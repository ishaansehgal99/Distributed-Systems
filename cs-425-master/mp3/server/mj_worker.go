package server

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"strings"

	pb "../ProtocolBuffers/ProtoPackage"
	"../config"
	"../fileops"
	"../logger"
	"../networking"
)

var (
	mjGUID = ""
)

// fetches “maple_exe” if it doesn’t already exist, and
// fetch maple input file to local dir system
// start goroutine for maple task
func handleMapleAssign(mjMessage *pb.MapleJuiceMessage) {
	logger.PrintInfo("maple task assigned, input file:", mjMessage.MapleInputFile, ", guid:", mjMessage.MjGUID)

	if mjGUID != mjMessage.MjGUID {
		fetchSdfsFileToLocal(mjMessage.ExeFilename)
		mjGUID = mjMessage.MjGUID
	}

	fetchSdfsFileToLocal(mjMessage.MapleInputFile)

	go mapleTaskRoutine(mjMessage)
}

// Recieve juice task information from master
// fetches “juice_exe” if it doesn’t already exist
// fetch the required intermediate files for this task to local dir system
// start goroutine for juice task
func handleJuiceAssign(mjMessage *pb.MapleJuiceMessage) {
	logger.PrintInfo("juice task assigned, input files:", mjMessage.JuiceInputKeys, ", guid:", mjMessage.MjGUID)

	if mjGUID != mjMessage.MjGUID {
		fetchSdfsFileToLocal(mjMessage.ExeFilename)
		mjGUID = mjMessage.MjGUID
	}

	// fetch intermediate files
	for _, inputKey := range mjMessage.JuiceInputKeys {
		fetchSdfsFileToLocal(mjMessage.IntermediatePrefix + "_" + inputKey)
	}

	go juiceTaskRoutine(mjMessage)
}

// runs maple_exe, pass in maple input file as argument
// capture (key, value) pairs from maple_exe and for each key PUT to intermediate file in SDFS system w/ key as filename
// send result back to master
func mapleTaskRoutine(mjMessage *pb.MapleJuiceMessage) {
	tempFilename := "temp"
	inputFile, _ := fileops.GetFilePointer(mjMessage.MapleInputFile)

	fileScanner := fileops.OpenFileScanner(inputFile)
	nextLines := fileops.ReadNextNLines(fileScanner, config.MAP_LINES)

	keysToValues := make(map[string][]string)

	for len(nextLines) > 0 {
		data := make([]byte, 0)

		for _, line := range nextLines {
			data = append(data, []byte(line+"\n")...)
		}

		ioutil.WriteFile(tempFilename, data, config.FILE_PERM)

		command := "./" + mjMessage.ExeFilename
		exeOutput := runExecutable(command, tempFilename)

		capturedLines := strings.Split(exeOutput, "\n")
		addToOutput(&keysToValues, capturedLines)

		nextLines = fileops.ReadNextNLines(fileScanner, config.MAP_LINES)
	}

	inputFile.Close()
	fileops.DeleteLocalFile(tempFilename)
	fileops.DeleteLocalFile(mjMessage.MapleInputFile)

	logger.PrintInfo("sending maple result to master")
	sendResultToMaster(keysToValues, mjMessage)
}

// read lines from intermediate files, run juice_exe and pass in all lines for a given key as argument
// capture (key, value) pairs from juice_exe and for each key PUT to “sdfs_dest_filename”
func juiceTaskRoutine(mjMessage *pb.MapleJuiceMessage) {
	keysToValues := make(map[string][]string)

	for _, inputKey := range mjMessage.JuiceInputKeys {
		inputFile := mjMessage.IntermediatePrefix + "_" + inputKey

		// juice_exe should expect a key to be passed in along with a list of values (delimited by the newline character)
		command := "./" + mjMessage.ExeFilename
		exeOutput := runExecutable(command, inputFile, inputKey)

		fileops.DeleteLocalFile(inputFile)

		capturedLines := strings.Split(exeOutput, "\n")
		addToOutput(&keysToValues, capturedLines)
	}

	logger.PrintInfo("sending juice result to master")
	sendResultToMaster(keysToValues, mjMessage)
}

func runExecutable(command string, args ...string) string {
	out, err := exec.Command(command, args...).Output()

	if err != nil {
		logger.PrintError("runExecutable:", err)
		return ""
	}

	return string(out)
}

func addToOutput(keysToValues *map[string][]string, lines []string) {
	for _, line := range lines {
		pair := strings.Split(line, " ")

		if len(pair) < 2 {
			break
		}

		key := pair[0]
		value := pair[1]

		if values, exists := (*keysToValues)[key]; exists {
			(*keysToValues)[key] = append(values, value)
		} else {
			newArr := make([]string, 0)
			(*keysToValues)[key] = append(newArr, value)
		}
	}
}

func sendResultToMaster(keysToValues map[string][]string, mjMessage *pb.MapleJuiceMessage) {
	// finalResult := make(map[string]*pb.StringList)

	// for key, values := range keysToValues {
	// 	finalResult[key] = &pb.StringList{
	// 		Items: values,
	// 	}
	// }
	// finalResult := make([]*pb.MJPair, 0)

	// for key, values := range keysToValues {
	// 	finalResult = append(finalResult, &pb.MJPair{
	// 		Key: key,
	// 		Values: &pb.StringList{
	// 			Items: values,
	// 		},
	// 	})

	// 	break
	// }

	// Marshal the map into a JSON string.
	empData, err := json.Marshal(keysToValues)
	if err != nil {
		logger.PrintError("json.Marshal:", err.Error())
		return
	}

	mjMessage.KeysToValues = empData
	mjMessage.Phase = pb.MapleJuicePhase_COMPLETE

	networking.SendMapleJuiceTCP(GetMasterIP(), mjMessage)
}
