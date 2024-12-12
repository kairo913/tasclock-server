package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/entity"
	"github.com/kairo913/tasclock-server/app/core/service"
)

type UserController struct {
	userAppService     *service.UserAppService
	tokenDomainService *service.TokenDomainService
}

func NewUserController(userAppService *service.UserAppService, tokenDomainService *service.TokenDomainService) *UserController {
	return &UserController{userAppService, tokenDomainService}
}

func (uc *UserController) Get(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	var res GetUserResponse
	res.Lastname = user.Lastname
	res.Firstname = user.Firstname
	res.CreatedAt = user.CreatedAt
	res.UpdatedAt = user.UpdatedAt

	c.JSON(http.StatusOK, res)
}

func (uc *UserController) Put(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := c.MustGet("user").(*entity.User)

	if user.Lastname == req.Lastname && user.Firstname == req.Firstname {
		c.Status(http.StatusBadRequest)
		return
	}

	err := uc.userAppService.UpdateUser(user, req.Lastname, req.Firstname)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (uc *UserController) Delete(c *gin.Context) {
	user := c.MustGet("user").(*entity.User)

	err := uc.userAppService.DeleteUser(user.Id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (uc *UserController) UpdateEmail(c *gin.Context) {
	var req UpdateUserEmailRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := c.MustGet("user").(*entity.User)

	if user.Email == req.Email {
		c.Status(http.StatusBadRequest)
		return
	}

	err := uc.userAppService.UpdateEmail(user, req.Email)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "Email already used" {
			status = http.StatusConflict
		}
		c.Status(status)
		return
	}

	c.Status(http.StatusNoContent)
}

func (uc *UserController) UpdatePassword(c *gin.Context) {
	var req UpdateUserPasswordRequest
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := c.MustGet("user").(*entity.User)

	err := uc.userAppService.UpdatePassword(user, req.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
