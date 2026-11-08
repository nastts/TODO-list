package httptransport

import (
	"net/http"
	"to-do-list/internal/service"
	
)


type Server struct{
	srv *http.Server
	mux *http.ServeMux
}

func NewServer() *Server{
	mux := http.DefaultServeMux
	
	srv := &http.Server{
		Addr: ":8080",
	}

	return &Server{
		srv: srv,
		mux: mux,
	}
}

func (s *Server) RegisterHandlers() error{
	service := service.NewToDoService()
	transport := NewTodoHTTP(service)

	s.mux.Handle("/api/task/create", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost{
			http.Error(w, "Method is not allowed", 405)
			return
		}
		transport.CreateTask(w,r)
	}))
	
	s.mux.Handle("/api/tasks", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet{
			http.Error(w, "Method is not allowed", 405)
			return
		}
		transport.ShowAllTasks(w,r)
	}))

	s.mux.Handle("/api/tasks/inprogress", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet{
			http.Error(w, "Method is not allowed", 405)
			return
		}
		transport.ShowAllTasksInProgress(w,r)
	}))

	s.mux.Handle("/api/tasks/done", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet{
			http.Error(w, "Method is not allowed", 405)
			return
		}
		transport.ShowAllTasksDone(w,r)
	}))

	s.mux.Handle("/api/tasks/complete", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost{
			http.Error(w, "Method is not allowed", 405)
			return
		}
		transport.CompleteTask(w, r)
	}))

	s.mux.Handle("/api/tasks/update", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost{
			http.Error(w, "Method is not allowed", 405)
			return
		}
		transport.UpdateTask(w, r)
	}))

	s.mux.Handle("/api/tasks/delete", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete{
			http.Error(w, "Method is not allowed", 405)
			return
		}
		transport.DeleteTask(w, r)
	}))

	return nil
}

func (s *Server) Start() error{
	return s.srv.ListenAndServe()
}