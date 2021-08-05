package controller

import "git.garena.com/russell.chanxl/be-class/assignment_2/database"

func GetMemberIdsByTeamId(teamid int) (memberIds []int) {
	for _, team := range database.Teams {
		if team.TeamId == teamid {
			memberIds = team.MemberIds
		}
	}
	return
}
