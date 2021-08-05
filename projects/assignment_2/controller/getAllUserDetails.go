package controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/jedib0t/go-pretty/table"
)

func GetAllUserDetails() {
	t := table.NewWriter()

	t.AppendHeader(table.Row{"#", "User ID", "Name", "Team ID", "Leave balance"})

	for i, user := range database.Users {
		t.AppendRow(table.Row{i + 1, user.UserId, user.Name, user.TeamId, user.LeaveBalance})
	}

	fmt.Println(t.Render())
}
