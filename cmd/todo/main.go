package main

import (
	"errors"
	"log"
	"net/http"
	/*"to-do-list/internal/service"
	"to-do-list/internal/transport/cli"*/
	"to-do-list/internal/transport/httptransport"
)

func main(){
	/*todoService := service.NewToDoService()
	todoCLI := cli.NewTodoCLI(todoService)

	todoCLI.Run()*/

	server := httptransport.NewServer()
	err := server.RegisterHandlers()
	if err != nil{
		log.Printf("error register handlers: %s", err)
		return
	}

	log.Printf("server started successfully!! port: :8080")
	log.Printf("POST /api/task/create")
	if err := server.Start(); !errors.Is(err, http.ErrServerClosed){
		log.Printf("error start server: %s", err)
		return
	}
}