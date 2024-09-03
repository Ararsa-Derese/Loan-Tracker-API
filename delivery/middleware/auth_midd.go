package middleware

import (
	"loan/internal/tokenutil"
	// "fmt"
	"loan/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMidd(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Err:     nil,
			Message: "Authorization header is required",
		})
		c.Abort()
		return
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Err:     nil,
			Message: "Token is required",
		})
		c.Abort()
		return
	}

	claims, err := tokenutil.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Err:     err,
			Message: "Invalid token",
		})
		c.Abort()
		return
	}

	c.Set("claim", *claims)
	c.Next()
}
