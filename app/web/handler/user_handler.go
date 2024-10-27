package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/core/service"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/infra/repository"
	"github.com/kairo913/tasclock-server/app/web/controller"
	csrf "github.com/utrack/gin-csrf"
)

func NewUserHandler(ctx context.Context, userGroup *gin.RouterGroup, sqlHandler *database.SqlHandler) {
	userController := controller.NewUserController(service.NewUserAppService(ctx, repository.NewUserRepository(sqlHandler)))

	userGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"token": csrf.GetToken(c)})
	})
	userGroup.POST("/signup", userController.SignUp)
	userGroup.POST("/signin", userController.SignIn)
}
