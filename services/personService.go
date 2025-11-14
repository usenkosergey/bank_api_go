package services

import (
	"bank_api_go/database"
	"bank_api_go/models"
	"encoding/json"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	db, _ := database.GetDB()

	_, _ = db.Exec(
		`INSERT INTO persons(full_name) VALUES ($1)`, &person.FullName)
}
