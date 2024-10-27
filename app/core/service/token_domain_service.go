package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kairo913/tasclock-server/app/core/entity"
	"github.com/kairo913/tasclock-server/app/core/repository"
	"github.com/kairo913/tasclock-server/app/util/config"
)

type TokenDomainService struct {
	userTokenRepository repository.UserRefreshTokenRepository
	sessionConfig       *config.SessionConfig
}

func NewTokenDomainService(ctx context.Context, userTokenRepository repository.UserRefreshTokenRepository) *TokenDomainService {
	return &TokenDomainService{userTokenRepository, config.NewSessionConfig(ctx)}
}

func (tds *TokenDomainService) GetRefreshTokenAge() int {
	return int(tds.sessionConfig.RefreshTokenExpire.Seconds())
}

func (tds *TokenDomainService) GenerateToken(userId string) (accessToken, refreshToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "tasclock",
		Subject:   "access",
		Audience:  []string{userId},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tds.sessionConfig.AccessTokenExpire)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	accessToken, err = token.SignedString([]byte(tds.sessionConfig.JWTSecret))
	if err != nil {
		return
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "tasclock",
		Subject:   "refresh",
		Audience:  []string{userId},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tds.sessionConfig.RefreshTokenExpire)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	})

	refreshToken, err = token.SignedString([]byte(tds.sessionConfig.JWTSecret))
	if err != nil {
		return
	}

	userRefreshToken := entity.NewUserRefreshToken(refreshToken, time.Now().Add(tds.sessionConfig.RefreshTokenExpire))

	err = tds.userTokenRepository.Store(userRefreshToken)
	if err != nil {
		return
	}

	return
}

func (tds *TokenDomainService) VerifyToken(token string) (userId string, err error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tds.sessionConfig.JWTSecret), nil
	}, jwt.WithSubject("access"), jwt.WithIssuer("tasclock"), jwt.WithExpirationRequired(), jwt.WithIssuedAt())
	if err != nil {
		return
	}

	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		err = jwt.ErrTokenInvalidClaims
		return
	}

	if claims.IssuedAt.After(time.Now()) {
		err = jwt.ErrTokenUsedBeforeIssued
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		err = jwt.ErrTokenExpired
		return
	}

	userId = claims.Audience[0]

	return
}

func (tds *TokenDomainService) RefreshToken(token string) (accessToken, refreshToken string, err error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tds.sessionConfig.JWTSecret), nil
	}, jwt.WithSubject("refresh"), jwt.WithIssuer("tasclock"), jwt.WithExpirationRequired(), jwt.WithIssuedAt())
	if err != nil {
		return
	}

	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok {
		err = jwt.ErrTokenInvalidClaims
		return
	}

	if claims.IssuedAt.After(time.Now()) {
		err = jwt.ErrTokenUsedBeforeIssued
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		err = jwt.ErrTokenExpired
		return
	}

	exist, err := tds.userTokenRepository.Exist(token)
	if err != nil {
		return
	}

	if !exist {
		err = jwt.ErrTokenExpired
		return
	}

	userRefreshToken, err := tds.userTokenRepository.Get(token)
	if err != nil {
		return
	}

	// TODO: refresh token reused handling
	if userRefreshToken.IsUsed {
		err = jwt.ErrTokenExpired
		return
	}

	userRefreshToken.Use()

	err = tds.userTokenRepository.Update(userRefreshToken)
	if err != nil {
		return
	}

	accessToken, refreshToken, err = tds.GenerateToken(claims.Audience[0])
	if err != nil {
		return
	}

	return
}

func (tds *TokenDomainService) RevokeToken(token string) error {
	userRefreshToken, err := tds.userTokenRepository.Get(token)
	if err != nil {
		return err
	}

	return tds.userTokenRepository.Delete(userRefreshToken.Id)
}
