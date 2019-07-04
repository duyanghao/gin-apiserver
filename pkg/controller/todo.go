package controller

import (
	"github.com/duyanghao/GinApiServer/pkg/models"
	"github.com/duyanghao/GinApiServer/pkg/service"
	"github.com/gin-gonic/gin"
)

type ToDoController interface {
	GetToDo(c *gin.Context)
}

func NewToDoController() ToDoController {
	return &toDoController{
		toDoService: service.NewToDoService(),
	}
}

type toDoController struct {
	toDoService service.ToDoService
}

// @Summary GetToDo
// @Description GetToDo
// @Success 200 {object} models.Response OK
// @Failure 400 {object} models.Response Bad Request
// @Failure 401 {object} models.Response Unauthorized
// @Failure 403 {object} models.Response Forbidden
// @Failure 500 {object} models.Response Internal Server Error
// @router /todo/get [get]
func (this *toDoController) GetToDo(c *gin.Context) {
	this.toDoService.Get()
	c.JSON(200, models.Response{Code: 0, Message: "todo demo", Data: struct{}{}})
	return
}
