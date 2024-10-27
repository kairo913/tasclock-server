package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/core/entity"
	"github.com/kairo913/tasclock-server/app/core/repository"
	"github.com/kairo913/tasclock-server/app/util"
	"github.com/kairo913/tasclock-server/app/util/config"
)

type UserAppService struct {
	userRepository repository.UserRepository
	hashConfig     *config.HashConfig
	sessionConfig  *config.SessionConfig
}

func NewUserAppService(ctx context.Context, userRepository repository.UserRepository) *UserAppService {
	return &UserAppService{userRepository, config.NewHashConfig(ctx), config.NewSessionConfig(ctx)}
}

func (uas *UserAppService) CreateUser(lastname, firstname, email, password string) (user *entity.User, err error) {
	id := uuid.New()

	salt := util.MakeRandomString(64)

	password = util.Hash(password+salt+uas.hashConfig.SecretSalt, uas.hashConfig.HashCount)

	user = entity.NewUser(id, lastname, firstname, email, password, salt)

	err = uas.userRepository.Store(user)
	if err != nil {
		return
	}

	return
}

func (uas *UserAppService) ExistByEmail(email string) (bool, error) {
	return uas.userRepository.ExistByEmail(email)
}

func (uas *UserAppService) GetUser(userId string) (user *entity.User, err error) {
	user, err = uas.userRepository.GetByUserId(userId)
	if err != nil {
		return
	}

	return
}

func (uas *UserAppService) GetUserByEmail(email string) (user *entity.User, err error) {
	user, err = uas.userRepository.GetByEmail(email)
	if err != nil {
		return
	}

	return
}

func (uas *UserAppService) VerifyPassword(user *entity.User, password string) bool {
	hashedPassword := util.Hash(password+user.Salt+uas.hashConfig.SecretSalt, uas.hashConfig.HashCount)

	return user.Password == hashedPassword
}

func (uas *UserAppService) UpdateUser(userId, lastname, firstname, email string) (err error) {
	user, err := uas.userRepository.GetByUserId(userId)
	if err != nil {
		return
	}

	user.UpdateLastname(lastname)
	user.UpdateFirstname(firstname)
	user.UpdateEmail(email)

	err = uas.userRepository.Update(user)
	if err != nil {
		return
	}

	return
}

func (uas *UserAppService) UpdateEmail(userId, email string) (err error) {
	user, err := uas.userRepository.GetByUserId(userId)
	if err != nil {
		return
	}

	user.UpdateEmail(email)

	err = uas.userRepository.Update(user)
	if err != nil {
		return
	}

	return
}

func (uas *UserAppService) UpdatePassword(userId, password string) (err error) {
	user, err := uas.userRepository.GetByUserId(userId)
	if err != nil {
		return
	}

	salt := util.MakeRandomString(64)

	password = util.Hash(password+salt+uas.hashConfig.SecretSalt, uas.hashConfig.HashCount)

	user.UpdatePassword(password, salt)

	err = uas.userRepository.Update(user)
	if err != nil {
		return
	}

	return
}
