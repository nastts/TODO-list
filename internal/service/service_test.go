package service_test

import (
	"testing"

	"github.com/google/uuid"

	"to-do-list/internal/service"
)

var testService service.TodoService

func TestCreateRoom(t *testing.T){
	tests := []struct{
        name string
		expectedTitle string
    }{
		{
			name: "create task",
			expectedTitle: "shopping",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			id := service.NewToDoService().CreateTask(test.expectedTitle)

			if id == uuid.Nil{
				t.Errorf("Expected not nil ID")
			}
		})
	}
}

func TestTaskByID(t *testing.T){
	t.Run("complete task successfully", func(t *testing.T) {
		id := testService.CreateTask("test-task-name")

		isCompleted := testService.CompleteTaskByID(id)
		if !isCompleted {
			t.Errorf("CompleteTaskByID() returned false, expect true")
		}

		isUpdated := testService.UpdateTaskByID(id, "test-task-update-name")
		if !isUpdated {
			t.Errorf("UpdateTaskByID() returned false, expect true")
		}

		isDeleted := testService.DeleteTaskByID(id)
		if !isDeleted {
			t.Errorf("DeleteTaskByID() returned false, expect true")
		}
	})
}

func TestUnexistedTaskByID(t *testing.T){
	t.Run("complete unexisted task", func(t *testing.T) {
		id := uuid.New()

		isCompleted := testService.CompleteTaskByID(id)
		if isCompleted {
			t.Errorf("CompleteTaskByID() returned true, expect false")
		}

		isUpdated := testService.UpdateTaskByID(id, "test-task-name")
		if isUpdated {
			t.Errorf("UpdateTaskByID() returned true, expect false")
		}

		isDeleted := testService.DeleteTaskByID(id)
		if isDeleted {
			t.Errorf("DeleteTaskByID() returned true, expect false")
		}
	})
}