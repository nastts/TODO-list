package main

import (
	"fmt"

	"github.com/google/uuid"
)

func CreateTask(title string) (Task, uuid.UUID){//done
	var progress string
	progress = "in progress"
	id := uuid.New()

	newTask := Task{
		ID: id,
		Name: title,
		Progress: progress,
	}
	return newTask, id
}

func ShowAllTasks(tasks []Task){//done
	fmt.Println(tasks)
}
func ShowAllTasksInProgress(tasks []Task){//done
	for i, task := range tasks {
		if tasks[i].Progress == "in progress"{
			fmt.Println(task)
		}
	}
}
func ShowAllTasksDone(tasks []Task){//done
	for i, task := range tasks {
		if tasks[i].Progress == "done"{
			fmt.Println(task)
		}
	}
}

func CompleteTaskByID(tasks []Task, id uuid.UUID){
	for i := range tasks{
		if tasks[i].ID == id{
			tasks[i].Progress = "done"
			fmt.Printf("the task: %s is done!!\n", tasks[i].Name)
			return
		}
	}
	fmt.Println("Task was not found\n")
}

func UpdateTaskByID(tasks []Task, id uuid.UUID, newName string){ //done
	for i := range tasks{
		if tasks[i].ID == id{

			tasks[i].Name = newName

			fmt.Printf("the task: %s was updated!!\n", tasks[i].Name)
			return
		}
	}
	fmt.Println("Task was not found\n")
}

func DeleteTaskByID(tasks *[]Task, id uuid.UUID){
	for i := range *tasks{
		if (*tasks)[i].ID == id{
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			fmt.Printf("the task was deleted!!\n")
			return
		}
	}
	fmt.Println("Task was not found\n")
}

func checkID(id string) (uuid.UUID, error){
	parsedID, err := uuid.Parse(id)
	if err != nil{
		return uuid.Nil, err
	}
	return parsedID, nil
}
