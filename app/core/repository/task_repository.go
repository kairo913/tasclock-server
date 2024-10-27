package repository

import (
	"github.com/kairo913/tasclock-server/app/core/entity"
)

type TaskRepository interface {
	Store(*entity.Task) error
	Get(id string) (*entity.Task, error)
	GetAll(userId string) (entity.Tasks, error)
	Update(*entity.Task) error
	Delete(id string) error
}
