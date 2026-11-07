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
	return nil
}