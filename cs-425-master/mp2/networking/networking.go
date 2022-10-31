package networking

import (
	"encoding/binary"
	"net"

	pb "../ProtocolBuffers/ProtoPackage"
	"../config"
	"../logger"
	"../utils"
	"github.com/jinzhu/copier"
)

// Get this machine's IP address
func GetLocalIPAddr() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.PrintError("net.Dial")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// Send a Membership message to the given IP addresses
func SendAll(destinations []string, serviceMessage *pb.MembershipMessage) error {
	masterMessage, err := EncodeMembershipMessage(serviceMessage)
	if err != nil {
		return err
	}

	memberListCpy := pb.MembershipMessage{}
	copier.Copy(&memberListCpy, &serviceMessage)
	memberListCpy.FileMap = nil
	message, err := EncodeMembershipMessage(&memberListCpy)
	if err != nil {
		return err
	}

	for _, dest := range destinations {
		messageToSend := message

		if utils.GetIPFromID(serviceMessage.MasterID) == dest {
			messageToSend = masterMessage
		}

		err := Send(dest, messageToSend)
		if err != nil {
			return err
		}
	}

	return nil
}

// Send a byte array to the given IP address
func Send(dest string, message []byte) error {
	if len(message) > config.BUFFER_SIZE {
		logger.WarningLogger.Println("Send: message is larger than BUFFER_SIZE")
	}

	addr, err := net.ResolveUDPAddr("udp", dest+":"+config.PORT)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Write(message)
	if err != nil {
		return err
	}

	return nil
}

// Set up a UDP server on the given port, which calls the given callback function when it receives a new message
func Listen(port string, callback func(message []byte) error) error {
	port = ":" + port

	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	defer conn.Close()
	buffer := make([]byte, config.BUFFER_SIZE)

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			return err
		}

		callback(buffer[0:n])
	}
}

// Send the given byte array over the given TCP connection
func sendTCP(conn net.Conn, message []byte) error {
	numBytes := make([]byte, 8)

	messageSize := len(message)
	binary.BigEndian.PutUint64(numBytes, uint64(messageSize))

	conn.Write(numBytes)
	conn.Write(message)

	return nil
}

// Create a TCP connection with the given IP address and send a byte array over that connection
// Waits for a reply message from the other side
func sendTCPWithReply(dest string, message []byte) (*pb.SdfsMessage, net.Conn) {
	conn, err := net.Dial("tcp", dest+":"+config.TCP_PORT)
	if err != nil {
		logger.PrintError(err)
		return nil, nil
	}

	sendTCP(conn, message)

	replyMessage, err := GetSdfsMessageFromConn(conn)

	if err != nil {
		logger.PrintError("error in getSdfsMessageFromConn")
	}

	return replyMessage, conn
}

// Set up a TCP server on the given port, which calls the given callback function when it receives a new message
func ListenTCP(callback func(conn net.Conn, message *pb.SdfsMessage)) error {
	listener, err := net.Listen("tcp", ":"+config.TCP_PORT)
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

		go handleIncomingSdfsMessage(conn, callback)
	}
}
