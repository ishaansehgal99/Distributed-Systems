package server

import (
	"fmt"
	"strconv"
	"strings"

	pb "../ProtocolBuffers/ProtoPackage"
	"../fileops"
	"../logger"
	"../networking"
	"github.com/google/uuid"
)

// Expected Args:
// <maple_exe> <num_maples>
// <sdfs_intermediate_filename_prefix> <sdfs_src_directory>
func validateMapleArgs(args []string) bool {
	if len(args) != 4 {
		logger.PrintError("Incorrect number of maple arguments")
		return false
	}

	mapleExe, numMaples := args[0], args[1]

	// Check if maple exe file exists in local dir
	if !fileops.DoesLocalFileExist(mapleExe) {
		logger.PrintError("maple exe local file does not exist")
		return false
	}

	// Check if numMaples is a number
	if _, err := strconv.Atoi(numMaples); err != nil {
		fmt.Printf("num_maples %q is not a number\n", numMaples)
		return false
	}

	return true
}

// Expected Args:
// <juice_exe> <num_juices>
// <sdfs_intermediate_filename_prefix> <sdfs_dest_filename>
// delete_input={0,1} partition_type={range,hash}
func validateJuiceArgs(args []string) bool {
	if len(args) != 6 {
		logger.PrintError("Incorrect number of juice arguments")
		return false
	}

	juiceExe, numJuices, deleteInput, partitionType := args[0], args[1], args[4], args[5]

	// Check if juice exe file exists in local dir
	if !fileops.DoesLocalFileExist(juiceExe) {
		logger.PrintError("juice exe local file does not exist")
		return false
	}

	// Check if numJuices is a number
	if _, err := strconv.Atoi(numJuices); err != nil {
		fmt.Printf("num_juices %q is not a number\n", numJuices)
		return false
	}

	// Check that deleteInput is either 0 or 1
	if deleteInput != "delete_input=0" && deleteInput != "delete_input=1" {
		fmt.Println("delete_input={num} must be 0 or 1")
		return false
	}

	// Check that partitionType is either range or hash
	if partitionType != "partition_type=range" && partitionType != "partition_type=hash" {
		fmt.Println("partition_type={type} must be range or hash")
		return false
	}

	// return true
	return true
}

// ensure that "maple_exe" exists in *local* filesystem (not in sdfs)
// Send PUT request of “maple_exe” into SDFS system
// send MapleJuiceMessage to master
func InitiateMaple(args []string) {
	if validateMapleArgs(args) == false {
		logger.PrintError("Invalid maple args")
		return
	}
	mapleExe, numMaples, filePrefix, srcDir := args[0], args[1], args[2], args[3]
	putLocalFileToSdfs(mapleExe, false)

	numMaplesInt, _ := strconv.ParseInt(numMaples, 10, 32)

	mapleCommand := &pb.MapleJuiceMessage{
		Phase:              pb.MapleJuicePhase_INIT,
		TaskType:           pb.MapleJuiceType_MAPLE,
		NumWorkers:         int32(numMaplesInt),
		ExeFilename:        strings.ReplaceAll(mapleExe, "/", "_"),
		IntermediatePrefix: strings.ReplaceAll(filePrefix, "/", "_"),
		SrcDirectory:       srcDir,
		MjGUID:             uuid.New().String(),
	}

	sendRequestToMaster(mapleCommand)
}

// validate args
// ensure that "juice_exe" exists in *local* filesystem (not in sdfs)
// Send PUT request of "juice_exe" into SDFS systems
// send MapleJuiceMessage to master
func InitiateJuice(args []string) {
	if validateJuiceArgs(args) == false {
		logger.PrintError("Invalid juice args")
		return
	}

	juiceExe, numJuices, filePrefix, destFile, deleteInput, partitionType := args[0], args[1], args[2], args[3], args[4], args[5]
	putLocalFileToSdfs(juiceExe, false)

	numJuicesInt, _ := strconv.ParseInt(numJuices, 10, 32)

	partition := pb.Partition_HASH

	if partitionType == "partition_type=range" {
		partition = pb.Partition_RANGE
	}

	juiceCommand := &pb.MapleJuiceMessage{
		Phase:              pb.MapleJuicePhase_INIT,
		TaskType:           pb.MapleJuiceType_JUICE,
		NumWorkers:         int32(numJuicesInt),
		ExeFilename:        strings.ReplaceAll(juiceExe, "/", "_"),
		IntermediatePrefix: strings.ReplaceAll(filePrefix, "/", "_"),
		DestFilename:       strings.ReplaceAll(destFile, "/", "_"),
		DeleteInput:        deleteInput == "1",
		PartitionType:      partition,
		MjGUID:             uuid.New().String(),
	}

	sendRequestToMaster(juiceCommand)
}

func sendRequestToMaster(request *pb.MapleJuiceMessage) {
	// call networking code to send to master
	networking.SendMapleJuiceTCP(GetMasterIP(), request)
	logger.PrintInfo("sent MapleJuice request to master")
}
