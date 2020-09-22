package middleware

import (
	"github.com/duyanghao/GinApiServer/pkg/config"
	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		config.GetString(config.FLAG_KEY_AUTH_USERNAME): config.GetString(config.FLAG_KEY_AUTH_PASSWORD),
	})
}
