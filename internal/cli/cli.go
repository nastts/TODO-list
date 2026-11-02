package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"to-do-list/internal/service"
)

type TodoCLI struct{
	TodoService *service.TodoService
}

func NewTodoCLI(service *service.TodoService) *TodoCLI{
	return &TodoCLI{
		TodoService: service,
	}
}

func (c *TodoCLI) Run(){
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
			id := c.TodoService.CreateTask(action)
			fmt.Printf("ID: %v\n", id)
		
		case "show":
			switch action {
			case "done":
				c.TodoService.ShowAllTasksDone()
			case "in-progress":
				c.TodoService.ShowAllTasksInProgress()
			default:
				c.TodoService.ShowAllTasks()
			}
		
		case "complete":
			parsedID, err := service.CheckID(action)
			if err != nil{
				log.Fatal(err)
			}
			c.TodoService.CompleteTaskByID(parsedID)
		
		case "update":
			parsedID, err := service.CheckID(action)
			if err != nil{
				log.Fatal(err)
			}
			readerForUpdate := bufio.NewReader(os.Stdin)
			newName, err := readerForUpdate.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			correctedNewName := strings.TrimSpace(newName)
			c.TodoService.UpdateTaskByID(parsedID, correctedNewName)
		
		case "delete":
			parsedID, err := service.CheckID(action)
			if err != nil{
				log.Fatal(err)
			}
			c.TodoService.DeleteTaskByID(parsedID)
		default:
			fmt.Println("The function was not found. Try again")
		}
	}
}