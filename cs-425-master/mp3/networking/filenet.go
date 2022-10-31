package networking

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"net"
	"syscall"

	pb "../ProtocolBuffers/ProtoPackage"
	"../config"
	"../logger"
	proto "github.com/golang/protobuf/proto"
)

// Encode an SDFS message into a byte array
func EncodeSdfsMessage(serviceMessage *pb.SdfsMessage) ([]byte, error) {
	// logger.PrintInfo("outgoing SDFS Message Fields:", serviceMessage.ReturnString, serviceMessage.FileSeqNum, serviceMessage.FileOp, serviceMessage.SdfsFilename)
	message, err := proto.Marshal(serviceMessage)

	return message, err
}

// Decode a byte array into an SDFS message
func DecodeSdfsMessage(message []byte) (*pb.SdfsMessage, error) {
	incomingMessage := &pb.SdfsMessage{}
	err := proto.Unmarshal(message, incomingMessage)
	// logger.PrintInfo("incoming SDFS Message Fields:", incomingMessage.ReturnString, incomingMessage.FileSeqNum, incomingMessage.FileOp, incomingMessage.SdfsFilename)

	return incomingMessage, err
}

// Given an existing TCP connection, send an SDFS message through that connection without waiting for a reply
func SendSdfsMessageAndClose(conn net.Conn, sdfsMessage *pb.SdfsMessage) error {
	message, err := EncodeSdfsMessage(sdfsMessage)

	if err != nil {
		logger.PrintError(err)
		return err
	}

	sendTCP(conn, message)
	return nil
}

// Given an IP address, create a TCP connection and send an SDFS message and wait for a reply SDFS message
func SendSdfsMessageWithReply(dest string, sdfsMessage *pb.SdfsMessage) (*pb.SdfsMessage, net.Conn) {
	message, err := EncodeSdfsMessage(sdfsMessage)
	if err != nil {
		logger.PrintError(err)
		return nil, nil
	}

	replyMessage, conn := sendTCPWithReply(dest, message)
	return replyMessage, conn
}

// Given an IP address, create a TCP connection and send a one-off SDFS message without waiting for a reply
func SendSoloSdfsMessage(dest string, sdfsMessage *pb.SdfsMessage) bool {
	message, err := EncodeSdfsMessage(sdfsMessage)

	if err != nil {
		logger.PrintError(err)
		return false
	}

	conn, err := net.Dial("tcp", dest+":"+config.TCP_PORT)
	if err != nil {
		logger.PrintError(err)
		return false
	}

	sendTCP(conn, message)
	return true
}

// Given a connection, read an SDFS message from that connection and decode it
func GetSdfsMessageFromConn(conn net.Conn) (*pb.SdfsMessage, error) {
	numBytes := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	mem := make([]byte, 0)
	messageSize := 0
	hasReadSize := false

	for {
		bufSize := math.Min(1024, float64(messageSize))

		if !hasReadSize {
			bufSize = 8
		}

		buf := make([]byte, int64(bufSize))

		n, err := conn.Read(buf)

		if err == io.EOF || n == 0 {
			break
		}

		if err != nil {
			logger.PrintError(err)
			return nil, err
		}

		startIndex := 0

		if messageSize == 0 {
			numBytes[0] = buf[0]
			numBytes[1] = buf[1]
			numBytes[2] = buf[2]
			numBytes[3] = buf[3]
			numBytes[4] = buf[4]
			numBytes[5] = buf[5]
			numBytes[6] = buf[6]
			numBytes[7] = buf[7]
			startIndex = 8

			messageSize = int(binary.BigEndian.Uint64(numBytes))
			hasReadSize = true
		}

		messageSize -= (n - startIndex)
		mem = append(mem, buf[startIndex:n]...)

		if messageSize <= 0 {
			break
		}
	}

	if len(mem) == 0 {
		logger.PrintError("read 0 bytes from conn -- connection likely closed on the other end")
		return nil, errors.New("Connection close")
	}

	sdfsMessage, err := DecodeSdfsMessage(mem)
	if err != nil {
		logger.PrintError(err)
		return nil, err
	}

	return sdfsMessage, nil
}

// This function is called whenever a new TCP connection is established, which reads an SDFS message and calls the given callback
func handleIncomingSdfsMessage(conn net.Conn, callback func(conn net.Conn, message *pb.SdfsMessage)) error {
	defer conn.Close()
	// logger.PrintInfo("handleIncomingSdfsMessage")
	sdfsMessage, err := GetSdfsMessageFromConn(conn)

	if err != nil {
		logger.PrintError(err)
		return err
	}

	callback(conn, sdfsMessage)

	return nil
}

// Checks whether a connection is closed by using syscalls and MSG_PEEK
func IsConnClosed(conn net.Conn) bool {
	rc, err := conn.(syscall.Conn).SyscallConn()
	if err != nil {
		return false
	}

	isConnClosed := false

	err = rc.Read(func(fd uintptr) bool {
		n, _, err := syscall.Recvfrom(int(fd), []byte{0}, syscall.MSG_PEEK|syscall.MSG_DONTWAIT)
		if n == 0 && err == nil {
			isConnClosed = true
		}
		return true
	})

	return isConnClosed && err != nil
}
