package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/lib"
)

type User struct {
	Id        int64     `db:"id" json:"-"`
	UserId    uuid.UUID `db:"user_id" json:"user_id"`
	Lastname  string    `db:"lastname" json:"lastname"`
	Firstname string    `db:"firstname" json:"firstname"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	Salt      string    `db:"salt" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Users []User

func NewUser(lastname, firstname, email, password, secretSalt string, hashCount int) User {
	salt := lib.MakeRandomString(64)

	password = lib.HashString(password+salt+secretSalt, hashCount)

	return User{
		UserId:    uuid.New(),
		Lastname:  lastname,
		Firstname: firstname,
		Email:     email,
		Password:  password,
		Salt:      salt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
