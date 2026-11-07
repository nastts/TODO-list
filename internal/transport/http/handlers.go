package http

import (
	"encoding/json"
	"net/http"
	"to-do-list/internal/service"
	models "to-do-list/internal/task"
)

type TodoHTTP struct{
	TodoService *service.TodoService
}

func NewTodoHTTP(service *service.TodoService) TodoHTTP{
	return TodoHTTP{
		TodoService: service,
	}
}

func (h *TodoHTTP) CreateTask(w http.ResponseWriter, r *http.Request){
	var createTaskReq CreateTaskRequest

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&createTaskReq); err != nil{
		http.Error(w, "Invalid request body", 400)
		return
	}

	if createTaskReq.name == ""{
		http.Error(w, "Title is empty", 400)
		return
	}

	id := h.TodoService.CreateTask(createTaskReq.name)

	createTaskResp := CreateTaskResponse{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(createTaskResp)
}

func (h *TodoHTTP) ShowAllTasks(w http.ResponseWriter, r *http.Request){
	//depending on showing all task or in progress or done

	var showAllTasksDoneReq ShowAllTasksDoneRequest
	var showAllTasksInProgressReq ShowAllTasksInProgressRequest
	var tasks []models.Task

	if err := json.NewDecoder(r.Body).Decode(&showAllTasksDoneReq); err != nil{
		http.Error(w, "Invalid request body", 400)
		return
	}
	
	if err := json.NewDecoder(r.Body).Decode(&showAllTasksInProgressReq); err != nil{
		http.Error(w, "Invalid request body", 400)
		return
	}

	if showAllTasksDoneReq.done == "done" {
		tasks = h.TodoService.ShowAllTasksDone()
	}else if showAllTasksInProgressReq.inProgress == "in-progress"{
		tasks = h.TodoService.ShowAllTasksInProgress()
	}else{
		tasks = h.TodoService.ShowAllTasks()
	}

	showAllTasksResp := ShowAllTasksResponse{
		tasks: tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(showAllTasksResp)
}