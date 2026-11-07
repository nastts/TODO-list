package http

import (
	"to-do-list/internal/service"
)

type TodoHTTP struct{
	TodoService *service.TodoService
}

func NewTodoHTTP(service *service.TodoService) TodoHTTP{
	return TodoHTTP{
		TodoService: service,
	}
}