package infra

import (
	"context"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(c context.Context) (*gin.Engine, error) {
	router := gin.New()
	router.ContextWithFallback = true

	return router, nil
}