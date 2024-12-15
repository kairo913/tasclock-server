package handler

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kairo913/tasclock-server/app/core/service"
)

func CORSMiddleware(port string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:" + port}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "X-CSRF-Token"}
	return cors.New(config)
}

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		t := strings.TrimPrefix(header, "Bearer ")
		if t == "" {
			c.Header("WWW-Authenticate", "Bearer error=\"token_required\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		token, err := jwt.ParseWithClaims(t, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			c.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok {
			c.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			c.Header("WWW-Authenticate", "Bearer error=\"expired_token\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		c.Next()
	}
}

func UserMiddleware(userAppService *service.UserAppService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
        if err != nil {
            c.Status(http.StatusBadRequest)
            c.Abort()
        }

		user, err := userAppService.GetUser(id)
		if err == sql.ErrNoRows {
			c.Status(http.StatusNotFound)
			c.Abort()
		}

		if err != nil {
			c.Status(http.StatusInternalServerError)
			c.Abort()
		}

		c.Set("user", user)

		c.Next()
	}
}
