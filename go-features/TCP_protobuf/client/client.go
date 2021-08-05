package main

import (
	"fmt"
	pb "git.garena.com/russell.chanxl/personal/TCP_protobuf/proto"
	"github.com/golang/protobuf/proto"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// get port number
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide a port number")
		return
	}

	// parse port number
	PORT := ":" + args[1]

	// dial into port
	conn, err := net.Dial("tcp", "localhost"+PORT)
	if err != nil {
		fmt.Printf("Failed to dial into PORT%s due to:\n=====> %v", PORT, err)
	}

	// interrupts any cancellation signals and closes the connection properly
	interruptSignals := make(chan os.Signal)
	signal.Notify(interruptSignals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-interruptSignals
		fmt.Println("\nSystem interrupt. Closing the connection...")
		conn.Write([]byte("STOP\n"))
		conn.Close()
		os.Exit(0)
	}()

	defer conn.Close()

	// initialise client reader (from os) and server reader (from connection)
	//clientReader := bufio.NewReader(os.Stdin)
	//serverReader := bufio.NewReader(conn)

	addr := conn.LocalAddr().String()

	fmt.Println("You are on", addr)

	req := &pb.RequestMessage{Msg: "Hello from client"}
	reqBytes, _ := proto.Marshal(req)
	fmt.Printf("Sending message to server: %v\n", reqBytes)

	_, err = conn.Write(reqBytes)
	if err != nil {
		fmt.Println("Failed to send message to server:\n===>", err)
	}

}
