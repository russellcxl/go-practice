package controller

import (
	"git.garena.com/russell.chanxl/be-class/assignment_2/constants"
	"git.garena.com/russell.chanxl/be-class/assignment_2/database"
)

// SetUserLeaveBalance action is 0 for approve, 1 for reject; check constants
func SetUserLeaveBalance(userid, leaveAmt int, action int) {
	for i := 0 ; i < len(database.Users) ; i++ {

		if database.Users[i].UserId == userid {
			if action == constants.LEAVE_ADD {
				database.Users[i].LeaveBalance += leaveAmt
			} else {
				database.Users[i].LeaveBalance -= leaveAmt
			}
		}
	}
}

