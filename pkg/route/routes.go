package route

import (
	"os"

	_ "github.com/duyanghao/gin-apiserver/docs"
	"github.com/duyanghao/gin-apiserver/pkg/controller"
	"github.com/duyanghao/gin-apiserver/pkg/log"
	"github.com/duyanghao/gin-apiserver/pkg/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger gin-apiserver
// @version 0.1.0
// @description This is a gin-apiserver.
// @contact.name duyanghao
// @contact.url https://duyanghao.github.io
// @contact.email 1294057873@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func InstallRoutes(r *gin.Engine) {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// a ping api test
	r.GET("/ping", controller.Ping)

	// get gin-apiserver version
	r.GET("/version", controller.Version)

	// config reload
	r.Any("/-/reload", func(c *gin.Context) {
		log.Info("===== Server Stop! Cause: Config Reload. =====")
		os.Exit(1)
	})

	rootGroup := r.Group("/api/v1")
	// AuthRequired middleware that provide basic auth
	rootGroup.Use(middleware.BasicAuthMiddleware())

	{
		// a ping api to test basic auth
		rootGroup.GET("/ping", controller.Ping)
	}

	{
		toDoController := controller.NewToDoController()
		rootGroup.GET("/todo/get", toDoController.GetToDo)
	}
}
