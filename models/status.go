package models

type Status int

const (
	Active Status = iota
	Blocked
	Expired
)
