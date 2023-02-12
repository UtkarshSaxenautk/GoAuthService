package model

import "time"

type DateOfBirth struct {
	Year      int
	Month     int
	MonthName string
	Date      int
}

type User struct {
	Email             string
	Username          string
	PasswordHash      string
	FullName          string
	Role              string
	Dob               DateOfBirth
	CreateTs          time.Time
	UpdateTs          time.Time
	LoginTs           time.Time
	PreviousPasswords []string
	Otp               string
}
