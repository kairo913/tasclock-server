package controller

import "time"

type SignUpResponse struct {
	Id        string    `json:"id"`
	Lastname  string    `json:"lastname"`
	Firstname string    `json:"firstname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

