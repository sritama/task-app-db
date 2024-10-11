package api

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

type Service struct {
	router chi.Router
	TaskDB *sql.DB
}

func NewService() *Service {

	db, err := initDB()
	if err != nil {
		log.Panicf("error in starting sqlite server \n%s", err)
	}

	service := &Service{
		router: chi.NewRouter(),
		TaskDB: db,
	}

	return service
}

func initDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./db/tasks.db")
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to tasks DB")

	// Create the tasks table if it doesn't exist
	_, err = db.Exec(`
	   CREATE TABLE IF NOT EXISTS tasks (
	       ID TEXT PRIMARY KEY,
	       DESCRIPTION TEXT,
	       CHECKED INTEGER,
	       CREATED_AT INTEGER
	   )
	`)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully created the tasks table")

	return db, nil
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
