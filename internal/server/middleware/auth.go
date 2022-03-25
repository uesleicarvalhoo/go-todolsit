package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uesleicarvalhoo/go-todolist/internal/config"
	"github.com/uesleicarvalhoo/go-todolist/pkg/auth"
	"github.com/uesleicarvalhoo/go-todolist/pkg/utils"
)

func validateSession(ctx context.Context, token string) (string, error) {
	return auth.ValidateJWT(token, config.GetEnv().SecretKey)
}

func AuthenticationMiddleware(function func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, utils.ErrAuthHeaderNotFound)
			return
		}

		userId, err := validateSession(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}

		c.Request.Header.Set("X-User-Id", userId)
		function(c)
	}
}
