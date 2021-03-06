package main

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/prompts"
	"os"
	"os/signal"
)

func main() {

	// open DB
	database.InitDb()

	// divert any system interrupts to the signals channel and makes a graceful exit instead
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<- signals
		fmt.Println("\nExiting program without saving changes")
		os.Exit(0)
	}()

	// start program
	prompts.LoginPrompt()

}
