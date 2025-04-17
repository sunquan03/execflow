package repositories

import (
	"context"
	"encoding/json"
	"exec_flow/pkg/models"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	client *redis.Client
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{
		client: client,
	}
}

const (
	TaskQueue    = "task_queue"
	TaskMetadata = "task_metadata"
	TaskStatus   = "task_status"
)

func (r *Repository) Ping() error {
	return nil
}

func (r *Repository) RegisterTask(id string, task *models.TaskRequest) (err error) {
	ctx := context.TODO()
	jsonB, err := json.Marshal(task)
	if err != nil {
		return err
	}
	_, err = r.client.HSet(ctx, TaskMetadata, id, string(jsonB)).Result()
	return err
}

func (r *Repository) ScheduleTask(id string, score float64) (err error) {
	ctx := context.TODO()
	_, err = r.client.ZAdd(ctx, TaskQueue, redis.Z{Score: score, Member: id}).Result()
	return err
}

func (r *Repository) SetTaskStatus(id string, status string) (err error) {
	ctx := context.TODO()
	_, err = r.client.HSet(ctx, TaskStatus, id, status).Result()
	return err
}

func (r *Repository) GetTask(id string) (task string, err error) {
	ctx := context.TODO()
	err = r.client.HGet(ctx, TaskMetadata, id).Scan(&task)
	return task, err
}

func (r *Repository) GetTaskStatus(id string) (status string, err error) {
	ctx := context.TODO()
	err = r.client.HGet(ctx, TaskStatus, id).Scan(&status)
	return status, err
}

func (r *Repository) GetTaskQueue(n int64) ([]models.Task, error) {
	ctx := context.TODO()
	jsonTasks, err := r.client.ZRange(ctx, TaskQueue, 0, n-1).Result()
	if err != nil {
		return nil, err
	}

	tasks := make([]models.Task, len(jsonTasks))
	for i, j := range jsonTasks {
		var task models.Task
		err := json.Unmarshal([]byte(j), &task)
		if err != nil {
			return nil, err
		}
		tasks[i] = task
	}

	return tasks, nil
}

func (r *Repository) DelTaskSchedule(id string) error {
	ctx := context.TODO()
	_, err := r.client.ZRem(ctx, TaskQueue, id).Result()
	return err
}
