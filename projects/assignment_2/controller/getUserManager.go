package controller

import "git.garena.com/russell.chanxl/be-class/assignment_2/database"

func GetUsersManagerId(userid int) (managerId int) {
	teamid := database.CurrentUser.TeamId
	for _, t := range database.Teams {
		if t.TeamId == teamid {
			managerId = t.ManagerId
		}
	}
	return
}
