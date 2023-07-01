package entity

import "time"

type User struct {
	ID        string
	FirstName string
	LastName  string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
}
