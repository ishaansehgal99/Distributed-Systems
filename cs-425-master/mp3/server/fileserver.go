package server

import (
	"fmt"
	"net"
	"strings"
	"time"

	pb "../ProtocolBuffers/ProtoPackage"
	"../client"
	"../config"
	"../fileops"
	"../logger"
	"../membership"
	"../networking"
	"../utils"
	"github.com/jinzhu/copier"
)

func checkDirectoryExists(dir string) bool {
	mux.Lock()
	defer mux.Unlock()

	for k := range localMessage.FileMap {
		kDir := strings.Split(k, "_")[0]
		if kDir == dir {
			return true
		}
	}

	return false
}

func listFilesInDirectory(dir string) *[]string {
	out := make([]string, 0)

	mux.Lock()
	defer mux.Unlock()

	for k := range localMessage.FileMap {
		kDir := strings.Split(k, "_")[0]
		if kDir == dir {
			out = append(out, k)
		}
	}

	return &out
}

// Called when a failure is detected.
// Go through the master filemap and check files that have < configured replicas and replicate
func replicateAsNeeded() {
	// Gives buffer before starting replication in case of false positives
	// and for combining multiple failures into one replication step
	time.Sleep(3 * config.PULSE_TIME * time.Millisecond)

	// logger.PrintInfo("replication start...\n---")
	filesToReplicate := membership.GetFilesToReplicate(localMessage)

	for _, file := range filesToReplicate {
		machineInfos := localMessage.FileMap[file].MachineFileInfos

		machineIDs := membership.GetIDMap(localMessage)

		for machineID := range machineInfos {
			delete(machineIDs, machineID)
		}

		machineIPs := []string{}

		for machineID := range machineIDs {
			machineIPs = append(machineIPs, utils.GetIPFromID(machineID))
		}

		utils.ShuffleList(&machineIPs)

		offset := 0

		for i := 0; i < config.NUM_REPLICAS-len(machineInfos) && i+offset < len(machineIPs); i++ {
			machineToReplicate := machineIPs[i+offset]

			replicationMessage := &pb.SdfsMessage{
				Type:         pb.MessageType_REPLICATION,
				SdfsFilename: file,
			}

			// logger.PrintInfo("replicating file \"", file, "\" to", machineToReplicate)
			didSucceed := networking.SendSoloSdfsMessage(machineToReplicate, replicationMessage)

			if !didSucceed {
				offset++
				i--
			}
		}
	}

	// logger.PrintInfo("---\nreplication end")
}

// Called when master wishes to replicate a file on this server.
// Initiates a GET request to fetch the file to store on this machine's SDFS storage
func handleReplicationMessage(sdfsMessage *pb.SdfsMessage) {
	client.HandleCommands(
		fmt.Sprintf("get %s %s", sdfsMessage.SdfsFilename, sdfsMessage.SdfsFilename),
		utils.GetIPFromID(selfID),
		masterIP,
		true,
		false,
	)
}

// Process local file operations on the replica and send completion/failure messages back to master
func handleLocalFileOp(conn net.Conn, sdfsMessage *pb.SdfsMessage) {
	if sdfsMessage.FileOp == pb.FileOp_GET {
		c := make(chan []byte)
		go fileops.ReadFileThread(sdfsMessage.SdfsFilename, true, config.FILE_BUFFER_SIZE, c)

		hasRead := false
		for i := 0; ; i++ {
			read := <-c

			if read == nil {
				break
			}

			// send partial file
			sdfsMessage.Type = pb.MessageType_RETURN
			sdfsMessage.File = read
			sdfsMessage.ReturnString = "Get was successful"
			sdfsMessage.FileSeqNum = int32(i)
			networking.SendSdfsMessageAndClose(conn, sdfsMessage)

			hasRead = true
		}

		if hasRead {
			sdfsMessage.File = nil
			sdfsMessage.ReturnString = "Get was successful"
			sdfsMessage.FileSeqNum = -1
		} else {
			sdfsMessage.File = nil
			sdfsMessage.ReturnString = "Get was not successful"
		}

		networking.SendSdfsMessageAndClose(conn, sdfsMessage)
	} else if sdfsMessage.FileOp == pb.FileOp_PUT {
		shouldOverwrite := !sdfsMessage.ShouldAppend && sdfsMessage.FileSeqNum == 0

		didSucceed := fileops.PutFile(sdfsMessage.SdfsFilename, sdfsMessage.File, true, shouldOverwrite)

		if !isMaster {
			sdfsMessage.File = nil
		}

		if didSucceed {
			sdfsMessage.ReturnString = "Put was successful"
		} else {
			sdfsMessage.ReturnString = "Put was not successful"
		}

		if conn != nil {
			sdfsMessage.Type = pb.MessageType_RETURN
			networking.SendSdfsMessageAndClose(conn, sdfsMessage)
		}
	} else if sdfsMessage.FileOp == pb.FileOp_DELETE {
		didSucceed := fileops.DeleteFile(sdfsMessage.SdfsFilename)
		mux.Lock()
		delete(localMessage.FileMap, sdfsMessage.SdfsFilename)
		mux.Unlock()

		if didSucceed {
			sdfsMessage.ReturnString = "Delete was successful"
		} else {
			sdfsMessage.ReturnString = "Delete was not successful"
		}

		if conn != nil {
			sdfsMessage.Type = pb.MessageType_RETURN
			networking.SendSdfsMessageAndClose(conn, sdfsMessage)
		}
	}
}

