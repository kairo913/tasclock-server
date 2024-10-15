package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          int64     `db:"id" json:"-"`
	TaskId      uuid.UUID `db:"task_id" json:"task_id"`
	UserId      int64     `db:"user_id" json:"-"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	IsDone      bool      `db:"is_done" json:"is_done"`
	Reward      float64   `db:"reward" json:"reward"`
	Deadline    time.Time `db:"deadline" json:"deadline"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Tasks []Task
