package repository

import (
	"github.com/kairo913/tasclock-server/app/core/entity"
)

type TaskRepository interface {
	Store(*entity.Task) error
	Get(id int64) (*entity.Task, error)
	GetByTaskId(taskId string) (*entity.Task, error)
	GetAll(userId int64) (*entity.Tasks, error)
	Update(*entity.Task) error
	Delete(id int64) error
}
