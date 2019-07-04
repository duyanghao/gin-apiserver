package main

import (
	"fmt"

	"github.com/duyanghao/GinApiServer/pkg/config"
	"github.com/duyanghao/GinApiServer/pkg/log"
	"github.com/duyanghao/GinApiServer/pkg/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	m := config.GetString(config.FLAG_KEY_GIN_MODE)
	gin.SetMode(m)

	route.InstallRoutes(r)
	serverBindAddr := fmt.Sprintf("%s:%d", config.GetString(config.FLAG_KEY_SERVER_HOST), config.GetInt(config.FLAG_KEY_SERVER_PORT))
	log.Infof("Run server at %s", serverBindAddr)
	r.Run(serverBindAddr) // listen and serve
}
