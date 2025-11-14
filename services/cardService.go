package services

import (
	"bank_api_go/database"
	"bank_api_go/models"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func CreateCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card
	err := json.NewDecoder(r.Body).Decode(&card)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	db, _ := database.GetDB()

	card.NumberFull = generateCardNumber()
	card.ValidityPeriod = time.Now().AddDate(3, 0, 0)

	_, _ = db.Exec(
		`INSERT INTO cards (full_number, owner_id, validity_period) VALUES($1, $2, $3)`,
		&card.NumberFull, &card.OwnerId, &card.ValidityPeriod)
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	query := `SELECT CA.id, CA.full_number, PP.full_name, CA.validity_period, CA.balance
				FROM cards CA
				LEFT JOIN persons PP ON PP.id = CA.owner_id `

	db, _ := database.GetDB()
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var card models.Card
		err := rows.Scan(&card.Id, &card.NumberFull, &card.OwnerFullName, &card.ValidityPeriod, &card.Balance)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}
	writeJSON(w, http.StatusOK, cards)
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func generateCardNumber() string {
	return strconv.Itoa(1000+rand.Intn(1000)) + " " + strconv.Itoa(1000+rand.Intn(1000)) + " " +
		strconv.Itoa(1000+rand.Intn(1000)) + " " + strconv.Itoa(1000+rand.Intn(1000))
}
