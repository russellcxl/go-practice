package validations

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/models"
	"regexp"
	"strconv"
	"time"
)

func CheckUserIdIsValid(userid int) bool {
	for _, row := range database.Users {
		if row.UserId == userid {
			return true
		}
	}
	return false
}

func CheckPasswordMatch(password string) (isMatch bool) {
	if database.CurrentUser.Password == password {
		isMatch = true
	}
	return
}

// CheckDateFormat checks that date matches dd/mm/yy format; can only be this year
func CheckDateFormat(date string) bool {
	matched, _ := regexp.MatchString(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[012])/[0-9]{2}$`, date)
	yy, _ := strconv.Atoi(date[len(date) - 2:])
	if yy != (time.Now().Year() % 2000) {
		matched = false
	}
	return matched
}


// CheckStartDate checks that start date is later than now
func CheckStartDate(date string) (isLater bool) {
	d, _ := time.Parse("02/01/06", date)
	if d.Unix() > time.Now().Unix() {
		isLater = true
	}
	return
}

// CheckEndDate checks that start is earlier than end
func CheckEndDate(start, end string) (isLater bool) {
	s, _ := time.Parse("02/01/06", start)
	e, _ := time.Parse("02/01/06", end)
	if s.Unix() < e.Unix() {
		isLater = true
	}
	return
}

func CheckSufficientLeavesForCurrentUser(start, end string) (bool,int) {
	userLeaveBalance := database.CurrentUser.LeaveBalance

	s, _ := time.Parse("02/01/06", start)
	e, _ := time.Parse("02/01/06", end)
	startDay := s.Truncate(time.Hour * 24)
	endDay := e.Truncate(time.Hour * 24)
	difference := int(endDay.Sub(startDay).Hours() / 24)

	if userLeaveBalance >= difference {
		return true, difference
	}
	return false, difference
}

// CheckNotPeakPeriod true means leaves DO NOT overlap with peak period
func CheckNotPeakPeriod(start, end string) bool {
	s, _ := time.Parse("02/01/06", start)
	e, _ := time.Parse("02/01/06", end)

	sMonth := int(s.Month())
	eMonth := int(e.Month())

	if eMonth >= 9 {
		if sMonth <= 11 {
			return false
		}
	}

	return true
}

// CheckTeamLeaveStatus true means <= 50% of team members are taking leave during that period
func CheckTeamLeaveStatus(start, end string, memberIds []int) bool {

	// get all team members' userid
	// check if members are taking leave during that period
	// if > 50%, return false
	teamHeadcount := len(memberIds)

	numberOfOverlaps := 0

	for _, id := range memberIds {
		if CheckUserTakingLeaveByPeriod(start, end, id) {
			fmt.Println("DEBUG")
			numberOfOverlaps++
		}
	}

	fmt.Println("DEBUG", numberOfOverlaps, teamHeadcount)
	if numberOfOverlaps >= teamHeadcount / 2 {
		return false
	}

	return true
}

// CheckUserTakingLeaveByPeriod checks Leaves DB if specified users are taking leave during specified period
func CheckUserTakingLeaveByPeriod(start, end string, userid int) bool {
	s, _ := time.Parse("02/01/06", start)
	e, _ := time.Parse("02/01/06", end)
	sUnix := int(s.Unix())
	eUnix := int(e.Unix())

	var userLeaves []models.Leave

	// get all of user's leaves
	for _, l := range database.Leaves {
		if l.Userid == userid {
			userLeaves = append(userLeaves, l)
		}
	}

	// check if any of the user's approved leaves overlap with given duration
	for _, l := range userLeaves {
		if l.Status == 1 {
			if l.StartDay <= sUnix {
				if l.EndDay >= eUnix {
					return true
				}
			}
		}
	}

	return false
}


func CheckApproveRejectLeaveResp(input string) bool {
	matched, _ := regexp.MatchString(`^[0-9]+,(approve|reject)$`, input)

	if matched == false {
		fmt.Println("> Invalid input format. Example: 1001,approve")
		time.Sleep(time.Second)
		return false
	}

	return true

}

func CheckLeaveIdExists (leaveid int) (bool, *models.Leave) {

	for i := 0 ; i < len(database.Leaves) ; i++ {
		if database.Leaves[i].LeaveId == leaveid {
			return true, &database.Leaves[i]
		}
	}
	return false, nil
}

func CheckTeamIdExists(teamid int) bool {

	for _, t := range database.Teams {
		if t.TeamId == teamid {
			return true
		}
	}
	return false
}


func CheckUserHasTeam(userid int) bool {
	for _, u := range database.Users {
		if u.UserId == userid {
			if u.TeamId > 0 {
				return true
			}
		}
	}
	return false
}
