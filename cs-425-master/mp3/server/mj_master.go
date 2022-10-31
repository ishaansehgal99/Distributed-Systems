package server

import (
	"encoding/json"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	pb "../ProtocolBuffers/ProtoPackage"
	"../config"
	"../fileops"
	"../logger"
	"../membership"
	"../networking"
	"../utils"
)

var (
	taskMap        map[string][]string   // map from machineIP (not ID) to list of task ids
	currentMessage *pb.MapleJuiceMessage // contains details about the current run (such as are we doing Maple or Juice)
	mjStartTime    time.Time
)

// Check that files with prefix “sdfs_src_directory” exist in the system (using filemap)
// The master GETS all of the sdfs_src_directory files into its local system and then PUTs into the SDFS system “num_maples” files of even length (and then deletes all of the local system files that it used) - remember to use lines, not bytes (Aggregate into one big file and split into smaller files)
// Master sends “num_maples” messages to machines with maple task information
func handleMapleInit(mjMessage *pb.MapleJuiceMessage) {
	mjStartTime = time.Now()
	logger.PrintInfo("=====Scheduling new Maple tasks=====")

	// initialize globals for current maple run
	initializeTaskMap()
	currentMessage = mjMessage

	// check that sdfs_src_directory exists and contains files
	dir := mjMessage.SrcDirectory
	inputFiles := listFilesInDirectory(dir)

	if len(*inputFiles) == 0 {
		logger.PrintError("Specified directory (" + dir + ") contained no files!")
		popFromIncomingMessagesQueue()
		return
	}

	// split up entire dataset into num_maples number of files
	numMaples := int(mjMessage.NumWorkers)
	numTasks, err := createMapleInputFiles(inputFiles, numMaples)
	mjMessage.NumWorkers = int32(numTasks)
	currentMessage = mjMessage

	if err != nil {
		logger.PrintError("Error splitting maple input", err)
		popFromIncomingMessagesQueue()
		return
	}

	// initial assignment of tasks
	tasks := make([]string, 0)
	for i := 0; i < numTasks; i++ {
		tasks = append(tasks, strconv.Itoa(i))
	}

	// TODO: might want to do error handling
	assignTasksToMachines(&tasks) // takes in *[]string
}

// Check that files with prefix “sdfs_intermediate_filename_prefix” exist in the system (using filemap)
// assign intermediate files to different tasks based on either hash partitioning or range partitioning (this is included as a command-line input for juice)
// Master sends “num_juices” messages to machines with juice task information.
func handleJuiceInit(mjMessage *pb.MapleJuiceMessage) {
	mjStartTime = time.Now()
	logger.PrintInfo("=====Received Juice Init message=====")

	// init map
	initializeTaskMap()
	currentMessage = mjMessage

	// get keys by listing all files with the specified intermediate prefix
	prefix := mjMessage.IntermediatePrefix
	inputFiles := listFilesInDirectory(prefix)

	if len(*inputFiles) == 0 {
		logger.PrintError("Specified intermediate prefix (" + prefix + ") contained no files!")
		popFromIncomingMessagesQueue()
		return
	}

	// extract keys from intermediate filenames
	keys := make([]string, 0)
	for _, filename := range *inputFiles {
		k := getTaskIDFromInputFile(filename)
		keys = append(keys, k)
	}

	// partition keys into groups (tasks)
	numWorkers := int(mjMessage.NumWorkers)
	var tasks *[]string

	if mjMessage.PartitionType == pb.Partition_HASH {
		tasks = hashPartition(&keys, numWorkers) // returns *[]string
	} else if mjMessage.PartitionType == pb.Partition_RANGE {
		tasks = rangePartition(&keys, numWorkers)
	}

	// assign groups to workers
	// TODO: might want to do error handling
	assignTasksToMachines(tasks) // takes in *[]string
}

