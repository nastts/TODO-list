package main

import "github.com/google/uuid"


type Task struct{
	ID uuid.UUID
	Name string
	//Description string
	Progress string
}