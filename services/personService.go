package services

import (
	"bank_api_go/database"
	"bank_api_go/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("+++++++")
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	db := database.DbConnect()
	defer db.Close()

	var id int
	err = db.QueryRow(
		"INSERT INTO persons (full_name) VALUES ($1) RETURNING id",
		person.FullName,
	).Scan(&id)
	fmt.Println("aaaaaaaaaaaaaaa", id)

	//_, _ = db.Exec(
	//	`INSERT INTO persons(full_name) VALUES ('adsf')`, &person.FullName)
	fmt.Print(person)
}
