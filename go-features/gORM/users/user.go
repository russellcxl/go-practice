package users

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var db *gorm.DB
var err error

type User struct {
	Name  string
	Email string
}

func InitialMigration() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test2")
	if err != nil {
		log.Panicln("failed to open DB:", err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AllUsers endpoint hit\n")
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test2")
	if err != nil {
		log.Panicln("failed to open DB,", err)
	}
	var users []User
	db.Find(&users)
	fmt.Println(users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NewUser enpoint hit\n")
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test2")
	if err != nil {
		log.Panicln("failed to open DB,", err)
	}
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	newUser := &User{
		Name:  name,
		Email: email,
	}
	db.Create(newUser)
	fmt.Fprintf(w, "New user successfully created\n")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteUser enpoint hit\n")
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test2")
	if err != nil {
		log.Panicln("failed to open DB,", err)
	}

	vars := mux.Vars(r)
	name := vars["name"]
	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(user)
	fmt.Fprintf(w, "Deleted user\n")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateUser enpoint hit\n")
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test2")
	if err != nil {
		log.Panicln("failed to open DB,", err)
	}
	vars := mux.Vars(r)
	name := vars["name"]
	newEmail := vars["email"]
	var user User
	db.Where("name = ?",name).Find(&user)
	user.Email = newEmail
	db.Save(&user)
	fmt.Fprintf(w, "Updated user\n")

}
