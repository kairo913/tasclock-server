package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/service"
)

type AuthController struct {
	userAppService *service.UserAppService
}

func NewAuthController(userAppService *service.UserAppService) *AuthController {
	return &AuthController{userAppService}
}

func (ac *AuthController) SignUp(c *gin.Context) {

}

func (ac *AuthController) Login(c *gin.Context) {

}

func (ac *AuthController) Logout(c *gin.Context) {

}

func (ac *AuthController) Refresh(c *gin.Context) {

}
