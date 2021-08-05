package controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/models"
	"time"
)

func DeleteAppliedLeave(leaveId int) {

	// get details of the leave application
	var leaveApplication models.Leave
	var idx int
	for i, l := range database.Leaves {
		if leaveId == l.LeaveId {
			leaveApplication = l
			idx = i
		}
	}

	// check if leave id exists
	if leaveApplication.LeaveId == 0 {
		fmt.Println("> Invalid leave ID")
		return
	}

	// check if leave id belongs to current user
	if leaveApplication.Userid != database.CurrentUser.UserId {
		fmt.Println("> Invalid leave ID")
		return
	}

	// check if leave is in pending status
	if leaveApplication.Status != 0 {
		fmt.Println("> Leave has already been approved/rejected")
		return
	}

	// splice leave application from database
	database.Leaves = append(database.Leaves[:idx], database.Leaves[idx + 1:]...)
	for i := 2 ; i > 0 ; i-- {
		fmt.Printf("> Deleting leave... %ds remaining\n", i)
		time.Sleep(time.Second)
	}
	fmt.Println("> Successfully deleted your leave application, ID:", leaveId)

	// credit leave to user
	database.CurrentUser.LeaveBalance += leaveApplication.Days

}
