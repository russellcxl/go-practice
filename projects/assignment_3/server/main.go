package main

import (
	"git.garena.com/russell.chanxl/be-class/assignment_3/server/cache"
	"git.garena.com/russell.chanxl/be-class/assignment_3/server/server"
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
