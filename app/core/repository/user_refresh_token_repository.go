package repository

import "github.com/kairo913/tasclock-server/app/core/entity"

type UserRefreshTokenRepository interface {
	Store(*entity.UserRefreshToken) error
	Update(*entity.UserRefreshToken) error
	Exist(token string) (bool, error)
	Get(token string) (*entity.UserRefreshToken, error)
	Delete(id int64) error
}