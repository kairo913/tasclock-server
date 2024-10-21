package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/service"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/infra/repository"
	"github.com/kairo913/tasclock-server/app/web/controller"
)

func NewUserHandler(ctx context.Context, userGroup *gin.RouterGroup, sqlHandler *database.SqlHandler) {
	userController := controller.NewUserController(service.NewUserAppService(ctx, repository.NewUserRepository(sqlHandler)))

	userGroup.POST("/signup", userController.SignUp)
}
