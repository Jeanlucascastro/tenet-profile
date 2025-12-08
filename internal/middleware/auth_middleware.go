package middleware

import (
	service "tenet-profile/internal/client"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authClient *service.AuthClient
}

func NewAuthMiddleware(authClient *service.AuthClient) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}

func (m *AuthMiddleware) MiddlewareFunc() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		valid, err := m.authClient.ValidateToken(token)
		if err != nil || !valid {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
