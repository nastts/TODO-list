package main

import (
	"to-do-list/internal/service"
	"to-do-list/internal/transport/cli"
)

func main(){
	todoService := service.NewToDoService()
	todoCLI := cli.NewTodoCLI(todoService)

	todoCLI.Run()
}