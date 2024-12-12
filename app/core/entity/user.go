package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        int64     `json:"-"`
	UserId    uuid.UUID `json:"id"`
	Lastname  string    `json:"lastname"`
	Firstname string    `json:"firstname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Salt      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

func NewUser(userId uuid.UUID, lastname, firstname, email, password, salt string) *User {
	return &User{
		UserId:    userId,
		Lastname:  lastname,
		Firstname: firstname,
		Email:     email,
		Password:  password,
		Salt:      salt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) UpdateName(lastname, firstname string) {
	u.Lastname = lastname
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
