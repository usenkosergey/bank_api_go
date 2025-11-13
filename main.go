package main

import (
	"./services"
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//r.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/person", СreatePerson).Methods("POST")
	//r.HandleFunc("/users/{id}", getUser).Methods("GET")
	//r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	//r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("------")
}
