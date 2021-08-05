package prompts

import (
	"bufio"
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/controller"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/validations"
	"github.com/jedib0t/go-pretty/table"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)


func LoginPrompt() {
	fmt.Println("> Enter userid for login")

	t := table.NewWriter()
	t.SetStyle(table.StyleDouble)
	t.SetTitle("Sample accounts")
	t.AppendHeader(table.Row{"Role", "User ID", "Password", "Test actions"})
	t.AppendRows([]table.Row{
		{"Admin", 0, 123},
		{"Staff", 5, 123, "Apply leave from 09/07/21 --> 10/07/21 to check overlap"},
		{"Manager", 3, 123, "View/approve/reject team leaves"},
	})
	fmt.Println(t.Render())

	for {
		res := getResponse()

		userid, err := strconv.Atoi(res)
		if err != nil {
			fmt.Println(err, "Please try again (enter userid)")
			continue
		}

		// if userid is valid, set CurrentUser
		if validations.CheckUserIdIsValid(userid) {
			user, err := controller.GetUserDetails(userid)
			if err != nil {
				fmt.Println("> Unable to retrieve user info:", err)
			}
			database.CurrentUser = user
			break
		}

		fmt.Println("> Invalid userid. Try again")
	}
	passwordPrompt()
}

func passwordPrompt() {
	fmt.Println("> Enter password for userid:", database.CurrentUser.UserId)
	for {
		res := getResponse()

		if validations.CheckPasswordMatch(res) {
			break
		}
		fmt.Println("> Wrong password. Try again")
	}

	switch database.CurrentUser.Role {
	case "staff": staffMenu()
	case "manager": managerMenu()
	case "admin": adminMenu()
	default: log.Fatalln("> Role error. Please check DB for userid:", database.CurrentUser.UserId)
	}
}

func staffMenu() {
	for {
		fmt.Println("> Select a service,", database.CurrentUser.Name)
		fmt.Println("(1) Check leave balance")
		fmt.Println("(2) Apply for leave")
		fmt.Println("(3) Check/delete leave status")
		fmt.Println("(4) Log out")

		res := getResponse()

		switch res {
		case "1": fmt.Println("> You have", database.CurrentUser.LeaveBalance, "days left\n")
		case "2": applyLeavePrompt()
		case "3": checkDeleteLeavePrompt()
		case "4": {
			fmt.Println("> Goodbye")
			LoginPrompt()
		}
		default: fmt.Println("> Invalid command. Try again\n")
		}
	}
}

func applyLeavePrompt() {

	var startDate, endDate string
	var wg sync.WaitGroup
	var mu sync.Mutex

	// get start date
	for {
		fmt.Println("> Key in the start day in the format dd/mm/yy")
		res := getResponse()
		if validations.CheckDateFormat(res) {
			if validations.CheckStartDate(res) {
				startDate = res
				break
			} else {
				fmt.Println("> Start date should be later than today")
				continue
			}
		} else {
			fmt.Println("> Invalid date format")
		}
	}


	// get end date
	for {
		fmt.Println("> Key in end day in the format dd/mm/yy")
		res := getResponse()
		if validations.CheckDateFormat(res) {
			if validations.CheckEndDate(startDate, res) {
				endDate = res
				break
			} else {
				fmt.Println("> End date should be later than start date")
				continue
			}
		} else {
			fmt.Println("> Invalid date format")
		}
	}

	// create leave application once all validations pass
	wg.Add(1)
	go controller.CreateLeaveApplication(startDate, endDate, &wg, &mu)
	wg.Wait()

}

func checkDeleteLeavePrompt() {

	for {
		_, tbl := controller.GetAppliedLeavesByUser(database.CurrentUser.UserId)
		fmt.Println(tbl)
		fmt.Println("> To delete a leave application, enter the leave ID (enter blank to go back to main menu)")

		res := getResponse()
		if res == "" {
			break
		}

		leaveid, err := strconv.Atoi(res)
		if err != nil {
			fmt.Println(err)
		}

		controller.DeleteAppliedLeave(leaveid)
		break
	}

}

func managerMenu() {
	for {
		fmt.Println("> Select a service,", database.CurrentUser.Name)
		fmt.Println("> (1) View/approve/reject team members' leaves")
		fmt.Println("> (2) Log out")

		res := getResponse()

		switch res {
		case "1": viewTeamOverviewPrompt()
		case "2": {
			fmt.Println("> Goodbye")
			LoginPrompt()
		}
		default: fmt.Println("> Invalid command. Try again\n")
		}

	}
}


