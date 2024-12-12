package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/core/entity"
	"github.com/kairo913/tasclock-server/app/core/repository"
)

type TaskAppService struct {
	taskRepository repository.TaskRepository
}

func NewTaskAppService(ctx context.Context, taskRepository repository.TaskRepository) *TaskAppService {
	return &TaskAppService{taskRepository}
}

func (tas *TaskAppService) CreateTask(userId int64, title, description string, reward int, deadline time.Time) (task *entity.Task, err error) {
	id := uuid.New()

	task = entity.NewTask(id, userId, title, description, reward, deadline)

	err = tas.taskRepository.Store(task)
	if err != nil {
		return
	}

	return
}

func (tas *TaskAppService) GetTask(taskId string) (task *entity.Task, err error) {
	return tas.taskRepository.GetByTaskId(taskId)
}

func (tas *TaskAppService) GetTasks(userId int64) (tasks *entity.Tasks, err error) {
	return tas.taskRepository.GetAll(userId)
}

func (tas *TaskAppService) DeleteTask(taskId int64) (err error) {
	return tas.taskRepository.Delete(taskId)
}

func (tas *TaskAppService) UpdateTask(task *entity.Task, title, description string, is_done bool, reward, elapsed int, deadline time.Time) (err error) {
	task.UpdateTask(title, description, is_done, reward, elapsed, deadline)

	return tas.taskRepository.Update(task)
}
