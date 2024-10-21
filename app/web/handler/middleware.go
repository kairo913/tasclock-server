package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(port string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:" + port}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	return cors.New(config)
}