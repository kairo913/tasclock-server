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
	_, err := repo.SqlHandler.Execute("INSERT INTO users (user_id, lastname, firstname, email, password, salt, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.Id.String(), user.Lastname, user.Firstname, user.Email, user.Password, user.Salt, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo *DBUserRepository) ExistByEmail(email string) (bool, error) {
	row, err := repo.SqlHandler.Query("SELECT COUNT(*) FROM users WHERE email = ?", email)
	if err != nil {
		return false, err
	}

	var count int
	if row.Next() {
		err = row.Scan(&count)
		if err != nil {
			return false, err
		}
	}

	return count > 0, nil
}

func (repo *DBUserRepository) Get(id int64) (*entity.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE id = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	var user DBUser
	if row.Next() {
		err = row.Scan(&user.Id, &user.UserId, &user.Lastname, &user.Firstname, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &entity.User{
		Id:        uuid.MustParse(user.UserId),
		Lastname:  user.Lastname,
		Firstname: user.Firstname,
		Email:     user.Email,
		Password:  user.Password,
		Salt:      user.Salt,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (repo *DBUserRepository) GetByUserId(userId string) (*entity.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE user_id = ? LIMIT 1", userId)
	if err != nil {
		return nil, err
	}

	var user DBUser
	if row.Next() {
		err = row.Scan(&user.Id, &user.UserId, &user.Lastname, &user.Firstname, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &entity.User{
		Id:        uuid.MustParse(user.UserId),
		Lastname:  user.Lastname,
		Firstname: user.Firstname,
		Email:     user.Email,
		Password:  user.Password,
		Salt:      user.Salt,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (repo *DBUserRepository) GetByEmail(email string) (*entity.User, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE email = ? LIMIT 1", email)
	if err != nil {
		return nil, err
	}

	var user DBUser
	if row.Next() {
		err = row.Scan(&user.Id, &user.UserId, &user.Lastname, &user.Firstname, &user.Email, &user.Password, &user.Salt, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &entity.User{
		Id:        uuid.MustParse(user.UserId),
		Lastname:  user.Lastname,
		Firstname: user.Firstname,
		Email:     user.Email,
		Password:  user.Password,
		Salt:      user.Salt,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (repo *DBUserRepository) Update(user *entity.User) error {
	_, err := repo.SqlHandler.Execute("UPDATE users SET lastname = ?, firstname = ?, email = ?, password = ?, salt = ?, updated_at = ? WHERE user_id = ?", user.Lastname, user.Firstname, user.Password, user.Salt, user.UpdatedAt, user.Id.String())
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
