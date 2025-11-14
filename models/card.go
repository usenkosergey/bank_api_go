package models

import "time"

type Card struct {
	Id             int64     `json:"id"`
	NumberFull     string    `json:"numberFull"`
	NumberHidden   string    `json:"numberHidden"`
	OwnerId        int64     `json:"ownerId"`
	OwnerFullName  string    `json:"ownerFullName"`
	ValidityPeriod time.Time `json:"validityPeriod"`
	Status         Status    `json:"status"`
	Balance        int64     `json:"balance"`
}
