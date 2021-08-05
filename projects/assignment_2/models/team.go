package models


type Team struct {
	TeamId    int           `json:"teamid,omitempty,string"`
	TeamName  string        `json:"team_name,omitempty"`
	ManagerId   int        `json:"managerid,omitempty,string"`
	MemberIds []int `json:"member_ids,omitempty"`
}
