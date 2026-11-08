package httptransport

import (
	"encoding/json"
	"log"
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

func decodeBody(w http.ResponseWriter, r *http.Request, request any) error{
	if err := json.NewDecoder(r.Body).Decode(request); err != nil{
		http.Error(w, "Invalid request body", 400)
		return err
	}
	return nil
}

func returnGoodResponse(w http.ResponseWriter, response any) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(response); err != nil{
		http.Error(w, "Invalid response", 500)
		return err
	}
	return nil
}

func (h *TodoHTTP) CreateTask(w http.ResponseWriter, r *http.Request){
	var request CreateTaskRequest

	defer r.Body.Close()

	if err := decodeBody(w, r, &request); err != nil{
		return
	}

	if request.Name == ""{
		http.Error(w, "Title is empty", 400)
		return
	}

	id := h.TodoService.CreateTask(request.Name)

	response := CreateTaskResponse{
		ID: id,
	}

	if err := returnGoodResponse(w,response); err != nil{
		return
	}
}

func (h *TodoHTTP) ShowAllTasks(w http.ResponseWriter, r *http.Request){
	var tasks []models.Task
	
	tasks = h.TodoService.ShowAllTasks()

	response := ShowAllTasksResponse{
		Tasks: tasks,
	}

	if err := returnGoodResponse(w,response); err != nil{
		return
	}
}

func (h *TodoHTTP) ShowAllTasksInProgress(w http.ResponseWriter, r *http.Request){
	var tasks []models.Task

	tasks = h.TodoService.ShowAllTasksInProgress()
	response := ShowAllTasksResponse{
		Tasks: tasks,
	}

	if err := returnGoodResponse(w,response); err != nil{
		return
	}
}

func (h *TodoHTTP) ShowAllTasksDone(w http.ResponseWriter, r *http.Request){
	var tasks []models.Task

	tasks = h.TodoService.ShowAllTasksDone()
	response := ShowAllTasksResponse{
		Tasks: tasks,
	}

	if err := returnGoodResponse(w,response); err != nil{
		return
	}
}


func (h *TodoHTTP) CompleteTask(w http.ResponseWriter, r *http.Request){
	var request CompleteTaskRequest
	var response CompleteTaskResponse

	if err := decodeBody(w, r, &request); err != nil{
		return
	}
	complete := h.TodoService.CompleteTaskByID(request.ID)

	response = CompleteTaskResponse{
			Text: complete,
		}
	
	if err := returnGoodResponse(w,response); err != nil{
		return
	}
}

func (h *TodoHTTP) UpdateTask(w http.ResponseWriter, r *http.Request){
	var request UpdateTaskRequest
	var response UpdateTaskResponse

	if err := decodeBody(w, r, &request); err != nil{
		return
	}
	log.Printf("the name: %s", request.Name)
	update := h.TodoService.UpdateTaskByID(request.ID, request.Name)

	response = UpdateTaskResponse{
			Text: update,
		}
	
	if err := returnGoodResponse(w,response); err != nil{
		return
	}
}

func (h *TodoHTTP) DeleteTask(w http.ResponseWriter, r *http.Request){
	var request DeleteTaskRequest
	var response DeleteTaskResponse

	if err := decodeBody(w, r, &request); err != nil{
		return
	}
	update := h.TodoService.DeleteTaskByID(request.ID)

	response = DeleteTaskResponse{
			Text: update,
		}
	
	if err := returnGoodResponse(w,response); err != nil{
		return
	}
}