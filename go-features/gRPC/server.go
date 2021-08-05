package main

import (
	"git.garena.com/russell.chanxl/personal/gRPC/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	PORT := "9000"
	lis, err := net.Listen("tcp", ":" + PORT)
	if err != nil {
		log.Fatalf("failed to listen on 9000: %v", err)
	}

	// initialise gRPC server
	grpcServer := grpc.NewServer()

	// create shell for your service
	s := chat.Server{}

	// register chat service on gRPC server
	// what does it mean to register a service?
	chat.RegisterChatServiceServer(grpcServer, &s)

	// accepts incoming connections on the listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}

}
