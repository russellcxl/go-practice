package controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/constants"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/models"
	"github.com/russellcxl/go-practice/assignment_2/validations"
	"sync"
	"time"
)

// CreateLeaveApplication creates leave application, deducts from user leave balance
func CreateLeaveApplication(start, end string, wg *sync.WaitGroup, mu *sync.Mutex) {

	// reduce waitgroup counter
	defer wg.Done()

	// lock the operation so only one go routine can access at a time
	mu.Lock()

	// check if user has enough leave days to take
	allowed, daysTaken := validations.CheckSufficientLeavesForCurrentUser(start, end)
	if allowed == false {
		fmt.Printf("> Insufficient leaves. You have %d days, but wanted to take %d days",database.CurrentUser.LeaveBalance, daysTaken)
		return
	}

	// check if user leave days fall within peak period
	if validations.CheckNotPeakPeriod(start, end) != true {
		fmt.Println("> Cannot take leaves during peak period (Sep - Nov)")
		return
	}

	// get user's member ids
	memberIds := GetMemberIdsByTeamId(database.CurrentUser.TeamId)

	// check if more than 50% of team members have taken leave in the same period
	if validations.CheckTeamLeaveStatus(start, end, memberIds) != true {
		fmt.Println("> 50% or more of your teammates have taken leave within your specified period")
		return
	}

	s, _ := time.Parse("02/01/06", start)
	e, _ := time.Parse("02/01/06", end)
	sUnix := int(s.Unix())
	eUnix := int(e.Unix())

	// get user's manager id
	manager := GetUsersManagerId(database.CurrentUser.UserId)

	// get latest leaveid, so can do 'autoincrement'
	latestLeaveId := database.Leaves[len(database.Leaves) - 1].LeaveId

	var leave = models.Leave{
		LeaveId:    latestLeaveId + 1, // this is clunky
		Userid:     database.CurrentUser.UserId,
		StartDay:   sUnix,
		EndDay:     eUnix,
		Days:		daysTaken,
		Status:		constants.LEAVE_PENDING,
		ApproverId: manager,
	}

	database.Leaves = append(database.Leaves, leave)

	database.CurrentUser.LeaveBalance -= daysTaken

	fmt.Printf("> Successfully created leave from %s to %s\n", start, end)
	fmt.Printf("> You now have %d days of leave remaining\n", database.CurrentUser.LeaveBalance)
	fmt.Println("")

	// unlock the operation
	mu.Unlock()

}
