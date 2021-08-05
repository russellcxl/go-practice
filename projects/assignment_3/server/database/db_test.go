package database

import (
	"fmt"
	"github.com/russellcxl/go-practice/assignment_3/protos"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	user, _ := GetUser(2000)
	fmt.Println(user)
}

func TestGetAllUsers(t *testing.T) {
	GetAllUsers()
}

func TestSetUser(t *testing.T) {
	t.Run("Create new user", func(t *testing.T) {
		SetUser(0, "Bill", 100, 10, int32(protos.UserRole_ROLE_MANAGER))
	})
	t.Run("Update user's name", func(t *testing.T) {
		SetUser(1001, "Elle", 100, 11, int32(protos.UserRole_ROLE_STAFF))
	})
}

func TestSetUserLeaves(t *testing.T) {
	SetUserLeaveBalance(1003, 15)
}

func TestDeletUser(t *testing.T) {
	DeleteUser(1006)
}

func TestGetLeavesByUserId(t *testing.T) {
	leaves, _ := GetLeavesByUserId(1000)
	for _, l := range leaves.GetLeaves() {
		fmt.Println(l)
	}
}

func TestGetLeaveByLeaveId(t *testing.T) {
	leave, _ := GetLeaveByLeaveId(2000)
	fmt.Println(leave)
}

func TestSetLeave(t *testing.T) {
	start  := "11/11/21"
	end  := "12/11/21"
	layout := "02/01/06"
	s, _ := time.Parse(layout, start)
	e, _ := time.Parse(layout, end)
	sUnix := uint32(s.Unix())
	eUnix := uint32(e.Unix())

	fmt.Println(sUnix, eUnix)

	CreateLeave(1001, 100, sUnix, eUnix, 1, 0, 1007)
}

func TestDeleteLeave(t *testing.T) {
	DeleteLeave(123)
}
