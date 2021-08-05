package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/russellcxl/go-practice/gRPC_CRUD/ideas"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {

	// dials in to port which grpc is serving
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	// create a new client service
	cli := ideas.NewIdeaServiceClient(conn)

	user := &ideas.User{Id: 1235}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		switch scanner.Text() {
		case "get":
			res, err := cli.GetIdeas(context.Background(), user)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Resp from server:", res.GetIdeas())
		case "exit":
			fmt.Printf("Exiting\n")
			os.Exit(0)
		default:
			fmt.Println("Invalid")
			continue
		}
	}

}
