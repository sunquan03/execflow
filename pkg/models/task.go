package models

import "time"

type TaskRequest struct {
	TaskType string `json:"task_type"`
	Payload  string `json:"payload"`
	Priority int    `json:"priority"`
	Deadline string `json:"deadline"`
}

type Task struct {
	ID       string    `json:"ID"`
	TaskType string    `json:"task_type"`
	Payload  string    `json:"payload"`
	Priority int       `json:"priority"`
	Deadline time.Time `json:"deadline"`
	Status   string    `json:"status"`
}
