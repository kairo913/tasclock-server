package repository

import (
	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/core/entity"
)

type UserRepository interface {
	Store(*entity.User) error
	ExistByEmail(email string) (bool, error)
	Get(id int64) (*entity.User, error)
	GetByUserId(userId uuid.UUID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(*entity.User) error
	Delete(id int64) error
}
