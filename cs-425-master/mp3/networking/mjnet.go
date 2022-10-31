package networking

import (
	"io"

	"../config"
	"../logger"
	proto "github.com/golang/protobuf/proto"

	"net"

	pb "../ProtocolBuffers/ProtoPackage"
)

// Set up a TCP server on the given port, which calls the given callback function when it receives a new message
func ListenMapleJuiceTCP(callback func(message *pb.MapleJuiceMessage)) error {
	listener, err := net.Listen("tcp", ":"+config.MAPLEJUICE_PORT)
	if err != nil {
		logger.PrintError(err)
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.PrintError(err)
			return err
		}

		handleIncomingMapleJuiceMessage(conn, callback)
	}
}

// Send the given byte array to destination on MAPLEJUICE_PORT
func SendMapleJuiceTCP(destIP string, mjMessage *pb.MapleJuiceMessage) error {
	bytes, err := encodeMapleJuiceMessage(mjMessage)
	if err != nil {
		logger.PrintError(err)
		return err
	}

	conn, err := net.Dial("tcp", destIP+":"+config.MAPLEJUICE_PORT)

	if err != nil {
		logger.PrintError(err)
		return err
	}

	defer conn.Close()

	conn.Write(bytes)

	return nil
}

// This function is called whenever a new TCP connection is established, which reads an SDFS message and calls the given callback
func handleIncomingMapleJuiceMessage(conn net.Conn, callback func(message *pb.MapleJuiceMessage)) error {
	defer conn.Close()
	mjMessage, err := getMapleJuiceMessageFromConn(conn)

	if err != nil {
		logger.PrintError("getMapleJuiceMessageFromConn", err)
		return err
	}

	if mjMessage.Phase == pb.MapleJuicePhase_INIT {
		go callback(mjMessage)
	} else {
		callback(mjMessage)
	}

	return nil
}

func getMapleJuiceMessageFromConn(conn net.Conn) (*pb.MapleJuiceMessage, error) {
	buffer := make([]byte, 0)
	tmp := make([]byte, 1024*1024)
	for {
		n, err := conn.Read(tmp)

		if err != nil {
			if err != io.EOF {
				logger.PrintError("read error:", err)
			}
			break
		}

		buffer = append(buffer, tmp[:n]...)
	}

	mjMessage, err := decodeMapleJuiceMessage(buffer)
	if err != nil {
		return nil, err
	}

	return mjMessage, nil
}

// Decode a byte array into a MapleJuiceMessage message
func decodeMapleJuiceMessage(message []byte) (*pb.MapleJuiceMessage, error) {
	incomingMessage := &pb.MapleJuiceMessage{}
	err := proto.Unmarshal(message, incomingMessage)

	return incomingMessage, err
}

// Encode a MapleJuiceMessage message into a byte array
func encodeMapleJuiceMessage(mjMessage *pb.MapleJuiceMessage) ([]byte, error) {
	message, err := proto.Marshal(mjMessage)
	return message, err
}
