package controller

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/validations"
	"time"
)

func SetUserRole(userid int, role string) {

	if validations.CheckUserIdIsValid(userid) != true {
		fmt.Println("> User ID invalid")
		return
	}


	for i := 0 ; i < len(database.Users) ; i++ {
		if database.Users[i].UserId == userid {
			database.Users[i].Role = role
		}
	}
	fmt.Printf("> User %d is now an %s\n", userid, role)
	time.Sleep(time.Second)
}
