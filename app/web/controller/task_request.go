package controller

import "time"

type CreateTaskRequest struct {
	Title       string    `json:"title" validate:"required, max=50"`
	Description string    `json:"description" validate:"max=255"`
	Reward      int       `json:"reward"`
	Deadline    time.Time `json:"deadline"`
}

type GetTaskRequest struct {
	Id string `json:"id" validate:"required"`
}

type UpdateTaskRequest struct {
	Id          string    `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"max=50"`
	Description string    `json:"description" validate:"max=255"`
	IsDone      bool      `json:"is_done"`
	Reward      int       `json:"reward"`
	Elapsed     int       `json:"elapsed"`
	Deadline    time.Time `json:"deadline"`
}

type DeleteTaskRequest struct {
	Id string `json:"id" validate:"required"`
}
