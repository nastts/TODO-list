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

type ShowAllTasksResponse struct{
	Tasks []models.Task `json:"tasks"`
}

type CompleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type CompleteTaskResponse struct {
	Text bool `json:"text"`
}

type UpdateTaskRequest struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type UpdateTaskResponse struct {
	Text bool `json:"text"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type DeleteTaskResponse struct {
	Text bool `json:"text"`
}