// Process LS by reading master's filemap
func getLsReturnMessage(sdfsMessage *pb.SdfsMessage) *pb.SdfsMessage {
	sdfsMessage.Type = pb.MessageType_RETURN
	lsString := "ls successful. The file " + sdfsMessage.SdfsFilename + " is stored at the following machines:\n"

	mux.Lock()

	if fileInfo, exists := localMessage.FileMap[sdfsMessage.SdfsFilename]; exists && len(fileInfo.MachineFileInfos) > 0 {
		for machineID := range fileInfo.MachineFileInfos {
			lsString += utils.GetIPFromID(machineID) + "\n"
		}
	} else {
		lsString = "File not found in the system"
	}

	mux.Unlock()

	sdfsMessage.ReturnString = lsString
	return sdfsMessage
}

// 1. Find all machine ids with the file name
// 2. Create a sdfsmessage to get the file from the first machine
// 3. If failed, iterate through the machines to get the file
func handleGetFromClientAtMaster(clientConn net.Conn, sdfsMessage *pb.SdfsMessage) {
	sdfsMessage.Type = pb.MessageType_FILE_OP

	mux.Lock()
	fileInfo, exists := localMessage.FileMap[sdfsMessage.SdfsFilename]
	fileInfoCpy := pb.FileInfo{}
	copier.Copy(&fileInfoCpy, &fileInfo)
	mux.Unlock()

	if exists {
		for machineID := range fileInfoCpy.MachineFileInfos {
			machineIP := utils.GetIPFromID(machineID)

			if machineIP == utils.GetIPFromID(selfID) {
				handleLocalFileOp(clientConn, sdfsMessage)
				break
			} else {
				sdfsMessage.Type = pb.MessageType_FILE_OP
				replyMessage, conn := networking.SendSdfsMessageWithReply(machineIP, sdfsMessage)

				if replyMessage == nil {
					continue
				}

				networking.SendSdfsMessageAndClose(clientConn, replyMessage)

				seqNum := replyMessage.FileSeqNum

				didSucceed := false

				for seqNum != -1 {
					nextMessage, err := networking.GetSdfsMessageFromConn(conn)

					if err != nil {
						logger.PrintError("client.go:", err)
						break
					} else {
						networking.SendSdfsMessageAndClose(clientConn, nextMessage)
					}

					seqNum = nextMessage.FileSeqNum

					didSucceed = seqNum == -1
				}

				if didSucceed {
					if sdfsMessage.LocalFilename != "sdfs" {
						break
					}

					// If the GET request was for a file replication to a new replica,
					// sdfsMessage.LocalFilename == "sdfs", and we should add this replica machine to filemap
					mux.Lock()

					clientMachineID := ""

					for machineID := range localMessage.MemberList {
						if strings.HasPrefix(machineID, sdfsMessage.ClientIP) {
							clientMachineID = machineID
							break
						}
					}

					if clientMachineID != "" {
						membership.AddToFileMap(localMessage, clientMachineID, sdfsMessage.SdfsFilename)
					}

					mux.Unlock()
					break
				}
			}
		}
	} else {
		sdfsMessage.ReturnString = "Get was not successful - file not found"
		sdfsMessage.File = nil
		networking.SendSdfsMessageAndClose(clientConn, sdfsMessage)
	}
}

// Handle DELETE request by sending DELETEs to replica and replying to client
func handleDeleteFromClientAtMaster(clientConn net.Conn, sdfsMessage *pb.SdfsMessage) {
	sdfsMessage.Type = pb.MessageType_FILE_OP

	mux.Lock()
	fileInfo, exists := localMessage.FileMap[sdfsMessage.SdfsFilename]
	fileInfoCpy := pb.FileInfo{}
	copier.Copy(&fileInfoCpy, &fileInfo)
	mux.Unlock()

	if exists {
		for machineID := range fileInfoCpy.MachineFileInfos {
			machineIP := utils.GetIPFromID(machineID)

			if machineIP == utils.GetIPFromID(selfID) {
				handleLocalFileOp(nil, sdfsMessage)
			} else {
				networking.SendSdfsMessageWithReply(machineIP, sdfsMessage)
			}
		}
	} else {
		sdfsMessage.ReturnString = "Delete was not successful because file was not found"
	}

	mux.Lock()
	delete(localMessage.FileMap, sdfsMessage.SdfsFilename)
	mux.Unlock()

	sdfsMessage.Type = pb.MessageType_RETURN
	networking.SendSdfsMessageAndClose(clientConn, sdfsMessage)
}

