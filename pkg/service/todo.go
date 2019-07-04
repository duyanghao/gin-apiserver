package service

type ToDoService interface {
	Get()
}

type toDoService struct {
}

func NewToDoService() ToDoService {
	return &toDoService{}
}

func (c *toDoService) Get() {
	return
}
