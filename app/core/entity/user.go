package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Lastname  string    `json:"lastname"`
	Firstname string    `json:"firstname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

func NewUser(id uuid.UUID, lastname, firstname, email, password, salt string) *User {
	return &User{
		Id:        id,
		Lastname:  lastname,
		Firstname: firstname,
		Email:     email,
		Password:  password,
		Salt:      salt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) UpdateLastname(lastname string) {
	u.Lastname = lastname
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateFirstname(firstname string) {
	u.Firstname = firstname
	u.UpdatedAt = time.Now()
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

func (u *User) UpdatePassword(password, salt string) {
	u.Password = password
	u.Salt = salt
	u.UpdatedAt = time.Now()
}
