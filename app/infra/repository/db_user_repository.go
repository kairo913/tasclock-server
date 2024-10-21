package repository

import (
	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/core/entity"
)

type DBUserRepository struct {
	SqlHandler SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) *DBUserRepository {
	return &DBUserRepository{sqlHandler}
}

func (repo *DBUserRepository) Store(user *entity.User) error {
	_, err := repo.SqlHandler.Execute("INSERT INTO users (user_id, lastname) VALUES (?, ?)", user.Id.String(), user.Lastname)
	if err != nil {
		return err
	}

	return nil
}

func (repo *DBUserRepository) Get(id int64) (*entity.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	var user DBUser
	if row.Next() {
		err = row.Scan(&user.Id, &user.UserId, &user.Lastname)
		if err != nil {
			return nil, err
		}
	}

	return &entity.User{
		Id:       uuid.MustParse(user.UserId),
		Lastname: user.Lastname,
	}, nil
}

func (repo *DBUserRepository) GetByUserId(userId string) (*entity.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE user_id = ? LIMIT 1", userId)
	if err != nil {
		return nil, err
	}

	var user DBUser
	if row.Next() {
		err = row.Scan(&user.Id, &user.UserId, &user.Lastname)
		if err != nil {
			return nil, err
		}
	}

	return &entity.User{
		Id:       uuid.MustParse(user.UserId),
		Lastname: user.Lastname,
	}, nil
}

func (repo *DBUserRepository) Update(user *entity.User) error {
	_, err := repo.SqlHandler.Execute("UPDATE users SET lastname = ? WHERE user_id = ?", user.Lastname, user.Id.String())
	if err != nil {
		return err
	}

	return nil
}

func (repo *DBUserRepository) Delete(userId string) error {
	_, err := repo.SqlHandler.Execute("DELETE FROM users WHERE user_id = ?", userId)
	if err != nil {
		return err
	}

	return nil
}