func handleMapleWorkerResult(mjMessage *pb.MapleJuiceMessage) {
	logger.PrintInfo("=====Received completed Maple job=====")

	if mjMessage.Phase == pb.MapleJuicePhase_COMPLETE {
		// Mark task as completed by removing from taskMap
		inputFile := mjMessage.MapleInputFile
		taskID := getTaskIDFromInputFile(inputFile)
		removeTaskFromTaskMap(taskID)

		var keysToValues map[string][]string
		err := json.Unmarshal(mjMessage.KeysToValues, &keysToValues)
		if err != nil {
			logger.PrintError("json.Unmarshal:", err.Error())
			return
		}

		// PUT (append) result to intermediate file
		for key, values := range keysToValues {
			intermediateFilename := mjMessage.IntermediatePrefix + "_" + removeSpecialChars(key)
			fileContents := []byte(strings.Join(values, "\n") + "\n")

			fileops.PutFile(intermediateFilename, fileContents, false, false)
			putLocalFileToSdfs(intermediateFilename, true)

			fileops.DeleteLocalFile(intermediateFilename)
		}

		// if task map is empty, print out "maple is finished"
		if isTaskMapEmpty() {
			logger.PrintInfo("=====Maple task finished!=====")

			elapsedTime := time.Now().Sub(mjStartTime).Milliseconds()
			logger.PrintInfo("time elapsed:", elapsedTime, "ms")

			logger.PrintInfo("==============================")
			popFromIncomingMessagesQueue()
		}
	} else {
		// handle fail (may not be needed here)
	}
}

// after juice task has finished, if delete_input==1, delete all intermediate files
func handleJuiceWorkerResult(mjMessage *pb.MapleJuiceMessage) {
	logger.PrintInfo("=====Received completed Juice job=====")

	if mjMessage.Phase == pb.MapleJuicePhase_COMPLETE {
		taskID := strings.Join(mjMessage.JuiceInputKeys, "\n")
		removeTaskFromTaskMap(taskID)

		fileContents := make([]byte, 0)

		var keysToValues map[string][]string
		err := json.Unmarshal(mjMessage.KeysToValues, &keysToValues)
		if err != nil {
			logger.PrintError("json.Unmarshal:", err.Error())
			return
		}

		for key, values := range keysToValues {
			for _, value := range values {
				fileContents = append(fileContents, []byte(key+" "+value+"\n")...)
			}
		}

		fileops.DeleteLocalFile(mjMessage.DestFilename)
		fileops.PutFile(mjMessage.DestFilename, fileContents, false, false)
		putLocalFileToSdfs(mjMessage.DestFilename, true)

		if mjMessage.DeleteInput {
			for _, inputKey := range mjMessage.JuiceInputKeys {
				inputFile := mjMessage.IntermediatePrefix + "_" + inputKey
				deleteSdfsFile(inputFile)
			}
		}

		if isTaskMapEmpty() {
			logger.PrintInfo("=====Juice task finished!=====")

			elapsedTime := time.Now().Sub(mjStartTime).Milliseconds()
			logger.PrintInfo("time elapsed:", elapsedTime, "ms")

			logger.PrintInfo("==============================")

			fileops.DeleteLocalFile(mjMessage.DestFilename)
			popFromIncomingMessagesQueue()
		}
	} else {
		// handle fail (may not be needed here)
	}
}

func initializeTaskMap() {
	taskMap = make(map[string][]string)

	mux.Lock()
	machineIDs := membership.GetIDMap(localMessage)
	mux.Unlock()
	for id := range machineIDs {
		machineIP := utils.GetIPFromID(id)
		taskMap[machineIP] = make([]string, 0)
	}
}

func removeTaskFromTaskMap(taskID string) {
	for machineIP := range taskMap {
		tasks := taskMap[machineIP]
		isDone := false

		for i := range tasks {
			if tasks[i] == taskID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				taskMap[machineIP] = tasks

				isDone = true
				break
			}
		}

		if isDone {
			break
		}
	}
}

func isTaskMapEmpty() bool {
	for machineIP := range taskMap {
		tasks := taskMap[machineIP]

		if len(tasks) != 0 {
			return false
		}
	}

	return true
}

