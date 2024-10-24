package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kairo913/tasclock-server/app/infra/database"
	"github.com/kairo913/tasclock-server/app/util/config"
	csrf "github.com/utrack/gin-csrf"
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

	router.Use(csrf.Middleware(csrf.Options{
		Secret: cfg.CSRFSecret,
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "CSRF token mismatch"})
			c.Abort()
		},
	}))

	NewUserHandler(c, router.Group("/user"), sqlHandler)

	return router, nil
}
