package server

import (
	"bufio"
	"os"
	"strings"
	"sync"
	"time"

	pb "../ProtocolBuffers/ProtoPackage"
	"../client"
	"../config"
	"../fileops"
	"../logger"
	"../membership"
	"../networking"
	"../utils"
	"github.com/golang/protobuf/ptypes"
)

var (
	localMessage         *pb.MembershipMessage // Container for local membership list and filemap
	mux                  sync.Mutex            // Mutex lock for localMessage
	rwmux                sync.RWMutex          // Reader-writer lock for concurrent SDFS requests
	failureList          map[string]bool
	hashRing             map[string]uint32
	incomingMessageQueue []*pb.MapleJuiceMessage
	selfID               string
	masterIP             string
	isSending            bool
	isMaster             bool
	isJoining            bool
)

func GetMasterIP() string {
	mux.Lock()
	defer mux.Unlock()

	return masterIP
}

func GetSelfID() string {
	mux.Lock()
	defer mux.Unlock()

	return selfID
}

// CLI command parsing
func handleCommands(input string) {
	args := strings.Split(input, " ")
	cmd := args[0]
	param := ""

	if len(args) > 1 {
		param = args[1]
	}

	switch cmd {
	case "store":
		files := fileops.ListFiles()

		if files == nil || len(files) == 0 {
			logger.PrintInfo("No files are stored at this machine")
		}

		logger.PrintInfo("The following files are stored at this machine:\n", strings.Join(fileops.ListFiles(), "\n"))
	case "list":
		if param == "master" {
			logger.PrintInfo("master ID:", localMessage.MasterID)
		} else if param == "membership" {
			if localMessage != nil {
				mux.Lock()
				logger.PrintInfo("Printing membership list:\n", membership.GetMembershipListString(localMessage, failureList))
				mux.Unlock()
			} else {
				logger.PrintInfo("Membership list is nil")
			}
		} else if param == "self" {
			if selfID == "" {
				logger.PrintInfo("selfID is non-existent")
			} else {
				logger.PrintInfo(selfID)
			}
		} else {
			logger.PrintError("Invalid argument to 'list'")
		}
	case "leave":
		sendLeaveRequest()
	case "join":
		if param == "" {
			logger.PrintInfo("Please specify master IP address for joining")
		} else if !isSending {
			masterIP = param

			failureList = make(map[string]bool)
			hashRing = make(map[string]uint32)

			initLocalMessage()

			isJoining = true
			isSending = true
			go startHeartbeat()
			logger.PrintInfo("Successfully sent join request")
		} else {
			logger.PrintError("Cannot join, already actively sending")
		}
	case "maple":
		if len(args) == 5 {
			InitiateMaple(args[1:])
		} else {
			logger.PrintInfo("Invalid number of args for maple (need 4)")
		}
	case "juice":
		if len(args) == 7 {
			InitiateJuice(args[1:])
		} else {
			logger.PrintInfo("Invalid number of args for juice (need 6)")
		}
	default:
		client.HandleCommands(input, utils.GetIPFromID(selfID), masterIP, false, false)
	}
}

func sendLeaveRequest() {
	isSending = false

	mux.Lock()

	localMessage.MemberList[selfID].IsLeaving = true
	networking.HeartbeatGossip(localMessage, config.GOSSIP_FANOUT, selfID)

	selfID = ""
	localMessage = nil

	mux.Unlock()
	logger.PrintInfo("Successfully left")
}

func updateMasterIPIfNeeded() {
	if !strings.HasPrefix(localMessage.MasterID, masterIP) && len(localMessage.MasterID) > len(masterIP) {
		masterIP = utils.GetIPFromID(localMessage.MasterID)
	}
}

