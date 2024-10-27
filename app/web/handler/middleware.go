package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	csrf "github.com/utrack/gin-csrf"
)

type TokenClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func CORSMiddleware(port string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:" + port}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Authorization", "Content-Type", "X-CSRF-Token"}
	return cors.New(config)
}

func CSRFMiddleware(secret string) gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: secret,
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "CSRF token mismatch"})
			c.Abort()
		},
	})
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

		token, err := jwt.ParseWithClaims(t, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			c.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		claims, ok := token.Claims.(*TokenClaims)
		if !ok || claims.UserId == "" {
			c.Header("WWW-Authenticate", "Bearer error=\"invalid_token\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			c.Header("WWW-Authenticate", "Bearer error=\"expired_token\"")
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}

		c.Set("userId", claims.UserId)

		c.Next()
	}
}
