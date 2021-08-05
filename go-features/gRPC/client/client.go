package main

import (
	"context"
	"fmt"
	"git.garena.com/russell.chanxl/personal/gRPC/chat"
	"google.golang.org/grpc"
	"log"
)

func main() {

	// dial into the port which gRPC is on
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	// initialise chat service
	c := chat.NewChatServiceClient(conn)

	// message using protobuf generated
	message := chat.Message{Body: "Hello from the client"}

	// call a chat service function
	// this sends a message to the server and gets a response; defined in chat.go
	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	fmt.Printf("Response from server: %s", res.Body)

}
