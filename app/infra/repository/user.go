package repository

type DBUser struct {
	Id       int64  `db:"id"`
	UserId   string `db:"user_id"`
	Lastname string `db:"lastname"`
}
