package main

import (
	"fmt"
	"github.com/russellcxl/go-practice/gRPC_CRUD/ideas"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// handle interrupts
	exit := make(chan os.Signal, 2)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	// start tcp listener on 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}

	defer lis.Close()

	// initialise grpc server
	grpcServer := grpc.NewServer()

	// initialise service shell
	ideaService := &ideas.Service{}

	// register service on grpc server
	ideas.RegisterIdeaServiceServer(grpcServer, ideaService)

	// ?
	reflection.Register(grpcServer)

	fmt.Println("gRPC server listening on 9000")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()

	<-exit
	fmt.Println("\nStopping gRPC server..")
	grpcServer.GracefulStop()


}
