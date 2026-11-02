package main

import (
	"to-do-list/internal/service"
	"to-do-list/internal/cli"
)

func main(){
	todoService := service.NewToDoService()
	todoCLI := cli.NewTodoCLI(todoService)

	todoCLI.Run()
}