func getTaskIDFromInputFile(inputFile string) string {
	split := strings.Split(inputFile, "_")

	taskID := split[len(split)-1]
	return taskID
}

func fetchSdfsMapleInputFilesToLocal(linesLeftInFile *map[string]int, inputFiles *[]string) int {
	totalLines := 0
	for _, filename := range *inputFiles {
		// get file to local
		fetchSdfsFileToLocal(filename)

		numLinesInFile, err := fileops.NumLinesInFile(filename)
		if err != nil {
			logger.PrintError("Error splitting maple input", err)
			return -1
		}

		(*linesLeftInFile)[filename] = numLinesInFile
		// add number of lines in this file to "totalLines"
		totalLines += numLinesInFile
	}

	return totalLines
}

func putLocalSplitMapleInputFilesToSdfs(numMaples int) {
	for idx := 0; idx < numMaples; idx++ {
		filename := config.MAPLE_INPUT_FILE_PREFIX + strconv.Itoa(idx)
		putLocalFileToSdfs(filename, false)
		fileops.DeleteLocalFile(filename)
	}
}

// split up files for maple
func createMapleInputFiles(inputFiles *[]string, numMaples int) (int, error) {
	numLinesLeftToReadInFile := make(map[string]int)
	totalLines := fetchSdfsMapleInputFilesToLocal(&numLinesLeftToReadInFile, inputFiles)

	numLinesForEachMaple := totalLines / numMaples

	if numLinesForEachMaple < 1 {
		numLinesForEachMaple = totalLines
		numMaples = 1
	}

	currFileIdx := 0
	f, err := fileops.GetFilePointer((*inputFiles)[currFileIdx])

	if err != nil {
		logger.PrintError("Error opening file in Maple create input", err)
		return 0, err
	}

	scanner := fileops.OpenFileScanner(f)

	mFile := &os.File{}
	for idx := 0; idx < numMaples; idx++ {
		numLinesRead := 0
		mFile, err = os.Create(config.MAPLE_INPUT_FILE_PREFIX + strconv.Itoa(idx))

		if err != nil {
			logger.PrintError("Error creating maple intermediate input", err)
			return 0, err
		}

		for numLinesRead < numLinesForEachMaple {
			currFile := (*inputFiles)[currFileIdx]
			numLinesLeftInFile := numLinesLeftToReadInFile[currFile]

			if numLinesLeftInFile+numLinesRead <= numLinesForEachMaple {
				lines := fileops.ReadNextNLines(scanner, numLinesLeftInFile)

				fileops.WriteLinesToFile(lines, mFile)

				numLinesRead += numLinesLeftInFile
				numLinesLeftToReadInFile[currFile] = 0

				f.Close()
				currFileIdx++

				if currFileIdx >= len(*inputFiles) {
					break
				}

				f, err = fileops.GetFilePointer((*inputFiles)[currFileIdx])

				if err != nil {
					logger.PrintError("Error opening file in Maple create input", err)
					return 0, err
				}

				scanner = fileops.OpenFileScanner(f)
			} else {
				//  numLinesLeftInFile + numLinesRead > numLinesForEachMaple
				nMoreLinesToRead := numLinesForEachMaple - numLinesRead
				lines := fileops.ReadNextNLines(scanner, nMoreLinesToRead)

				fileops.WriteLinesToFile(lines, mFile)

				numLinesRead += nMoreLinesToRead
				numLinesLeftToReadInFile[currFile] -= nMoreLinesToRead
			}
		}
	}

	for currFileIdx < len((*inputFiles)) {
		currFile := (*inputFiles)[currFileIdx]
		numLinesLeftInFile := numLinesLeftToReadInFile[currFile]

		lines := fileops.ReadNextNLines(scanner, numLinesLeftInFile)
		fileops.WriteLinesToFile(lines, mFile)

		numLinesLeftToReadInFile[currFile] = 0
		currFileIdx++
	}

	putLocalSplitMapleInputFilesToSdfs(numMaples)

	return numMaples, nil
}

