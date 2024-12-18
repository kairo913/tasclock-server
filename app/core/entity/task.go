package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	Reward      int       `json:"reward"`
	Elapsed     int       `json:"elapsed"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tasks []Task

func NewTask(id, userId uuid.UUID, title, description string, reward int, deadline time.Time) *Task {
	return &Task{
		Id:          id,
		UserId:      userId,
		Title:       title,
		Description: description,
		IsDone:      false,
		Reward:      reward,
		Elapsed:     0,
		Deadline:    deadline,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) UpdateTask(title, description string, isDone bool, reward, elapsed int, deadline time.Time) {
	t.Title = title
	t.Description = description
	t.IsDone = isDone
	t.Reward = reward
	t.Elapsed = elapsed
	t.Deadline = deadline
	t.UpdatedAt = time.Now()
}
