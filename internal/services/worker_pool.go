package services

import (
	"exec_flow/internal/repositories"
	"exec_flow/pkg/models"
)

type WorkerPool struct {
	wn    int
	taskQ chan models.Task
	resQ  chan models.TaskResult
	repo  *repositories.Repository
}

func NewWorkerPool(wn int, repo *repositories.Repository) *WorkerPool {
	return &WorkerPool{
		wn:    wn,
		taskQ: make(chan models.Task, wn),
		resQ:  make(chan models.TaskResult, wn),
		repo:  repo,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.wn; i++ {
		w := Worker{id: i, taskQ: wp.taskQ, resQ: wp.resQ}
		w.Start()
	}
}

func (wp *WorkerPool) Submit(id string, taskData models.TaskData) {
	switch taskData.TaskType {
	case "SQL":
		t := models.Task{
			ID:       id,
			Exec:     ExecSQL,
			TaskData: taskData,
		}
		wp.taskQ <- t
	case "HTTP_REQUEST":
		t := models.Task{
			ID:       id,
			Exec:     ExecHttpRequest,
			TaskData: taskData,
		}
		wp.taskQ <- t
	}
}

func (wp *WorkerPool) HandleResults() {
	for {
		res := <-wp.resQ
		// TODO
		if res.Err != nil {

		}
	}
}

type Worker struct {
	id    int
	taskQ <-chan models.Task
	resQ  chan<- models.TaskResult
}

func (w *Worker) Start() {
	task := <-w.taskQ
	w.resQ <- task.Exec(task.ID, task.TaskData.Payload)
}
