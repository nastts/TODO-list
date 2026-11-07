package httptransport

import (
	"to-do-list/internal/task"

	"github.com/google/uuid"
)


type CreateTaskRequest struct{
	Name string `json:"name"`
}

type CreateTaskResponse struct{
	ID uuid.UUID `json:"id"`
}

type ShowAllTasksInProgressRequest struct{
	InProgress string `json:"pogress"`
}

type ShowAllTasksDoneRequest struct{
	Done string `json:"progress"`
}

type ShowAllTasksResponse struct{
	Tasks []models.Task `json:"tasks"`
}

type CompleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type CompleteTaskResponse struct {
	Text string `json:"text"`
}

type UpdateTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type UpdateTaskResponse struct {
	Text string `json:"text"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type DeleteTaskResponse struct {
	tasksext string `json:"text"`
}
