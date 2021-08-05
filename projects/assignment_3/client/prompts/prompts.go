package prompts

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//func LoginPrompt(cli *client.Client) {
//	for {
//		printLine("Welcome, please enter [user id] [password] e.g. '1001 123'")
//		input := getInput()
//		if util.ValidateLoginFormat(input) == false {
//			printLine("Wrong format. Please follow [user id] [password]")
//			continue
//		}
//		if client_controller.CheckUserLoginDetails(cli, input) == false {
//			printLine("Invalid user id / password combination")
//			continue
//		}
//	}
//	//staffMenu()
//}

//func staffMenu() {
//
//	for {
//		fmt.Printf("> Hello %s, what would you like to do?\n")
//		fmt.Println("(1) View leave balance")
//		fmt.Println("(2) Apply for leave")
//		fmt.Println("(3) Check leave status")
//		fmt.Println("(4) Logout")
//		fmt.Println(separator)
//
//		input := getInput()
//		switch input {
//		case "1":
//			userId, err := strconv.Atoi(input)
//			if err != nil {
//				fmt.Println("Invalid userId format (only enter digits) |", err)
//			}
//			client_controller.GetLeaveBalance(uint64(userId))
//		case "2": fmt.Println()
//		case "3": fmt.Println()
//		case "4": fmt.Println()
//		default:
//			fmt.Println("> Invalid command. Try again")
//			fmt.Println(separator)
//		}
//	}
//}

//================ scanner ==================//

func getInput() string {
	fmt.Println("")

	// init scanner
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

	// scan command line for input
	scanner.Scan()

	fmt.Println("\n--------------------------------------------\n ")
	if scanner.Text() == "exit" {
		for i := 3; i > 0; i-- {
			fmt.Printf("> Saving changes... %ds remaining\n", i)
			time.Sleep(time.Second)
		}
		fmt.Println("> All changes saved! Goodbye")
		os.Exit(0)
	}

	return scanner.Text()

}

func printLine(input string) {
	fmt.Printf("> %s", input)
}
