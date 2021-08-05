package main

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_3/client/client"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a port number for the client")
		return
	}
	PORT := ":" + args[1]

	client.NewClient(PORT).Run()
}
