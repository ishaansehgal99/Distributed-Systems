package networking

import (
	pb "../ProtocolBuffers/ProtoPackage"
	"../membership"
	"../utils"
	proto "github.com/golang/protobuf/proto"
)

// Encode a Membership message into a byte array
func EncodeMembershipMessage(serviceMessage *pb.MembershipMessage) ([]byte, error) {
	message, err := proto.Marshal(serviceMessage)

	return message, err
}

// Decode a byte array into a Membership message
func DecodeMembershipMessage(message []byte) (*pb.MembershipMessage, error) {
	list := &pb.MembershipMessage{}
	err := proto.Unmarshal(message, list)

	return list, err
}

// Send a gossip message to k random machines from your membership list
func SendGossip(serviceMessage *pb.MembershipMessage, k int, selfID string) error {
	dests := membership.GetOtherMembershipListIPs(serviceMessage, selfID)

	if k < len(serviceMessage.MemberList) {
		utils.ShuffleList(&dests)

		dests = dests[:(k - 1)]
	}

	dests = append(dests, utils.GetIPFromID(serviceMessage.MasterID))

	return SendAll(dests, serviceMessage)
}

// Initiate the gossip protocol
func HeartbeatGossip(serviceMessage *pb.MembershipMessage, k int, selfID string) error {
	return SendGossip(serviceMessage, k, selfID)
}