// Thread for parallelizing transfering files over to the replicas
func putThread(channel chan bool, sdfsMessage *pb.SdfsMessage, machineID string) {
	machineIP := utils.GetIPFromID(machineID)
	// logger.PrintInfo("putThread - machineIP", machineIP)

	if machineIP == utils.GetIPFromID(selfID) {
		handleLocalFileOp(nil, sdfsMessage)
		mux.Lock()
		membership.AddToFileMap(localMessage, selfID, sdfsMessage.SdfsFilename)
		mux.Unlock()

		channel <- true
	} else {
		replyMessage, _ := networking.SendSdfsMessageWithReply(machineIP, sdfsMessage)

		if replyMessage == nil {
			channel <- false
		} else {
			mux.Lock()
			membership.AddToFileMap(localMessage, machineID, sdfsMessage.SdfsFilename)
			mux.Unlock()
			channel <- true
		}
	}
}

// Choose machines, start threads to forward file data to replicas, and reply to client
func handlePutFromClientAtMaster(clientConn net.Conn, sdfsMessage *pb.SdfsMessage) {
	if sdfsMessage.FileSeqNum == 0 {
		rwmux.Lock()
	}

	sdfsMessage.Type = pb.MessageType_FILE_OP
	channel := make(chan bool)

	mux.Lock()
	fileInfo, exists := localMessage.FileMap[sdfsMessage.SdfsFilename]
	fileInfoCpy := pb.FileInfo{}
	copier.Copy(&fileInfoCpy, &fileInfo)
	mux.Unlock()

	if exists {
		// file exists in system - put to machines that already have the file

		for machineID := range fileInfoCpy.MachineFileInfos {
			go putThread(channel, sdfsMessage, machineID)
		}

		for i := 0; i < len(fileInfoCpy.MachineFileInfos); i++ {
			<-channel
		}
	} else {
		// file doesn't exist in system - put to REPLICA_NUM random machines
		mux.Lock()
		machineIDs := membership.GetAllIDs(localMessage)
		mux.Unlock()

		utils.ShuffleList(&machineIDs)

		for i := 0; i < config.NUM_REPLICAS && i < len(machineIDs); i++ {
			go putThread(channel, sdfsMessage, machineIDs[i])
		}

		for i := 0; i < config.NUM_REPLICAS && i < len(machineIDs); i++ {
			<-channel
		}
	}

	sdfsMessage.ReturnString = "Put was successful"

	sdfsMessage.Type = pb.MessageType_RETURN
	networking.SendSdfsMessageAndClose(clientConn, sdfsMessage)

	isConnClosed := networking.IsConnClosed(clientConn)

	if sdfsMessage.FileSeqNum == -1 || isConnClosed {
		rwmux.Unlock()
	}
}

// Parse and perform incoming SDFS request and reply as necessary
func readSdfsMessage(conn net.Conn, sdfsMessage *pb.SdfsMessage) {
	if sdfsMessage.Type == pb.MessageType_REPLICATION {
		handleReplicationMessage(sdfsMessage)
	} else if isMaster && sdfsMessage.Type == pb.MessageType_CLIENT_REQUEST {

		if sdfsMessage.FileOp == pb.FileOp_GET {
			rwmux.RLock()
			handleGetFromClientAtMaster(conn, sdfsMessage)
			rwmux.RUnlock()
		} else if sdfsMessage.FileOp == pb.FileOp_PUT {
			handlePutFromClientAtMaster(conn, sdfsMessage)
		} else if sdfsMessage.FileOp == pb.FileOp_DELETE {
			rwmux.Lock()
			handleDeleteFromClientAtMaster(conn, sdfsMessage)
			rwmux.Unlock()
		} else if sdfsMessage.FileOp == pb.FileOp_LS {
			// separate case: handle LS directly
			rwmux.RLock()
			returnMessage := getLsReturnMessage(sdfsMessage)
			networking.SendSdfsMessageAndClose(conn, returnMessage)
			rwmux.RUnlock()
		}

	} else {
		if sdfsMessage.Type == pb.MessageType_CLIENT_REQUEST {
			logger.PrintError("non-master nodes shouldn't directly handle client requests")
		} else if sdfsMessage.Type == pb.MessageType_FILE_OP && sdfsMessage.FileOp == pb.FileOp_GET {
			handleLocalFileOp(conn, sdfsMessage)
		} else if sdfsMessage.Type == pb.MessageType_FILE_OP {
			handleLocalFileOp(nil, sdfsMessage)
			sdfsMessage.Type = pb.MessageType_RETURN
			networking.SendSdfsMessageAndClose(conn, sdfsMessage)
		}
	}
}
