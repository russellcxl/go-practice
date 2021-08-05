package controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/models"
	"github.com/russellcxl/go-practice/assignment_2/validations"
	"time"
)

func CreateNewTeam(teamName string, managerId int) {

	if validations.CheckUserIdIsValid(managerId) != true {
		fmt.Println("> User does not exist")
		return
	}

	if validations.CheckUserHasTeam(managerId) == true {
		fmt.Println("> This user already has an assigned team")
		return
	}

	latestTeamId := database.Teams[len(database.Teams) - 1].TeamId

	var newTeam = models.Team{
		TeamId:    latestTeamId + 100,
		TeamName:  teamName,
		ManagerId: managerId,
		MemberIds: []int{managerId},
	}

	database.Teams = append(database.Teams, newTeam)

	fmt.Println("> Success!")
	time.Sleep(time.Second)
}
