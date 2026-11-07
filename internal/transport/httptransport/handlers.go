package httptransport

import (
	"encoding/json"
	"net/http"
	"to-do-list/internal/service"
	"to-do-list/internal/task"
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

	if createTaskReq.Name == ""{
		http.Error(w, "Title is empty", 400)
		return
	}

	id := h.TodoService.CreateTask(createTaskReq.Name)

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

	if showAllTasksDoneReq.Done == "done" {
		tasks = h.TodoService.ShowAllTasksDone()
	}else if showAllTasksInProgressReq.InProgress == "in-progress"{
		tasks = h.TodoService.ShowAllTasksInProgress()
	}else{
		tasks = h.TodoService.ShowAllTasks()
	}

	showAllTasksResp := ShowAllTasksResponse{
		Tasks: tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(showAllTasksResp)
}

func (s *Server) Start() error{
	return s.srv.ListenAndServe()
}