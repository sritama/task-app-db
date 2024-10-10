package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

func NewService() *Service {

	l := newTaskList()
	service := &Service{
		router: chi.NewRouter(),
		list:   l,
	}
	return service
}

func (s *Service) Routes() {

	c := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Origin", "X-Requested-With", "Content-Type", "Accept"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	})

	s.router.Use(c.Handler)

	s.router.Get("/tasks", s.GetTaskList())
	s.router.Post("/tasks", s.CreateTask())
	s.router.Delete("/tasks/{id}", s.DeleteTask())
	s.router.Put("/tasks/{id}", s.CheckTask())
}
