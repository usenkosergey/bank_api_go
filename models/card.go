package models

import "time"

type Card struct {
	id             int64
	numberFull     string
	numberHidden   string
	owner          Person
	validityPeriod time.Time
	status         Status
	balance        int64
}
