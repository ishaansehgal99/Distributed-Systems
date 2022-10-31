package client

import (
	"bufio"
	"os"
	"strings"
	"time"

	pb "../ProtocolBuffers/ProtoPackage"
	"../config"
	"../fileops"
	"../logger"
	"../networking"
)

func handleClientPut(args []string, masterIP string, clientMessage *pb.SdfsMessage) {
	clientMessage.FileOp = pb.FileOp_PUT
	clientMessage.SdfsFilename = args[2]

	c := make(chan []byte)
	go fileops.ReadFileThread(args[1], false, config.FILE_BUFFER_SIZE, c)

	hasRead := false
	for i := 0; ; i++ {
		read := <-c

		if read == nil {
			break
		}

		// send partial file
		clientMessage.File = read
		clientMessage.FileSeqNum = int32(i)
		networking.SendSdfsMessageWithReply(masterIP, clientMessage)

		hasRead = true
	}

	if hasRead {
		clientMessage.File = nil
		clientMessage.FileSeqNum = -1
		replyMessage, _ := networking.SendSdfsMessageWithReply(masterIP, clientMessage)

		if replyMessage != nil {
			logger.PrintInfo(replyMessage.ReturnString)
		} else {
			logger.PrintInfo("connection with master failed")
		}
	} else {
		logger.PrintError("Cannot find or read file:", args[1])
	}
}

func HandleCommands(input, clientIP, masterIP string, isSdfsDir bool) {
	startTime := time.Now()

	args := strings.Split(input, " ")

	clientMessage := &pb.SdfsMessage{
		Type:     pb.MessageType_CLIENT_REQUEST,
		ClientIP: clientIP,
	}

	if args[0] == "put" && len(args) == 3 {
		handleClientPut(args, masterIP, clientMessage)

		elapsedTime := time.Now().Sub(startTime).Milliseconds()
		logger.PrintInfo("time elapsed:", elapsedTime, "ms")
		return
	} else if args[0] == "get" && len(args) == 3 {
		clientMessage.FileOp = pb.FileOp_GET
		clientMessage.SdfsFilename = args[1]

		if isSdfsDir {
			clientMessage.LocalFilename = "sdfs"
		} else {
			clientMessage.LocalFilename = ""
		}
	} else if args[0] == "delete" && len(args) == 2 {
		clientMessage.FileOp = pb.FileOp_DELETE
		clientMessage.SdfsFilename = args[1]
	} else if args[0] == "ls" && len(args) == 2 {
		clientMessage.FileOp = pb.FileOp_LS
		clientMessage.SdfsFilename = args[1]
	} else {
		logger.PrintError("Invalid client command:", input)
		return
	}

	replyMessage, conn := networking.SendSdfsMessageWithReply(masterIP, clientMessage)

	if replyMessage == nil {
		logger.PrintError("failure communicating with master")
		return
	}

	if replyMessage.FileOp == pb.FileOp_GET && replyMessage.File != nil {
		localFilename := args[2]
		fileops.PutFile(localFilename, replyMessage.File, isSdfsDir, replyMessage.FileSeqNum == 0)
		seqNum := replyMessage.FileSeqNum

		for seqNum != -1 {
			nextMessage, err := networking.GetSdfsMessageFromConn(conn)

			if err != nil {
				logger.PrintError("client.go:", err)
				break
			} else {
				fileops.PutFile(localFilename, nextMessage.File, isSdfsDir, nextMessage.FileSeqNum == 0)
			}

			seqNum = nextMessage.FileSeqNum
		}
	}

	logger.PrintInfo(replyMessage.ReturnString)

	elapsedTime := time.Now().Sub(startTime).Milliseconds()
	logger.PrintInfo("time elapsed:", elapsedTime, "ms")
}

func Run(master string) {
	clientIP := networking.GetLocalIPAddr().String()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		HandleCommands(input, clientIP, master, false)
	}

	if scanner.Err() != nil {
		logger.PrintError("Error reading input from commandline")
	}
}
