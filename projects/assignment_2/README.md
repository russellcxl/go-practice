# Assignment 2 -- HR system

## Requirements
- Must have 3 user roles -- staff, manager, admin
    - **STAFF** 
      - check leave applications
      - apply leave
      - cancel leave application
    - **MANAGER**
      - view all leave applications made by team
      - approve / reject leave applications
    - **ADMIN** 
      - view all teams
      - add new team
      - remove team
      - view all users
      - add new user
      - change a user's role
  

- Leave application rules
  - Can only apply 1 day, no half days
  - Can only apply for this year  
  - Cannot apply during peak period (Sep - Nov)
  - Cannot apply if > 50% of the team has applied
  
## Database design

**Users**
```bigquery
UserId       int    `json:"userid,string"`
Password     string `json:"password"`
Name         string `json:"name"`
TeamId       int    `json:"teamid,string"`
Role         string `json:"role"` 
LeaveBalance int    `json:"leave_balance,string"`
```

**Teams**
```bigquery
TeamId      int     `json:"teamid,omitempty,string"`
TeamName    string  `json:"team_name,omitempty"`
ManagerId   int     `json:"managerid,omitempty,string"`
MemberIds   []int   `json:"member_ids,omitempty"`
```

**Leaves (or leave applications)**
```bigquery
LeaveId	   int `json:"leaveid,string"`
UserId     int `json:"userid,string"`
StartDay   int `json:"startDay,string"`
EndDay     int `json:"endDay,string"`
Days       int `json:"days,string"`
Status     int `json:"status,string"`
ApproverId int `json:"approverId,string"`
```


## Instructions:
- run main.go
- all changes will be saved only when you terminate the program by entering "exit"

## Resources:
- [How to organise your project](https://itnext.io/beautify-your-golang-project-f795b4b453aa)
- [Parsing JSON](https://www.sohamkamani.com/golang/parsing-json/)
- [Open JSON while using less space](https://medium.com/kanoteknologi/better-way-to-read-and-write-json-file-in-golang-9d575b7254f2)
- [Mutex](https://golangbot.com/mutex/)
- [Graceful termination and other tips](https://sayedalesawy.hashnode.dev/top-5-lessons-i-learned-while-working-with-go-for-two-years)