// Assign the task in mjMessage to the specified machine
func assignTasksToMachines(tasks *[]string) {
	// sort machines by lowest task count
	type kv struct {
		key   string
		value int
	}

	machines := make([]kv, 0)

	for k, v := range taskMap {
		machines = append(machines, kv{k, len(v)})
	}

	sort.Slice(machines, func(i, j int) bool {
		return machines[i].value < machines[j].value
	})

	// select the first len(tasks) machine to assign
	machineIdx := 0
	for i := 0; i < len(*tasks); i++ {
		machineIP := machines[machineIdx].key
		curTask := (*tasks)[i]

		// TODO: might want to do error handling here
		if currentMessage.TaskType == pb.MapleJuiceType_MAPLE {
			sendMapleTaskToWorker(machineIP, curTask, currentMessage)
		} else if currentMessage.TaskType == pb.MapleJuiceType_JUICE {
			sendJuiceTaskToWorker(machineIP, curTask, currentMessage)
		}

		taskMap[machineIP] = append(taskMap[machineIP], curTask)

		machineIdx = (machineIdx + 1) % len(machines)
	}
}

func sendMapleTaskToWorker(machineIP string, task string, mjMessage *pb.MapleJuiceMessage) {
	mjMessage.Phase = pb.MapleJuicePhase_ASSIGN
	mjMessage.MapleInputFile = config.MAPLE_INPUT_FILE_PREFIX + task

	networking.SendMapleJuiceTCP(machineIP, mjMessage)
}

func sendJuiceTaskToWorker(machineIP string, task string, mjMessage *pb.MapleJuiceMessage) {
	mjMessage.Phase = pb.MapleJuicePhase_ASSIGN

	// expand task out into list of keys
	keys := strings.Split(task, "\n")

	mjMessage.JuiceInputKeys = keys

	networking.SendMapleJuiceTCP(machineIP, mjMessage)
}

// method for re-assigning in the event of failure
// (should be called in membership, just like replicateAsNeeded())

// detect failure in membership. call handleMapleJuiceFailure()
// update the taskmap with updated membership list
// failedTasks = list what tasks the failed machines were working on
// assignMapleTasksToMachines(failedTasks)
func handleMapleJuiceFailure(aliveMachineIDs *[]string) {
	if taskMap == nil {
		return
	}

	// logger.PrintInfo("")
	// create aliveMachineIDs set
	isAlive := make(map[string]bool)
	for _, machineID := range *aliveMachineIDs {
		machineIP := utils.GetIPFromID(machineID)
		isAlive[machineIP] = true
	}

	// create failedTasks
	failedTasks := make([]string, 0)
	for machineIP, tasks := range taskMap {
		_, ok := isAlive[machineIP]

		if !ok {
			failedTasks = append(failedTasks, tasks...)
			delete(taskMap, machineIP)
		}
	}

	// add new alive processes to taskMap
	for machineIP := range isAlive {
		_, exists := taskMap[machineIP]
		if !exists {
			taskMap[machineIP] = make([]string, 0)
		}
	}

	// reassign failed tasks
	assignTasksToMachines(&failedTasks)
}

// partition keys into groups of keys (newline delimited string) (reducer only)
func hashPartition(keys *[]string, numWorkers int) *[]string {
	m := make(map[int][]string)

	for _, k := range *keys {
		hash := int(utils.HashString(k)) % numWorkers

		m[hash] = append(m[hash], k)
	}

	tasks := make([]string, 0)
	for hash := range m {
		list := m[hash]

		tasks = append(tasks, strings.Join(list, "\n"))
	}

	return &tasks
}

func rangePartition(keys *[]string, numWorkers int) *[]string {
	sort.Strings(*keys)

	tasks := make([]string, 0)

	groupSize := int(math.Ceil(float64(len(*keys)) / float64(numWorkers)))
	for i := 0; i < numWorkers; i++ {
		start := i * groupSize
		end := start + groupSize

		if end > len(*keys) {
			end = len(*keys)
		}

		if start > end {
			continue
		}

		tasks = append(tasks, strings.Join((*keys)[start:end], "\n"))
	}

	return &tasks
}
