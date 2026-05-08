package models

import "time"

type User struct {
	Id string `json:"id"`
	FName string `json:"fname"`
	LName string `json:"lname"`
	Email string `json:"email"`
	Age int `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserPayload struct {
	FName string `json:"fname"`
	LName string `json:"lname"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}