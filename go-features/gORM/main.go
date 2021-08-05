package main

import (
	"fmt"
	"git.garena.com/russell.chanxl/personal/gORM/users"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/*

========= SQL =========
- log into SQL `mysql -u root -p`
- password is root

 */

func main() {
	users.InitialMigration()
	handleRequest()
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// routes
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/users", users.AllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{name}/{email}", users.NewUser).Methods("POST")
	myRouter.HandleFunc("/users/{name}", users.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users/{name}/{email}", users.UpdateUser).Methods("PUT")

	log.Fatalln(http.ListenAndServe(":8081", myRouter))

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}
