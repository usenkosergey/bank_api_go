package main

import "time"

type Card struct {
	numberFull     string
	numberHidden   string
	owner          string
	validityPeriod time.Time
	status         string
	balance
}

const (
	Active = iota
	Blocked
	Expired
)
