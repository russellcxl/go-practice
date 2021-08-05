package controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/models"
	"github.com/russellcxl/go-practice/assignment_2/validations"
	"time"
)


// CreateNewUser all users will be created with password 123 as default
func CreateNewUser(name, role string, teamid, leaveBalance int) {


	if validations.CheckTeamIdExists(teamid) != true {
		fmt.Println("> Team ID does not exist")
		return
	}

	latestUserId := database.Users[len(database.Users) - 1].UserId

	newUser := models.User{
		UserId:       latestUserId + 1,
		Password:     "123",
		Name:         name,
		TeamId:       teamid,
		Role:         role,
		LeaveBalance: leaveBalance,
	}

	// add new user to Users
	database.Users = append(database.Users, newUser)

	// add new user to respective team
	// cannot use for/range because it makes a copy
	for i := 0 ; i < len(database.Teams) ; i++ {
		if database.Teams[i].TeamId == newUser.TeamId {
			database.Teams[i].MemberIds = append(database.Teams[i].MemberIds, newUser.UserId)
		}
	}

	fmt.Printf("> Succesfully created new user")
	time.Sleep(time.Second)

}
