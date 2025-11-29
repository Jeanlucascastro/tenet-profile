package middleware

import (
	service "tenet-profile/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService *service.AuthService
}

func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) MiddlewareFunc() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		valid, err := m.authService.ValidateToken(token)
		if err != nil || !valid {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
