package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Lastname string    `json:"lastname"`
}

type Users []User

func NewUser(lastname string) *User {
	return &User{
		Lastname: lastname,
	}
}

func (u *User) UpdateLastname(lastname string) {
	u.Lastname = lastname
}