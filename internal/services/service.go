package services

import (
	"exec_flow/internal/repositories"
	"exec_flow/pkg/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}

const (
	StatusPending    = "pending"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
	StatusFailed     = "failed"
)

func (s *Service) CreateTask(task *models.TaskData) (id string, err error) {
	id = uuid.New().String()
	err = s.repo.RegisterTask(id, task)
	if err != nil {
		return "", err
	}
	deadline, err := time.Parse("02.01.2006 15:04:05", task.Deadline)
	if err != nil {
		return "", fmt.Errorf("invalid deadline format: %w", err)
	}
	var score float64
	score = float64(task.Priority)*10e10 + float64(deadline.Unix())
	err = s.repo.ScheduleTask(id, score)
	if err != nil {
		return "", err
	}
	err = s.repo.SetTaskStatus(id, StatusPending)
	if err != nil {
		return "", err
	}
	return
}
