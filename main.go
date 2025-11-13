package main

import (
	"bank_api_go/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//r.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/person", services.CreatePerson).Methods("POST")
	//r.HandleFunc("/users/{id}", getUser).Methods("GET")
	//r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	//r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("------")

	log.Println("Server starting on :9090")
	log.Fatal(http.ListenAndServe(":9090", router))
}
