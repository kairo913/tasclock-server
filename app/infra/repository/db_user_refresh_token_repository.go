package repository

import "github.com/kairo913/tasclock-server/app/core/entity"

type DBUserRefreshTokenRepository struct {
	SqlHandler SqlHandler
}

func NewUserRefreshTokenRepository(sqlHandler SqlHandler) *DBUserRefreshTokenRepository {
	return &DBUserRefreshTokenRepository{SqlHandler: sqlHandler}
}

func (repo DBUserRefreshTokenRepository) Store(urt *entity.UserRefreshToken) error {
	_, err := repo.SqlHandler.Execute("INSERT INTO user_refresh_token (token, expire_at) VALUES (?, ?)", urt.RefreshToken, urt.ExpiredAt)
	if err != nil {
		return err
	}

	return nil
}

func (repo DBUserRefreshTokenRepository) Update(urt *entity.UserRefreshToken) error {
	_, err := repo.SqlHandler.Execute("UPDATE user_refresh_token SET is_used = ? WHERE id = ?", urt.IsUsed, urt.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repo DBUserRefreshTokenRepository) Exist(token string) (bool, error) {
	row, err := repo.SqlHandler.Query("SELECT COUNT(*) FROM user_refresh_token WHERE token = ?", token)
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

func (repo DBUserRefreshTokenRepository) Get(token string) (*entity.UserRefreshToken, error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM user_refresh_token WHERE token = ? LIMIT 1", token)
	if err != nil {
		return nil, err
	}

	var urt entity.UserRefreshToken
	if row.Next() {
		err = row.Scan(&urt.Id, &urt.RefreshToken, &urt.ExpiredAt, &urt.IsUsed)
		if err != nil {
			return nil, err
		}
	}

	return &urt, nil
}

func (repo DBUserRefreshTokenRepository) Delete(id int64) error {
	_, err := repo.SqlHandler.Execute("DELETE FROM user_refresh_token WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
