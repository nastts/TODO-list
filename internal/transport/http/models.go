package http

import (
	"to-do-list/internal/task"

	"github.com/google/uuid"
)


type CreateTaskRequest struct{
	name string
}

type CreateTaskResponse struct{
	ID uuid.UUID
}

type ShowAllTasksInProgressRequest struct{
	inProgress string
}

type ShowAllTasksDoneRequest struct{
	done string
}

type ShowAllTasksResponse struct{
	tasks []models.Task
}

type CompleteTaskRequest struct {
	ID uuid.UUID
}

type CompleteTaskResponse struct {
	text string
}

type UpdateTaskRequest struct {
	ID uuid.UUID
}

type UpdateTaskResponse struct {
	text string
}

type DeleteTaskRequest struct {
	ID uuid.UUID
}

type DeleteTaskResponse struct {
	text string
}