func viewTeamOverviewPrompt() {
	for {
		controller.GetAppliedLeavesByTeam(database.CurrentUser.TeamId)
		fmt.Println("> To approve/reject a leave, enter leave ID, approve/reject (e.g. `1001,approve`)")

		res := getResponse()
		if res == "" {break}

		// resp format check
		if validations.CheckApproveRejectLeaveResp(res) == false {
			continue
		}

		inputs := strings.Split(res, ",")
		leaveid, _ := strconv.Atoi(inputs[0])
		action := inputs[1]


		controller.SetAppliedLeaveStatus(leaveid, action)

	}
}


func adminMenu() {

	for {

		fmt.Println("> Select a service,", database.CurrentUser.Name)
		fmt.Println("> -------------------")
		fmt.Println("> (1) View all teams")
		fmt.Println("> (2) Add new team")
		fmt.Println("> (3) Remove team")
		fmt.Println("> -------------------")
		fmt.Println("> (4) View all users")
		fmt.Println("> (5) Change user role")
		fmt.Println("> (6) Add new user")
		fmt.Println("> -------------------")
		fmt.Println("> (7) Log out")


		switch getResponse() {
		case "1": controller.GetAllTeamDetails()
		case "2": createTeamPrompt()
		case "3": removeTeamPrompt()
		case "4": controller.GetAllUserDetails()
		case "5": changeUserRolePrompt()
		case "6": createUserPrompt()
		case "7": {
			fmt.Println("> Goodbye")
			LoginPrompt()
		}
		default: {
			fmt.Println("> Invalid action")
			continue
		}
		}
	}
}

func createUserPrompt() {
	for {

		// get name
		fmt.Println("> Enter the name of the user (enter blank to leave)")
		name := getResponse()


		// get team
		fmt.Println("> Enter team id for user")
		teamid, err := strconv.Atoi(getResponse())
		if err != nil {
			fmt.Println("> Invalid team ID:", err)
			continue
		}


		// get leave balance
		fmt.Println("> Enter leave balance for user")
		leaveBalance, err := strconv.Atoi(getResponse())
		if err != nil {
			fmt.Println("> Only enter integers")
			continue
		}


		// get role
		fmt.Println("> Enter user role")
		role := getResponse()
		if role != "staff" && role != "admin" {
			fmt.Println("> Invalid role")
			continue
		}

		controller.CreateNewUser(name, role, teamid, leaveBalance)
		controller.GetAllUserDetails()
		break

	}
}

func changeUserRolePrompt() {
	for {

		// get userid
		fmt.Println("> Enter userid of user whose role you want to change")

		userid, err := strconv.Atoi(getResponse())
		if err != nil {
			fmt.Println("> Invalid input:", err)
			continue
		}


		// get role
		fmt.Println("> Enter new role")

		role := getResponse()
		if role != "staff" && role != "manager" {
			fmt.Println("> Invalid role")
			continue
		}

		controller.SetUserRole(userid, role)
		break

	}
}

func createTeamPrompt() {
	for {
		fmt.Println("> Enter team name")
		teamName := getResponse()

		fmt.Println("> Enter manager id")
		managerId, err := strconv.Atoi(getResponse())
		if err != nil {
			fmt.Println("> Formatting error:", err)
			continue
		}

		controller.CreateNewTeam(teamName, managerId)
		controller.GetAllUserDetails()
		break
	}
}

func removeTeamPrompt() {
	for {
		controller.GetAllTeamDetails()

		fmt.Println("> Enter teamid of team to delete")
		teamid, err := strconv.Atoi(getResponse())
		if err != nil {
			fmt.Println("> Input format error:", err)
		}

		controller.DeleteTeam(teamid)
		controller.GetAllTeamDetails()
		break
	}
}









//========================= get resp ============================//


func getResponse() string {
	fmt.Println("")

	// init scanner
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

	// scan command line for input
	scanner.Scan()

	fmt.Println("----------------------------------")
	if scanner.Text() == "exit" {
		for i := 3 ; i > 0 ; i-- {
			fmt.Printf("> Saving changes... %ds remaining\n", i)
			time.Sleep(time.Second)
		}
		database.WriteToDbs()
		fmt.Println("> All changes saved! Goodbye")
		os.Exit(0)
	}

	return scanner.Text()
}