// Callback for when membership message is received on the UDP listener
func readNewMessage(message []byte) error {
	if !isSending {
		return nil
	}

	remoteMessage, err := networking.DecodeMembershipMessage(message)
	if err != nil {
		return err
	}

	mux.Lock()
	defer mux.Unlock()

	if isJoining && remoteMessage.Type == pb.MessageType_JOIN_REP {
		isJoining = false
		localMessage.Type = pb.MessageType_MEMBERSHIP
		localMessage.MasterID = remoteMessage.MasterID
		localMessage.MasterCounter = remoteMessage.MasterCounter
	}

	if !isMaster && remoteMessage.Type == pb.MessageType_JOIN_REQ {
		return nil
	}

	membership.MergeMembershipLists(localMessage, remoteMessage, failureList, hashRing, selfID)
	updateMasterIPIfNeeded()

	if isMaster && remoteMessage.Type == pb.MessageType_JOIN_REQ {
		logger.PrintInfo("Received join request")
		localMessage.Type = pb.MessageType_JOIN_REP
		message, err := networking.EncodeMembershipMessage(localMessage)
		localMessage.Type = pb.MessageType_MEMBERSHIP

		if err != nil {
			return err
		}

		dests := membership.GetOtherMembershipListIPs(remoteMessage, selfID)
		networking.Send(dests[0], message)
	}

	if localMessage.MasterID == selfID && !isMaster {
		isMaster = true
	}

	return nil
}

func reassignMapleJuiceTasksAsNeeded() {
	idMap := membership.GetIDMap(localMessage)

	idList := make([]string, 0)

	for id := range idMap {
		idList = append(idList, id)
	}

	handleMapleJuiceFailure(&idList)
}

// Goroutine function for membership + master gossip and updating local membership details
func startHeartbeat() {
	for isSending {
		mux.Lock()

		localMessage.MemberList[selfID].LastSeen = ptypes.TimestampNow()
		localMessage.MemberList[selfID].HeartbeatCounter++

		isNewFailure := membership.CheckAndRemoveMembershipListFailures(localMessage, &failureList, hashRing)

		if isNewFailure && isMaster {
			replicateAsNeeded()
			reassignMapleJuiceTasksAsNeeded()
		}

		updateMasterIPIfNeeded()
		logger.InfoLogger.Println("Current memlist:\n", membership.GetMembershipListString(localMessage, failureList))

		membership.UpdateLocalFileMap(localMessage, selfID)

		if isJoining {
			message, _ := networking.EncodeMembershipMessage(localMessage)
			networking.Send(masterIP, message)
		} else {
			networking.HeartbeatGossip(localMessage, config.GOSSIP_FANOUT, selfID)

			for machineID := range localMessage.MemberList {
				if localMessage.MemberList[machineID].IsLeaving && !failureList[machineID] {
					logger.PrintInfo("Received leave request from machine", machineID)
					failureList[machineID] = true
				}
			}
		}

		if localMessage.MasterID == selfID && !isMaster {
			isMaster = true
		}

		mux.Unlock()

		time.Sleep(config.PULSE_TIME * time.Millisecond)
	}
}

func initLocalMessage() {
	selfMember := pb.Member{
		HeartbeatCounter: 1,
		LastSeen:         ptypes.TimestampNow(),
	}

	localMessage = &pb.MembershipMessage{
		MemberList: make(map[string]*pb.Member),
		FileMap:    make(map[string]*pb.FileInfo),
	}

	localIP := networking.GetLocalIPAddr().String()
	selfID = localIP + ":" + ptypes.TimestampString(selfMember.LastSeen)

	if isMaster {
		localMessage.Type = pb.MessageType_MEMBERSHIP
		localMessage.MasterID = selfID
		localMessage.MasterCounter = 1
	} else {
		localMessage.Type = pb.MessageType_JOIN_REQ
		localMessage.MasterCounter = 0
	}

	membership.AddMemberToMembershipList(localMessage, selfID, &selfMember, hashRing)
}

// Server process entry point
func Run(isFirst bool, masIP string) {
	logger.PrintInfo("Starting server...\nIs master:", isFirst, "\nmasterIP:", masIP)
	isMaster = isFirst
	masterIP = masIP

	isSending = true
	isJoining = !isMaster

	fileops.InitSdfs()

	failureList = make(map[string]bool)
	hashRing = make(map[string]uint32)
	incomingMessageQueue = make([]*pb.MapleJuiceMessage, 0)
	initLocalMessage()

	logger.PrintInfo("Starting server with id", selfID, "on port", config.PORT)
	go networking.Listen(config.PORT, readNewMessage)
	go startHeartbeat()
	go networking.ListenTCP(readSdfsMessage)
	go networking.ListenMapleJuiceTCP(HandleMapleJuiceMessage)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		logger.InfoLogger.Println("Commandline input:", input)

		handleCommands(input)
	}

	if scanner.Err() != nil {
		logger.PrintError("Error reading input from commandline")
	}
}
