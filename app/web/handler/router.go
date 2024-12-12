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

func SetUpRouter(c context.Context) (*gin.Engine, error) {
	router := gin.Default()
	router.ContextWithFallback = true

	sqlHandler, err := database.NewSqlHandler(c)
	if err != nil {
		return nil, err
	}

	cfg := config.NewClientConfig(c)

	router.Use(CORSMiddleware(cfg.Port))

    authController := controller.NewAuthController(service.NewUserAppService(c, repository.NewUserRepository(sqlHandler)))

	router.POST("/signup", authController.SignUp)
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout)
	router.POST("/refresh", authController.Refresh)

	NewUserHandler(c, router.Group("/users"), sqlHandler)

	return router, nil
}
