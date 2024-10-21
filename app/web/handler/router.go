package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/util/config"
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

	NewUserHandler(c, router.Group("/user"), sqlHandler)

	return router, nil
}