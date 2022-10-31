package membership

import (
	"sort"
	"strconv"
	"strings"

	pb "../ProtocolBuffers/ProtoPackage"
	"../config"
	"../fileops"
	"../logger"
	"../utils"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
)

// MergeMembershipLists : merge remote membership list into local membership list
func MergeMembershipLists(
	localMessage,
	remoteMessage *pb.MembershipMessage,
	failureList map[string]bool,
	hashRing map[string]uint32,
	selfID string,
) {
	shouldOverrideMaster := remoteMessage.MasterID != localMessage.MasterID &&
		len(remoteMessage.MasterID) > 0 &&
		remoteMessage.MasterCounter == localMessage.MasterCounter &&
		utils.HashString(remoteMessage.MasterID) < utils.HashString(localMessage.MasterID)

	if remoteMessage.MasterCounter > localMessage.MasterCounter || shouldOverrideMaster {
		localMessage.MasterID = remoteMessage.MasterID
		localMessage.MasterCounter = remoteMessage.MasterCounter
		logger.PrintInfo("Change to new master:", localMessage.MasterID)
	}

	for machineID, member := range remoteMessage.MemberList {
		if _, ok := localMessage.MemberList[machineID]; !ok {
			if remoteMessage.MemberList[machineID].IsLeaving {
				break
			}

			memberCpy := pb.Member{}
			copier.Copy(&memberCpy, &member)
			AddMemberToMembershipList(localMessage, machineID, &memberCpy, hashRing)
			continue
		} else if remoteMessage.MemberList[machineID].IsLeaving {
			localMessage.MemberList[machineID].IsLeaving = true
		}

		remoteHeartBeat := remoteMessage.MemberList[machineID].HeartbeatCounter

		if localMessage.MemberList[machineID].HeartbeatCounter < remoteHeartBeat {
			delete(failureList, machineID)
			localMessage.MemberList[machineID].HeartbeatCounter = remoteHeartBeat
			localMessage.MemberList[machineID].LastSeen = ptypes.TimestampNow()
		}
	}

	for fileName, remoteFileInfo := range remoteMessage.FileMap {
		if _, fileExists := localMessage.FileMap[fileName]; !fileExists {
			// FileInfo doesn't exist, copy whole thing over
			fileInfoCpy := pb.FileInfo{}
			copier.Copy(&fileInfoCpy, &remoteFileInfo)

			localMessage.FileMap[fileName] = &fileInfoCpy
		} else {
			// FileInfo does exist, copy over MachineFileInfo objects that are not about this machine
			remoteFileMachineInfos := remoteMessage.FileMap[fileName].MachineFileInfos

			for machineID, remoteMachineFileInfo := range remoteFileMachineInfos {
				if machineID == selfID {
					continue
				}

				if _, machineInfoExists := localMessage.FileMap[fileName].MachineFileInfos[machineID]; !machineInfoExists {
					// MachineFileInfo doesn't exist, copy whole thing over

					machineFileInfoCpy := pb.MachineFileInfo{}
					copier.Copy(&machineFileInfoCpy, &remoteMachineFileInfo)

					localMessage.FileMap[fileName].MachineFileInfos[machineID] = &machineFileInfoCpy
				} else if localMessage.FileMap[fileName].MachineFileInfos[machineID].Counter < remoteMachineFileInfo.Counter {
					localMessage.FileMap[fileName].MachineFileInfos[machineID].Counter = remoteMachineFileInfo.Counter
					localMessage.FileMap[fileName].MachineFileInfos[machineID].LastSeen = ptypes.TimestampNow()
				}
			}
		}
	}
}

// Gets all the machine IDs from a given membership list
func GetAllIDs(message *pb.MembershipMessage) []string {
	machineIDs := []string{}

	for machineID := range message.MemberList {
		machineIDs = append(machineIDs, machineID)
	}

	return machineIDs
}

// Gets all the machine IDs from a given membership list in a map
func GetIDMap(message *pb.MembershipMessage) map[string]bool {
	machineIDs := make(map[string]bool)

	for machineID := range message.MemberList {
		machineIDs[machineID] = true
	}

	return machineIDs
}

// GetOtherMembershipListIPs : Expecting MachineID to be in format IP:timestamp
func GetOtherMembershipListIPs(message *pb.MembershipMessage, selfID string) []string {
	ips := make([]string, 0, len(message.MemberList))

	for machineID := range message.MemberList {
		if machineID != selfID && machineID != message.MasterID {
			ips = append(ips, utils.GetIPFromID(machineID))
		}
	}

	return ips
}

// Get a list of files that we should replicate based on the filemap
func GetFilesToReplicate(message *pb.MembershipMessage) []string {
	files := []string{}
	for fileName, fileInfo := range message.FileMap {
		if len(fileInfo.MachineFileInfos) < config.NUM_REPLICAS {
			files = append(files, fileName)
		}
	}
	return files
}

