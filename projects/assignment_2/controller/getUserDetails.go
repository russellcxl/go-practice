package controller

import (
	"errors"
	"github.com/russellcxl/go-practice/assignment_2/database"
	"github.com/russellcxl/go-practice/assignment_2/models"
)

var UserNotFound = errors.New("user not found")

func GetUserDetails(userid int) (*models.User, error) {
	var user *models.User
	err := UserNotFound

	for i := 0 ; i < len(database.Users) ; i++ {
		if database.Users[i].UserId == userid {
			user = &database.Users[i]
			err = nil
		}
	}
	return user, err
}
