package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/service"
)

type UserController struct {
	userAppService *service.UserAppService
}

func NewUserController(userAppService *service.UserAppService) *UserController {
	return &UserController{userAppService}
}

func (uc *UserController) SignUp(c *gin.Context) {
	var req SignUpRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userAppService.CreateUser(req.Lastname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res SignUpResponse
	res.Id = user.Id.String()
	res.Lastname = user.Lastname

	c.JSON(http.StatusCreated, res)
}
