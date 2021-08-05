package database

import (
	"encoding/json"
	"fmt"
	"git.garena.com/russell.chanxl/be-class/assignment_2/models"
	"io/ioutil"
	"log"
	"os"
)

var Users []models.User

var Teams []models.Team

var Leaves []models.Leave

var CurrentUser *models.User

var path = "/Users/russell.chanxl/go/src/git.garena.com/russell.chanxl/be-class/assignment_2/assets/"

func InitDb() {


	// open users json
	file, err := os.Open(path + "users.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &Users)
	fmt.Println("> Successfully opened Users DB")
	if err != nil {
		fmt.Println(err)
	}


	// open the teams json
	file, err = os.Open(path + "teams.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ = ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &Teams)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("> Successfully opened Teams DB")


	// open the leaves json
	file, err = os.Open(path + "leaves.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	byteValue, _ = ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &Leaves)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("> Successfully opened Leaves DB")
	fmt.Println("--------------------------------")

}

func WriteToDbs() {

	// save users json
	file, _ := json.MarshalIndent(Users, "", "	")
	err := ioutil.WriteFile(path + "users.json", file, 0644)
	if err != nil {
		fmt.Println("> Couldn't write to DB:", err)
	}

	// save leaves json
	file, _ = json.MarshalIndent(Leaves, "", "	")
	err = ioutil.WriteFile(path + "leaves.json", file, 0644)
	if err != nil {
		fmt.Println("> Couldn't write to DB:", err)
	}

	// save teams json
	file, _ = json.MarshalIndent(Teams, "", "	")
	err = ioutil.WriteFile(path + "teams.json", file, 0644)
	if err != nil {
		fmt.Println("> Couldn't write to DB:", err)
	}

}
