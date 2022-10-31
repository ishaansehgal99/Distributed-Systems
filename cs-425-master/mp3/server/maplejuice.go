package server

import (
	"fmt"
	"regexp"
	"time"

	"../client"
	"../logger"
	"../utils"

	pb "../ProtocolBuffers/ProtoPackage"
)

// Used as callback for HandleIncomingMapleJuiceMessage in the server
func HandleMapleJuiceMessage(mjMessage *pb.MapleJuiceMessage) {
	if mjMessage.Phase == pb.MapleJuicePhase_INIT {
		handleInit(mjMessage)
	} else if mjMessage.Phase == pb.MapleJuicePhase_ASSIGN {
		handleAssign(mjMessage)
	} else {
		handleWorkerResult(mjMessage)
	}
}

// Check that files with prefix “sdfs_src_directory” exist in the system  (using filemap)
func handleInit(mjMessage *pb.MapleJuiceMessage) {
	phaseType := "Maple"
	if mjMessage.TaskType == pb.MapleJuiceType_JUICE {
		phaseType = "Juice"
	}

	logger.PrintInfo("=====Received new", phaseType, "Init message=====")

	incomingMessageQueue = append(incomingMessageQueue, mjMessage)

	for len(incomingMessageQueue) > 1 && mjMessage != incomingMessageQueue[0] {
		time.Sleep(1 * time.Second)
	}

	// print queue

	if mjMessage.TaskType == pb.MapleJuiceType_MAPLE {
		handleMapleInit(mjMessage)
	} else {
		handleJuiceInit(mjMessage)
	}
}

func popFromIncomingMessagesQueue() {
	if len(incomingMessageQueue) <= 1 {
		incomingMessageQueue = incomingMessageQueue[:0]
	} else {
		incomingMessageQueue = incomingMessageQueue[1:]
	}
}

// Handle at worker
func handleAssign(mjMessage *pb.MapleJuiceMessage) {
	if mjMessage.TaskType == pb.MapleJuiceType_MAPLE {
		handleMapleAssign(mjMessage)
	} else {
		handleJuiceAssign(mjMessage)
	}
}

// Handle at master
func handleWorkerResult(mjMessage *pb.MapleJuiceMessage) {
	if mjMessage.TaskType == pb.MapleJuiceType_MAPLE {
		handleMapleWorkerResult(mjMessage)
	} else {
		handleJuiceWorkerResult(mjMessage)
	}
}

func removeSpecialChars(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		logger.PrintError(err)
	}

	return reg.ReplaceAllString(input, "")
}

func putLocalFileToSdfs(filename string, shouldAppend bool) {
	client.HandleCommands(
		fmt.Sprintf("put %s %s", filename, filename),
		utils.GetIPFromID(GetSelfID()),
		GetMasterIP(),
		true,
		shouldAppend,
	)
}

// Fetches the given file from SDFS to the local filesystem if it doesn't already exist
func fetchSdfsFileToLocal(filename string) {
	client.HandleCommands(
		fmt.Sprintf("get %s %s", filename, filename),
		utils.GetIPFromID(GetSelfID()),
		GetMasterIP(),
		false,
		false,
	)
}

func deleteSdfsFile(filename string) {
	client.HandleCommands(
		fmt.Sprintf("delete %s", filename),
		utils.GetIPFromID(GetSelfID()),
		GetMasterIP(),
		true,
		false,
	)
}
