package repository

import "time"

type DBUser struct {
	Id        int64     `db:"id"`
	UserId    string    `db:"user_id"`
	Lastname  string    `db:"lastname"`
	Firstname string    `db:"firstname"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Salt      string    `db:"salt"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
