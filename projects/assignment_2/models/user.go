package models

type User struct {
	UserId       int    `json:"userid,string"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	TeamId       int    `json:"teamid,string"`
	Role         string `json:"role"`
	LeaveBalance int    `json:"leave_balance,string"`
}
