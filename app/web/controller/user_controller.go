package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/kairo913/tasclock-server/app/core/service"
)

type UserController struct {
	userAppService *service.UserAppService
	tokenDomainService *service.TokenDomainService
}

func NewUserController(userAppService *service.UserAppService, tokenDomainService *service.TokenDomainService) *UserController {
	return &UserController{userAppService, tokenDomainService}
}

func (uc *UserController) SignUp(c *gin.Context) {
	var req SignUpRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist, err := uc.userAppService.ExistByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	user, err := uc.userAppService.CreateUser(req.Lastname, req.Firstname, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res SignUpResponse
	res.Id = user.Id.String()
	res.Lastname = user.Lastname
	res.Firstname = user.Firstname
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt

	c.JSON(http.StatusCreated, res)
}

func (uc *UserController) SignIn(c *gin.Context) {
	var req SignInRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exist, err := uc.userAppService.ExistByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !exist {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user, err := uc.userAppService.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if !uc.userAppService.VerifyPassword(user, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	accessToken, refreshToken, err := uc.tokenDomainService.GenerateToken(user.Id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", "Bearer " + accessToken)

	c.SetCookie("refresh_token", refreshToken, uc.tokenDomainService.GetRefreshTokenAge(), "/", "", false, true)

	c.Status(http.StatusNoContent)
}

func (uc *UserController) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := uc.tokenDomainService.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Authorization", "Bearer " + accessToken)

	c.SetCookie("refresh_token", refreshToken, uc.tokenDomainService.GetRefreshTokenAge(), "/", "", false, true)

	c.Status(http.StatusNoContent)
}

func (uc *UserController) SignOut(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.tokenDomainService.RevokeToken(refreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
