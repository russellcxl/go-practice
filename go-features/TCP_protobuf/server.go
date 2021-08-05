package main

import (
	"bufio"
	"fmt"
	pb "git.garena.com/russell.chanxl/personal/TCP_protobuf/proto"
	"github.com/golang/protobuf/proto"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {

	// gets port number from command line
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	// parses port number
	PORT := ":" + arguments[1]

	// starts server
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Printf("Could not listen on %s due to:\n =====> %v\n", PORT, err)
		return
	}

	defer lis.Close()

	// saw this somewhere but don't know what it is
	rand.Seed(time.Now().Unix())

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln("Could not accept connection:", err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	// will print address every time you make a new connection
	fmt.Printf("Now serving %s\n", conn.RemoteAddr().String())

	defer conn.Close()

	clientReader := bufio.NewReader(conn)

	for {

		// read message from client
		//bs := make([]byte, 128)
		receivedBytes, err := clientReader.ReadBytes(116)
		if err != nil {
			log.Println("Failed to read message from client:\n===>", err)
			break
		}

		log.Printf("Received bytes: %v\n", receivedBytes)

		req := new(pb.RequestMessage)
		if err = proto.Unmarshal(receivedBytes, req); err != nil {
			log.Println("Failed to unmarshal message")
			break
		}


		log.Printf("-> %s", req)

	}

	fmt.Println("No messages received")

}



