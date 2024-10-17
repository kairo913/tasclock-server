package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          int64         `db:"id" json:"-"`
	TaskId      uuid.UUID     `db:"task_id" json:"task_id"`
	UserId      int64         `db:"user_id" json:"-"`
	Title       string        `db:"title" json:"title"`
	Description string        `db:"description" json:"description"`
	IsDone      bool          `db:"is_done" json:"is_done"`
	Reward      float64       `db:"reward" json:"reward"`
	Deadline    time.Time     `db:"deadline" json:"deadline"`
	Elapsed     time.Duration `db:"elapsed" json:"elapsed"`
	CreatedAt   time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `db:"updated_at" json:"updated_at"`
}

type Tasks []Task

func NewTask(userId int64, title, description string, isDone bool, reward float64, deadline time.Time) Task {
	return Task{
		TaskId:      uuid.New(),
		UserId:      userId,
		Title:       title,
		Description: description,
		IsDone:      isDone,
		Reward:      reward,
		Deadline:    deadline,
		Elapsed:     time.Duration(0),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) UpdateTitle(title string) {
	t.Title = title
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateIsDone(isDone bool) {
	t.IsDone = isDone
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateReward(reward float64) {
	t.Reward = reward
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateDeadline(deadline time.Time) {
	t.Deadline = deadline
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateElapsed(elapsed time.Duration) {
	t.Elapsed = elapsed
	t.UpdatedAt = time.Now()
}

func (t *Task) IsDeadlineExceeded() bool {
	return time.Now().After(t.Deadline)
}
