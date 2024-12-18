package controller

import "time"

type CreateTaskResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	Reward      int       `json:"reward"`
	Elapsed     int       `json:"elapsed"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetTaskResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	Reward      int       `json:"reward"`
	Elapsed     int       `json:"elapsed"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetTasksResponse struct {
	Tasks []GetTaskResponse `json:"tasks"`
}