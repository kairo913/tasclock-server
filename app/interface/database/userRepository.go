package database

import "github.com/kairo913/tasclock-server/app/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int64, err error) {
	r, err := repo.SqlHandler.Execute(
		"INSERT INTO users (user_id, lastname, firstname, email, password, salt) VALUES (?, ?, ?, ?, ?, ?)", u.UserId, u.Lastname, u.Firstname, u.Email, u.Password, u.Salt,
	)

	if err != nil {
		return
	}

	id, err = r.LastInsertId()

	if err != nil {
		return -1, err
	}

	return
}

func (repo *UserRepository) FindById(id string) (user domain.User, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE user_id = ? LIMIT 1", id)

	if err != nil {
		return
	}

	defer row.Close()

	var u domain.User
	row.Next()
	if err = row.Scan(&u.Id, &u.UserId, &u.Lastname, &u.Firstname, &u.Email, &u.Password, &u.Salt, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return
	}

	user = u

	return
}

func (repo *UserRepository) FindByEmail(email string) (user domain.User, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM users WHERE email = ? LIMIT 1", email)

	if err != nil {
		return
	}

	defer row.Close()

	var u domain.User
	row.Next()
	if err = row.Scan(&u.Id, &u.UserId, &u.Lastname, &u.Firstname, &u.Email, &u.Password, &u.Salt, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return
	}

	user = u

	return
}

func (repo *UserRepository) Update(u domain.User) (err error) {
	_, err = repo.SqlHandler.Execute(
		"UPDATE users SET lastname = ?, firstname = ?, email = ?, password = ?, salt = ? WHERE id = ?;", u.Lastname, u.Firstname, u.Email, u.Password, u.Salt, u.Id,
	)

	if err != nil {
		return
	}

	return
}

func (repo *UserRepository) Delete(id string) (err error) {
	_, err = repo.SqlHandler.Execute("DELETE FROM users WHERE user_id = ?", id)

	if err != nil {
		return
	}

	return
}
