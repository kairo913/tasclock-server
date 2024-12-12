package repository

import "time"

type DBTask struct {
	Id          int64     `db:"id"`
	TaskId      string    `db:"task_id"`
	UserId      int64     `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	IsDone      bool      `db:"is_done"`
	Reward      int       `db:"reward"`
	Elapsed     int       `db:"elapsed"`
	Deadline    time.Time `db:"deadline"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
