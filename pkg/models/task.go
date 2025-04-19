package models

type TaskData struct {
	TaskType string `json:"task_type"`
	Payload  string `json:"payload"`
	Priority int    `json:"priority"`
	Deadline string `json:"deadline"`
}

type TaskResult struct {
	ID  string
	Res int
	Err error
}

type TaskExec func(id, payload string) (res TaskResult)

type Task struct {
	ID   string
	Exec TaskExec
	TaskData
}
