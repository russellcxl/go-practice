package controller

import (
	"git.garena.com/russell.chanxl/be-class/assignment_2/database"
	"git.garena.com/russell.chanxl/be-class/assignment_2/models"
	"github.com/jedib0t/go-pretty/table"
	"time"
)

// GetAppliedLeavesByUser returns []Leave and a table in string format (used for printing)
func GetAppliedLeavesByUser(userid int) ([]models.Leave, string) {

	// filter leaves taken by user
	var leaves []models.Leave
	for _, l := range database.Leaves {
		if l.Userid == userid {
			leaves = append(leaves, l)
		}
	}

	// initialise console table
	t := table.NewWriter()

	// write headers
	t.AppendHeader(table.Row{"#", "leaveid", "start", "end", "total days", "status", "approver"})

	for idx, l := range leaves {

		// convert unix to readable
		s := time.Unix(int64(l.StartDay), 0)
		e := time.Unix(int64(l.EndDay), 0)

		// convert status to readable
		status := "Pending"
		switch l.Status {
		case 2: status = "Approved"
		case 3: status = "Rejected"
		}

		// convert approverid to approver name
		approver, _ := GetUserDetails(l.ApproverId)

		t.AppendRow(table.Row{idx + 1, l.LeaveId, s.Format("02 Jan 2006"), e.Format("02 Jan 2006"), l.Days, status, approver.Name})
	}

	return leaves, t.Render()
}
