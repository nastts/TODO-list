package main

import "github.com/google/uuid"


type Task struct{
	ID uuid.UUID `json:"ID"`
	Name string `json:"name"`
	Progress string `json:"progress"`
}