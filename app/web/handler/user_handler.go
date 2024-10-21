package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/infra/repository"
	"github.com/kairo913/tasclock-server/app/core/service"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/web/controller"
)

func NewUserHandler(userGroup *gin.RouterGroup, sqlHandler *database.SqlHandler) {
	userController := controller.NewUserController(service.NewUserAppService(repository.NewUserRepository(sqlHandler)))

	userGroup.POST("/signup", userController.SignUp)
}