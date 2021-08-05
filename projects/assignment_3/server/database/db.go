package database

import (
	"fmt"
	pb "git.garena.com/russell.chanxl/be-class/assignment_3/protos"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db       *gorm.DB
	err      error
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbName   = "assignment_3"
)

//============= INITIALISING DB ==============//

func initDb() {
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName))
	if err != nil {
		log.Println("Failed to open DB:", err)
	}
}

//============= USERS =============//

func GetUser(userId uint64) (*pb.User, error) {
	log.Println("Getting user from DB")
	initDb()
	defer db.Close()
	var user pb.User
	if err = db.Where("user_id = ?", userId).Find(&user).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() *[]pb.User {
	log.Println("Getting all users from DB")
	initDb()
	var users []pb.User
	if err = db.Find(&users).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
	}
	defer db.Close()
	return &users
}

func SetUser(userId uint64, name string, teamId, leaveBalance uint32, role int32) {
	log.Println("Setting user in DB")
	initDb()
	switch userId {
	case 0:
		newUser := &pb.User{
			Password:     "123",
			Name:         name,
			TeamId:       teamId,
			Role:         role,
			LeaveBalance: leaveBalance,
		}
		db.Create(newUser)
	default:
		var user pb.User
		user.Name = name
		user.TeamId = teamId
		user.Role = role
		if err = db.Model(&user).Where("user_id = ?", userId).Update(&user).Error; err != nil {
			log.Printf("DB operation failed: %v", err)
		}
	}
	defer db.Close()

}

func SetUserLeaveBalance(userId uint64, leaveBalance uint32) error {
	log.Println("Setting user's leave balance in DB")
	initDb()
	var user pb.User
	if err = db.Model(&user).Where("user_id = ?", userId).Update("leave_balance", leaveBalance).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
		return err
	}
	defer db.Close()
	return nil
}

func DeleteUser(userId uint64) {
	log.Println("Deleting user in DB")
	initDb()
	var user pb.User
	if err = db.Model(&user).Where("user_id = ?", userId).Delete(&user).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
	}
	defer db.Close()
}

//============= LEAVES ===============//

func GetLeavesByUserId(userId uint64) (*pb.Leaves, error) {
	log.Println("Getting user's leaves from DB")
	initDb()
	defer db.Close()
	var l []*pb.Leave
	if err = db.Where("user_id = ?", userId).Find(&l).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
		return nil, err
	}
	leaves := pb.Leaves{Leaves: l}
	return &leaves, nil
}

func GetLeaveByLeaveId(leaveId uint64) (*pb.Leave, error) {
	log.Println("Getting leave from DB")
	initDb()
	defer db.Close()
	var leave pb.Leave
	if err = db.Where("leave_id = ?", leaveId).Find(&leave).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
		return nil, err
	}
	return &leave, nil
}

func CreateLeave(userId uint64, teamId, start, end, days uint32, status int32, approverId uint64) error {
	log.Println("Creating leave in DB")
	initDb()
	defer db.Close()
	newLeave := &pb.Leave{
		UserId:     userId,
		TeamId:     teamId,
		StartTime:  start,
		EndTime:    end,
		DaysTaken:  days,
		Status:     status,
		ApproverId: approverId,
	}
	if err = db.Create(&newLeave).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
		return err
	}
	return nil
}

func UpdateLeaveStatus(leaveId uint64, status int32) error {
	log.Println("Updating leave in DB")
	initDb()
	defer db.Close()
	var leave pb.Leave
	leave.Status = status
	if err = db.Model(&leave).Where("leave_id = ?", leaveId).Update(&leave).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
		return err
	}
	return nil
}

func DeleteLeave(leaveId uint64) error {
	log.Println("Deleting leave in DB")
	initDb()
	defer db.Close()
	var leave pb.Leave
	if err = db.Model(&leave).Where("leave_id = ?", leaveId).Delete(&leave).Error; err != nil {
		log.Printf("DB operation failed: %v", err)
		return err
	}
	return nil
}
