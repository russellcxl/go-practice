package controller

import "github.com/russellcxl/go-practice/assignment_2/database"

func GetUsersManagerId(userid int) (managerId int) {
	teamid := database.CurrentUser.TeamId
	for _, t := range database.Teams {
		if t.TeamId == teamid {
			managerId = t.ManagerId
		}
	}
	return
}
