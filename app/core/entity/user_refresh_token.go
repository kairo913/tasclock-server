package entity

import "time"

type UserRefreshToken struct {
	Id           int64     `json:"id"`
	RefreshToken string    `json:"token"`
	ExpiredAt    time.Time `json:"expired_at"`
	IsUsed       bool      `json:"is_used"`
}

type UserRefreshTokens []UserRefreshToken

func NewUserRefreshToken(refreshToken string, expiredAt time.Time) *UserRefreshToken {
	return &UserRefreshToken{
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
		IsUsed:       false,
	}
}

func (urt *UserRefreshToken) Use() {
	urt.IsUsed = true
}