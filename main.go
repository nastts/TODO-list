package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct{
	ID uuid.UUID
	Name string
	Description string
	Progress string
}


func CreateTask() (Task, uuid.UUID){
	var name string
	var description string
	var progress string

	fmt.Println("the title: ")
	fmt.Scan(&name)
	fmt.Println("the description (optional) :")
	fmt.Scan(&description)
	progress = "in progress"
	id := uuid.New()

	newTask := Task{
		ID: id,
		Name: name,
		Description: description,
		Progress: progress,
	}
	return newTask, id
}

func ShowAllTasks(tasks []Task){
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func CompleteTaskByID(tasks []Task, id uuid.UUID){
	for i := range tasks{
		if tasks[i].ID == id{
			tasks[i].Progress = "done"
			fmt.Printf("the task: %s is done!!", tasks[i].Name)
			return
		}
	}
	fmt.Println("Task was not found")
}

func UpdateTaskByID(tasks []Task, id uuid.UUID){
	for i := range tasks{
		if tasks[i].ID == id{
			fmt.Println("the title: ")
			fmt.Scan(&tasks[i].Name)
			fmt.Println("the description (optional) :")
			fmt.Scan(&tasks[i].Description)
			fmt.Printf("the task: %s was updated!!", tasks[i].Name)
			return
		}
	}
	fmt.Println("Task was not found")
}

func DeleteTaskByID(tasks []Task, id uuid.UUID){
	for i := range tasks{
		if tasks[i].ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("the task: %s was deleted!!", tasks[i].Name)
			return
		}
	}
	fmt.Println("Task was not found")
}

func checkID() (uuid.UUID, error){
	var id string
	fmt.Println("Type ID here: ")
	fmt.Scan(&id)
	parsedID, err := uuid.Parse(id)
	if err != nil{
		return uuid.Nil, err
	}
	return parsedID, nil
}

func main(){
	tasks := []Task{}
	working := true
	var input int
	for working{
		fmt.Println("What would you like to do?")
		fmt.Println("(1) create task")
		fmt.Println("(2) see all tasks")
		fmt.Println("(3) complete task by ID")
		fmt.Println("(4) update task by ID")
		fmt.Println("(5) delete task by ID")
		fmt.Println("(0) exit")
		fmt.Scan(&input)
		switch input{
		case 0:
			working = false
		case 1:
			task, id := CreateTask()
			tasks = append(tasks, task)
			fmt.Printf("ID: %d", id)
		case 2:
			ShowAllTasks(tasks)
		case 3:
			parsedID, err := checkID()
			if err != nil{
				return
			}
			CompleteTaskByID(tasks, parsedID)
		case 4:
			parsedID, err := checkID()
			if err != nil{
				return
			}
			UpdateTaskByID(tasks, parsedID)
		case 5:
			parsedID, err := checkID()
			if err != nil{
				return
			}
			DeleteTaskByID(tasks, parsedID)
		}
	}
}