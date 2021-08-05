package main

import (
	"github.com/russellcxl/go-practice/assignment_3/server/cache"
	"github.com/russellcxl/go-practice/assignment_3/server/server"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("Please provide a port number")
		return
	}
	PORT := ":" + args[1]

	redisRepo := cache.NewRedisRepo()

	server.NewServer(redisRepo, PORT).Run()
}