// CheckAndRemoveMembershipListFailures : Upon sending of membership list mark failures and remove failed machines
func CheckAndRemoveMembershipListFailures(message *pb.MembershipMessage, failureList *map[string]bool, hashRing map[string]uint32) bool {
	hasNewFailure := false

	for machineID, member := range message.MemberList {
		timeElapsedSinceLastSeen := float64(ptypes.TimestampNow().GetSeconds() - member.LastSeen.GetSeconds())

		if timeElapsedSinceLastSeen >= config.T_TIMEOUT+config.T_CLEANUP {
			delete(*failureList, machineID)
			RemoveMemberFromMembershipList(message, machineID, hashRing)
			hasNewFailure = true

			for _, fileInfo := range message.FileMap {
				for fileMachineID := range fileInfo.MachineFileInfos {
					if machineID == fileMachineID {
						delete(fileInfo.MachineFileInfos, machineID)
					}
				}
			}

		} else if !(*failureList)[machineID] && timeElapsedSinceLastSeen >= config.T_TIMEOUT {
			(*failureList)[machineID] = true
			logger.PrintInfo("Marking machine", machineID, "as failed")
		}
	}

	for _, fileInfo := range message.FileMap {
		for machineID, machineFileInfo := range fileInfo.MachineFileInfos {
			timeElapsedSinceLastSeen := float64(ptypes.TimestampNow().GetSeconds() - machineFileInfo.LastSeen.GetSeconds())

			if timeElapsedSinceLastSeen >= config.T_TIMEOUT+config.T_CLEANUP {
				delete(fileInfo.MachineFileInfos, machineID)
			}
		}
	}

	return hasNewFailure
}

// AddMemberToMembershipList : add new member to membership list
func AddMemberToMembershipList(message *pb.MembershipMessage, machineID string, member *pb.Member, hashRing map[string]uint32) {
	logger.PrintInfo("Adding machine", machineID, "to membership list")

	message.MemberList[machineID] = member
	hashRing[machineID] = utils.HashString(machineID)
}

// RemoveMemberFromMembershipList : remove member from membership list
func RemoveMemberFromMembershipList(message *pb.MembershipMessage, machineID string, hashRing map[string]uint32) {
	logger.PrintInfo("Removing machine", machineID, "from membership list")

	delete(message.MemberList, machineID)
	delete(hashRing, machineID)

	if machineID == message.MasterID {
		message.MasterCounter++
		message.MasterID = utils.GetLowestHashRingEntry(hashRing)
		logger.PrintInfo("Change to new master:", message.MasterID)
	}
}

// Get a printable string from the membership message
func GetMembershipListString(message *pb.MembershipMessage, failureList map[string]bool) string {
	var sb strings.Builder

	machineIDs := make([]string, 0)
	for k := range message.MemberList {
		machineIDs = append(machineIDs, k)
	}

	sort.Strings(machineIDs)

	for _, machineID := range machineIDs {
		if failureList[machineID] {
			sb.WriteString("FAILED:")
		}
		sb.WriteString(machineID +
			" - { HeartbeatCounter: " +
			strconv.Itoa(int(message.MemberList[machineID].HeartbeatCounter)) +
			", LastSeen: " +
			ptypes.TimestampString(message.MemberList[machineID].LastSeen) +
			" }\n")
	}

	sb.WriteString("\n")

	return sb.String()
}

// Add a file to the local filemap
func AddToFileMap(message *pb.MembershipMessage, machineID, filename string) {
	message.FileMap[filename] = &pb.FileInfo{
		MachineFileInfos: make(map[string]*pb.MachineFileInfo),
	}

	message.FileMap[filename].MachineFileInfos[machineID] = &pb.MachineFileInfo{
		Counter:  1,
		LastSeen: ptypes.TimestampNow(),
	}
}

// Update the local filemap based on this machine's files in the SDFS directory
func UpdateLocalFileMap(message *pb.MembershipMessage, selfID string) {
	for _, filename := range fileops.ListFiles() {
		if fileInfo, fileExists := message.FileMap[filename]; fileExists {

			if _, selfExists := fileInfo.MachineFileInfos[selfID]; selfExists {
				fileInfo.MachineFileInfos[selfID].Counter++
				fileInfo.MachineFileInfos[selfID].LastSeen = ptypes.TimestampNow()
				continue
			}

			if fileInfo.MachineFileInfos == nil {
				fileInfo.MachineFileInfos = make(map[string]*pb.MachineFileInfo)
			}

			fileInfo.MachineFileInfos[selfID] = &pb.MachineFileInfo{
				Counter:  1,
				LastSeen: ptypes.TimestampNow(),
			}

			continue
		}

		AddToFileMap(message, selfID, filename)
	}
}
