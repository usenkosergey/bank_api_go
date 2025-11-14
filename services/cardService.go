package services

import (
	"bank_api_go/database"
	"bank_api_go/models"
	"encoding/json"
	"fmt"
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
	db := database.DbConnect()
	defer db.Close()

	card.NumberFull = generateCardNumber()
	card.ValidityPeriod = time.Now().AddDate(3, 0, 0)

	_, _ = db.Exec(
		`INSERT INTO cards (full_number, owner_id, validity_period) VALUES($1, $2, $3)`,
		&card.NumberFull, &card.OwnerId, &card.ValidityPeriod)
}

func GetCards(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * FROM cards`

	db := database.DbConnect()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var card models.Card
		fmt.Print(card)
		err := rows.Scan(
			&card.Id, &card.NumberFull, &card.NumberHidden, &card.OwnerFullName, &card.ValidityPeriod, &card.Balance,
		)
		if err != nil {
			log.Fatal(err)
		}
		//card = append(cards, card)
	}

	fmt.Printf("Total users: %d\n", len(cards))
	//for _, card := range cards {
	//	fmt.Printf(" - %s (%s)\n", card.Name, card.Email)
	//}
}

func generateCardNumber() string {
	return strconv.Itoa(1000+rand.Intn(1000)) + " " + strconv.Itoa(1000+rand.Intn(1000)) + " " +
		strconv.Itoa(1000+rand.Intn(1000)) + " " + strconv.Itoa(1000+rand.Intn(1000))
}
