package model

import "time"

type Employee struct {
	ID          int64
	UserID      int64
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	JobTitle    string
	HireDate    time.Time
	Department  string
}
