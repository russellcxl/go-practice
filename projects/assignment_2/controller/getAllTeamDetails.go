package controller

import (
	"fmt"
	"git.garena.com/russell.chanxl/be-class/assignment_2/database"
	"github.com/jedib0t/go-pretty/table"
)

func GetAllTeamDetails() {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"#", "Team ID", "Name", "Manager ID", "Member IDs"})

	for i, team :=  range database.Teams {
		t.AppendRow(table.Row{i + 1,team.TeamId, team.TeamName, team.ManagerId, team.MemberIds})
	}

	fmt.Println(t.Render())
}
