package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/service"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/infra/repository"
	"github.com/kairo913/tasclock-server/app/util/config"
	"github.com/kairo913/tasclock-server/app/web/controller"
)

func NewUserHandler(ctx context.Context, userGroup *gin.RouterGroup, sqlHandler *database.SqlHandler) {
	userAppService := service.NewUserAppService(ctx, repository.NewUserRepository(sqlHandler))
	userController := controller.NewUserController(userAppService, service.NewTokenDomainService(ctx, repository.NewUserRefreshTokenRepository(sqlHandler)))

	sessonConfig := config.NewSessionConfig(ctx)

	userGroup.Use(AuthMiddleware(sessonConfig.JWTSecret))
	userGroup.Use(UserMiddleware(userAppService))

	userGroup.GET("/:id", userController.Get)
	userGroup.PUT("/:id", userController.Put)
	userGroup.DELETE("/:id", userController.Delete)
	userGroup.PUT("/:id/email", userController.UpdateEmail)
	userGroup.PUT("/:id/password", userController.UpdatePassword)

	NewTaskHandler(ctx, userGroup.Group("/:id/tasks"), sqlHandler)
}
