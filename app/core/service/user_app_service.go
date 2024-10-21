package service

import (
	"github.com/kairo913/tasclock-server/app/core/entity"
	"github.com/kairo913/tasclock-server/app/core/repository"
)

type UserAppService struct {
	userRepository repository.UserRepository
}

func NewUserAppService(userRepository repository.UserRepository) *UserAppService {
	return &UserAppService{userRepository}
}

func (uas *UserAppService) CreateUser(lastname string) (user *entity.User, err error) {
	user = entity.NewUser(lastname)

	err = uas.userRepository.Store(user)
	if err != nil {
		return
	}

	return
}

func (uas *UserAppService) UpdateLastname(userId, lastname string) (err error) {
	user, err := uas.userRepository.GetByUserId(userId)
	if err != nil {
		return
	}

	user.UpdateLastname(lastname)

	err = uas.userRepository.Update(user)
	if err != nil {
		return
	}

	return
}
