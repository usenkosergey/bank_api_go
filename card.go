package main

import "time"

type Card struct {
	numberFull     string
	numberHidden   string
	owner          string
	validityPeriod time.Time
	status         Status
	balance        int64
}

type Status int

const (
	Active Status = iota
	Blocked
	Expired
)
