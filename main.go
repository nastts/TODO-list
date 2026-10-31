package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	tasks := []Task{}
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
			task, id := CreateTask(action)
			tasks = append(tasks, task)
			fmt.Printf("ID: %v\n", id)
		
		case "show":
			ShowAllTasks(tasks)
		
		case "complete":
			parsedID, err := checkID(action)
			if err != nil{
				return
			}
			CompleteTaskByID(tasks, parsedID)
		
		case "update":
			parsedID, err := checkID(action)
			if err != nil{
				return
			}
			UpdateTaskByID(tasks, parsedID)
		
		case "delete":
			parsedID, err := checkID(action)
			if err != nil{
				return
			}
			DeleteTaskByID(&tasks, parsedID)
		default:
			fmt.Println("The function was not found. Try again")
		}
	}
}