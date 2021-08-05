package controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/validations"
	"time"
)

func DeleteTeam(teamid int) {

	// check if teamid is valid
	if validations.CheckTeamIdExists(teamid) != true {
		fmt.Println("> Team ID invalid")
		return
	}

	// delete team from Teams DB
	for i := 0 ; i < len(database.Teams) ; i++ {
		if database.Teams[i].TeamId == teamid {
			database.Teams = append(database.Teams[:i], database.Teams[i + 1:]... )
		}
	}

	// delete teamid for all members
	for i := 0 ; i < len(database.Users) ; i++ {
		if database.Users[i].TeamId == teamid {
			database.Users[i].TeamId = 0
		}
	}

	fmt.Printf("> Success!")
	time.Sleep(time.Second)
}
