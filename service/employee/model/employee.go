package model

import "time"

type Employee struct {
	ID             int64
	FirstName      string
	JobDescription string
	SignDate       time.Time
	UserID         int64
}
