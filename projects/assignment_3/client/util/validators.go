package util

import (
	"regexp"
	"strconv"
	"time"
)

func ValidateLoginFormat(credentials string) bool {
	matched, _ := regexp.MatchString(`^[0-9]+ .+$`, credentials)
	return matched
}

// CheckDateFormat checks that date matches dd/mm/yy format; can only be this year
func CheckDateFormat(date string) bool {
	matched, _ := regexp.MatchString(`^(0[1-9]|[12][0-9]|3[01])/(0[1-9]|1[012])/[0-9]{2}$`, date)
	if !matched {
		return false
	}
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

func CheckSufficientLeavesForCurrentUser(start, end string, userLeaveBalance int) (bool,int) {
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

