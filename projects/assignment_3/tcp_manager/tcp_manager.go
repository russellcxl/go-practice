package tcp_manager

import (
	"bufio"
	"encoding/binary"
	pb "git.garena.com/russell.chanxl/be-class/assignment_3/protos"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

var defaultHeaderLength = 8

func ReadMessage(conn net.Conn) *pb.Message {

	// get header (length of message)
	header := make([]byte, defaultHeaderLength)
	reader := bufio.NewReader(conn)
	reader.Read(header)
	byteLength := binary.LittleEndian.Uint64(header)

	// get actual message
	resBytes := make([]byte, byteLength)
	_, err := reader.Read(resBytes)
	if err != nil {
		log.Printf("Failed to read from TCP stream: %v\n", err)
	}

	// unmarshall proto message
	var msg pb.Message
	err = proto.Unmarshal(resBytes, &msg)
	if err != nil {
		log.Printf("Failed to unmarshal message: %v\n", err)
	}

	return &msg

}

func SendMessage(conn net.Conn, cmd pb.Command, data []byte, error pb.Error) {

	req := pb.Message{
		Command: cmd,
		Data:    data,
		Error:   error,
	}

	reqMarshalled, err := proto.Marshal(&req)
	if err != nil {
		log.Println("Failed to marshal message data |", err)
	}

	reqWithHeader := make([]byte, defaultHeaderLength)
	binary.LittleEndian.PutUint64(reqWithHeader, uint64(len(reqMarshalled)))
	reqWithHeader = append(reqWithHeader, reqMarshalled...)

	_, err = conn.Write(reqWithHeader)
	if err != nil {
		log.Println("Failed to write message to TCP stream |", err)
	}
}