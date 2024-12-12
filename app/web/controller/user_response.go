package controller

import "time"

type GetUserResponse struct {
	Lastname  string    `json:"lastname"`
	Firstname string    `json:"firstname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
