package service

import (
	"log"

	"github.com/google/uuid"

	"to-do-list/internal/task"
)


type TodoService struct{
	tasks []models.Task
}

func NewToDoService() *TodoService{
	return &TodoService{
		tasks: []models.Task{},
	}
}

func (s *TodoService) CreateTask(title string) uuid.UUID{
	var progress string
	progress = "in progress"
	id := uuid.New()

	newTask := models.Task{
		ID: id,
		Name: title,
		Progress: progress,
	}
	s.tasks = append(s.tasks, newTask)
	return id
}

func (s *TodoService) ShowAllTasks() []models.Task {
	return s.tasks
}

func (s *TodoService) ShowAllTasksInProgress() []models.Task {
	var tasksInProgress []models.Task
	for i, task := range s.tasks {
		if s.tasks[i].Progress == "in progress"{
			tasksInProgress = append(tasksInProgress, task)
		}
	}
	return tasksInProgress
}

func (s *TodoService) ShowAllTasksDone() []models.Task {
	var tasksDone []models.Task
	for i, task := range s.tasks {
		if s.tasks[i].Progress == "done"{
			tasksDone = append(tasksDone, task)
		}
	}
	return tasksDone
}

func (s *TodoService) CompleteTaskByID(id uuid.UUID) bool {
	for i := range s.tasks{
		if s.tasks[i].ID == id{
			s.tasks[i].Progress = "done"
			log.Println("service: found")
			return true
		}
	}
	log.Println("service: not found")
	return false
}

func (s *TodoService) UpdateTaskByID(id uuid.UUID, newName string) bool{
	for i := range s.tasks{
		if s.tasks[i].ID == id{
			s.tasks[i].Name = newName
			return true
		}
	}
	return false
}

func (s *TodoService) DeleteTaskByID(id uuid.UUID) bool{
	for i := range s.tasks{
		if s.tasks[i].ID == id{
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}
	return false
}
