package client

import (
	"bufio"
	"fmt"
	"github.com/russellcxl/go-practice/assignment_3/client/client_controller"
	pb "github.com/russellcxl/go-practice/assignment_3/protos"
	tcp "github.com/russellcxl/go-practice/assignment_3/tcp_manager"
	"github.com/jedib0t/go-pretty/table"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Client struct {
	port             string
	conn             net.Conn
	user             *pb.User
}

func NewClient(port string) *Client {
	return &Client{port: port}
}

func (c *Client) Run() {
	conn, err := net.Dial("tcp", "localhost"+c.port)
	if err != nil {
		fmt.Printf("Failed to dial into PORT%s: %v", c.port, err)
	}

	c.conn = conn

	defer conn.Close()

	addr := conn.LocalAddr().String()
	fmt.Println("You're connected on", addr)

	interruptSignals := make(chan os.Signal)
	signal.Notify(interruptSignals, syscall.SIGINT, syscall.SIGTERM)

	// alerts the server when the client disconnects unexpectedly
	go func() {
		<-interruptSignals
		fmt.Println("\nSystem interrupt. Closing the connection...")
		tcp.SendMessage(conn, pb.Command_CLIENT_EXIT, nil, pb.Error_SUCCESS)
		conn.Close()
		os.Exit(0)
	}()

	for {
		c.StartPrompt()
	}
}

func (c *Client) GetConn() net.Conn {
	return c.conn
}

func (c *Client) SetCurrentUser(user *pb.User) {
	c.user = user
}

func (c *Client) GetCurrentUser() *pb.User {
	return c.user
}

//******************* PROMPTS *******************//

func (c *Client) StartPrompt() {

	for {
		fmt.Printf("Welcome, please enter [user id] [password] e.g. '1001 123'\n")
		input := getInput()

		user, isSuccess := client_controller.CheckUserLoginDetails(c.GetConn(), input)
		if isSuccess == false {
			fmt.Printf("Invalid format OR userID/password combination\n")
			fmt.Println("\n**************************************\n ")
			continue
		}
		fmt.Println("Login success!")
		fmt.Println("\n**************************************\n ")
		c.SetCurrentUser(user)
		break
	}
	c.staffMenu()
}

func (c *Client) staffMenu() {
	for {
		fmt.Printf("What would you like to do, %s?\n", c.GetCurrentUser().GetName())
		fmt.Println("(1) View leave balance")
		fmt.Println("(2) Apply for leave")
		fmt.Println("(3) Check leave status")
		fmt.Println("(4) Logout")

		input := getInput()
		switch input {
		case "1":
			c.getLeaveBalance()
		case "2":
			c.applyLeavePrompt()
		case "3":
			c.checkLeaveStatus()
		case "4":
			fmt.Println("Logging out... Goodbye,", c.user.GetName())
			fmt.Println("\n**************************************\n ")
			c.user = nil
			c.StartPrompt()
		default:
			fmt.Println("Invalid command, please try again")
			fmt.Println("\n**************************************\n ")
		}
	}
}

func (c *Client) getLeaveBalance() {
	fmt.Printf("Leave balance: %v\n", c.GetCurrentUser().GetLeaveBalance())
	fmt.Println("\n**************************************\n ")
}

func (c *Client) applyLeavePrompt() {
	for {
		fmt.Println("Enter start date of leave in the format dd/mm/yy")
		startDate := getInput()
		fmt.Println("Enter end date of leave in the format dd/mm/yy")
		endDate := getInput()

		isSuccess, err := client_controller.ApplyLeave(c.GetConn(), c.GetCurrentUser(), startDate, endDate)
		if isSuccess == false {
			fmt.Println(err)
			fmt.Println("\n**************************************\n ")
			break
		}
		fmt.Printf("Successfully applied for leave. You have %v days left\n", c.GetCurrentUser().GetLeaveBalance())
		fmt.Println("\n**************************************\n ")
		break
	}
}

func (c *Client) checkLeaveStatus() {

	for {
		// initialise console table
		t := table.NewWriter()

		// append headers
		t.AppendHeader(table.Row{"#", "leaveid", "start", "end", "total days", "status"})

		leaves, err := client_controller.GetLeaveApplications(c.GetConn(), c.GetCurrentUser().GetUserId())
		if err != "" {
			fmt.Println(err)
			return
		}

		for i, l := range leaves {
			// convert unix to readable
			s := time.Unix(int64(l.GetStartTime()), 0)
			e := time.Unix(int64(l.GetEndTime()), 0)

			t.AppendRow(table.Row{i + 1, l.GetLeaveId(), s.Format("02 Jan 2006"), e.Format("02 Jan 2006"), l.GetDaysTaken(), l.GetStatus()})
		}
		fmt.Println(t.Render())

		fmt.Println("Enter the leave ID of the application you want to delete or hit ENTER to return to the main menu")

		input := getInput()
		switch input {
		case "":
			fmt.Println("\n**************************************\n ")
			return
		default:
			var validId bool
			for _, l := range leaves {
				if input == strconv.Itoa(int(l.GetLeaveId())) {
					validId = true
				}
			}
			if !validId {
				fmt.Println("Invalid leave ID entered. Please try again")
				fmt.Println("\n**************************************\n ")
				continue
			}
			client_controller.DeleteLeaveApplication(c.GetConn(), c.GetCurrentUser(), input)
		}
	}
}

//******************* SCANNER *******************//

func getInput() string {
	fmt.Printf("\n> ")

	// init scanner
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

	// scan command line for input
	scanner.Scan()

	fmt.Println("")
	return scanner.Text()
}
