package models

type Person struct {
	Id       int64  `json:"id"`
	FullName string `json:"fullName"`
	Role     Role   `json:"role"`
}
