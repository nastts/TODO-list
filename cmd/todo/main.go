package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"to-do-list/internal/task"
	"to-do-list/internal/service"
)

func main(){
	tasks := []models.Task{}
	working := true
	reader := bufio.NewReader(os.Stdin)
	var action string

	for working{
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		line := strings.TrimSpace(input)
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}
		commandWord := words[0]

		action = ""
		action = strings.Join(words[1:], " ")

		switch commandWord{
		case "exit":
			working = false
		
		case "create":
			task, id := service.CreateTask(action)
			tasks = append(tasks, task)
			fmt.Printf("ID: %v\n", id)
		
		case "show":
			switch action {
			case "done":
				service.ShowAllTasksDone(tasks)
			case "in progress":
				service.ShowAllTasksInProgress(tasks)
			default:
				service.ShowAllTasks(tasks)
			}
		
		case "complete":
			parsedID, err := service.CheckID(action)
			if err != nil{
				return
			}
			service.CompleteTaskByID(tasks, parsedID)
		
		case "update":
			parsedID, err := service.CheckID(action)
			if err != nil{
				return
			}
			readerForUpdate := bufio.NewReader(os.Stdin)
			newName, err := readerForUpdate.ReadString('\n')
			if err != nil {
				return
			}
			correctedNewName := strings.TrimSpace(newName)
			service.UpdateTaskByID(tasks, parsedID, correctedNewName)
		
		case "delete":
			parsedID, err := service.CheckID(action)
			if err != nil{
				return
			}
			service.DeleteTaskByID(&tasks, parsedID)
		default:
			fmt.Println("The function was not found. Try again")
		}
	}
}