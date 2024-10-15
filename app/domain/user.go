package domain

import (
	"time"

	"github.com/google/uuid"
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
