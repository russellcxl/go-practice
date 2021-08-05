package models

type Leave struct {
	LeaveId	   int `json:"leaveid,string"`
	Userid     int `json:"userid,string"`
	StartDay   int `json:"startDay,string"`
	EndDay     int `json:"endDay,string"`
	Days       int `json:"days,string"`
	Status     int `json:"status,string"`
	ApproverId int `json:"approverId,string"`
}
