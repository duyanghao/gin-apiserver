package controller

import (
	"github.com/duyanghao/GinApiServer/pkg/config"
	"github.com/gin-gonic/gin"
)

func Version(c *gin.Context) {
	c.JSON(200, config.FLAG_KEY_SERVER_VERSION)
}
