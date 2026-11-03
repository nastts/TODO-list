package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"

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

func convertToJSON(data any){
	jsonData, err := json.Marshal(data)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line := strings.TrimSpace(input)
	return line
}

func checkID(id string) (uuid.UUID, error){
	parsedID, err := uuid.Parse(id)
	if err != nil{
		return uuid.Nil, err
	}
	return parsedID, nil
}

func (c *TodoCLI) Run(){
	working := true
	var action string

	for working{
		line := readInput()
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
			task := c.TodoService.CreateTask(action)
			convertToJSON(task) // TODO: implement only ID, not the whole task
		
		case "show":
			switch action {
			case "done":
				tasksDone := c.TodoService.ShowAllTasksDone()
				convertToJSON(tasksDone)
			case "in-progress":
				tasksInProgress := c.TodoService.ShowAllTasksInProgress()
				convertToJSON(tasksInProgress)
			default:
				tasks := c.TodoService.ShowAllTasks()
				convertToJSON(tasks)
			}
		
		case "complete":
			parsedID, err := checkID(action)
			if err != nil{
				log.Fatal(err)
			}
			isCompleted := c.TodoService.CompleteTaskByID(parsedID)
			if isCompleted{
				fmt.Printf("done!!\n")
			}else{
				fmt.Println("not found\n")
			}
		
		case "update":
			parsedID, err := checkID(action)
			if err != nil{
				log.Fatal(err)
			}
			newName := readInput()
			isUpdated := c.TodoService.UpdateTaskByID(parsedID, newName)
			if isUpdated{
				fmt.Printf("updated!!\n")
			}else{
				fmt.Println("not found\n")
			}
		
		case "delete":
			parsedID, err := checkID(action)
			if err != nil{
				log.Fatal(err)
			}
			isDeleted := c.TodoService.DeleteTaskByID(parsedID)
			if isDeleted{
				fmt.Printf("deleted!!\n")
			}else{
				fmt.Println("not found\n")
			}
		default:
			fmt.Println("The function was not found. Try again\n")
		}
	}
}