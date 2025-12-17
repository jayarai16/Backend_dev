package models

import (
	"time"
)

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" validate:"required,min=1,max=100"`
	DOB  time.Time `json:"dob" db:"dob" validate:"required"`
}

type UserWithAge struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
	DOB  string `json:"dob" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
	DOB  string `json:"dob" validate:"required"`
}

func (u *UserWithAge) CalculateAge() {
	dob, err := time.Parse("2006-01-02", u.DOB)
	if err != nil {
		u.Age = 0
		return
	}
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	u.Age = age
}