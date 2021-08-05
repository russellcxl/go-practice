package controller

import (
	"fmt"
	"git.garena.com/russell.chanxl/be-class/assignment_2/constants"
	"git.garena.com/russell.chanxl/be-class/assignment_2/database"
	"git.garena.com/russell.chanxl/be-class/assignment_2/validations"
	"time"
)

func SetAppliedLeaveStatus(leaveid int, action string) {

	// check that leaveid exists
	doesExist, leave := validations.CheckLeaveIdExists(leaveid)
	if doesExist != true {
		fmt.Println("> Leave ID does not exist")
		time.Sleep(time.Second)
		return
	}

	// check that leave has the correct approver/manager
	if leave.ApproverId != database.CurrentUser.UserId {
		fmt.Println("> You are not allowed to approve/reject this leave")
		time.Sleep(time.Second)
		return
	}

	// cannot edit rejected leaves
	if leave.Status == constants.LEAVE_REJECTED {
		fmt.Println("> Cannot edit rejected leaves")
		time.Sleep(time.Second)
		return
	}

	// if the leave has been rejected, will need to credit the leave back to user's account
	if action == "approve" {
		leave.Status = constants.LEAVE_APPROVED
	} else {
		leave.Status = constants.LEAVE_REJECTED
		SetUserLeaveBalance(leave.Userid, leave.Days, constants.LEAVE_ADD)
	}

	fmt.Println("> Successfully modified leave application")
	time.Sleep(time.Second)


}
