package service

import (
	"fmt"

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

func (s *TodoService) CreateTask(title string) uuid.UUID{//done
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

func (s *TodoService) ShowAllTasks(){//done
	fmt.Println(s.tasks)
}

func ShowAllTasksInProgress(tasks []models.Task){//done
	for i, task := range tasks {
		if tasks[i].Progress == "in progress"{
			fmt.Println(task)
		}
	}
}

func (s *TodoService) ShowAllTasksDone(){//done
	for i, task := range s.tasks {
		if s.tasks[i].Progress == "done"{
			fmt.Println(task)
		}
	}
}

func (s *TodoService) CompleteTaskByID(id uuid.UUID){
	for i := range s.tasks{
		if s.tasks[i].ID == id{
			s.tasks[i].Progress = "done"
			fmt.Printf("the task: %s is done!!\n", s.tasks[i].Name)
			return
		}
	}
	fmt.Println("Task was not found\n")
}

func (s *TodoService) UpdateTaskByID(id uuid.UUID, newName string){ //done
	for i := range s.tasks{
		if s.tasks[i].ID == id{

			s.tasks[i].Name = newName

			fmt.Printf("the task: %s was updated!!\n", s.tasks[i].Name)
			return
		}
	}
	fmt.Println("Task was not found\n")
}

func (s *TodoService) DeleteTaskByID(id uuid.UUID){
	for i := range s.tasks{
		if s.tasks[i].ID == id{
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			fmt.Printf("the task was deleted!!\n")
			return
		}
	}
	fmt.Println("Task was not found\n")
}

func CheckID(id string) (uuid.UUID, error){
	parsedID, err := uuid.Parse(id)
	if err != nil{
		return uuid.Nil, err
	}
	return parsedID, nil
}
