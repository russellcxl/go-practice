package controller

import (
	"fmt"
	"git.garena.com/russell.chanxl/be-class/assignment_2/database"
	"git.garena.com/russell.chanxl/be-class/assignment_2/models"
	"github.com/jedib0t/go-pretty/table"
	"time"
)

// GetAppliedLeavesByTeam prints table and returns []Leave
func GetAppliedLeavesByTeam(teamid int) []models.Leave {
	var leaves []models.Leave

	// get approver id using teamid
	var approverid int
	for _, team := range database.Teams {
		if team.TeamId == teamid {
			approverid = team.ManagerId
		}
	}

	// get all leaves that belong to team (by using the approver id)
	// this, however, assumes that each team only has one approver >:(
	for idx, _ := range database.Leaves {
		if database.Leaves[idx].ApproverId == approverid {
			leaves = append(leaves, database.Leaves[idx])
		}
	}

	// initialise console table
	t := table.NewWriter()

	// write headers
	t.AppendHeader(table.Row{"#", "leaveid", "applicant", "start", "end", "total days", "status", "approver"})

	for idx, l := range leaves {

		// convert unix to readable
		s := time.Unix(int64(l.StartDay), 0)
		e := time.Unix(int64(l.EndDay), 0)

		// convert status to readable
		status := "Pending"
		switch l.Status {
		case 1: status = "Approved"
		case 2: status = "Rejected"
		}

		// convert approverid to approver name
		approver, _ := GetUserDetails(l.ApproverId)

		staff, _ := GetUserDetails(l.Userid)

		t.AppendRow(table.Row{idx + 1, l.LeaveId, staff.Name, s.Format("02 Jan 2006"), e.Format("02 Jan 2006"), l.Days, status, approver.Name})
	}

	fmt.Println(t.Render())

	return leaves
}