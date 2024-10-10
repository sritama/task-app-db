package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

type Service struct {
	router chi.Router
	list   taskLists
}

func (s *Service) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payload CreatePayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// fmt.Printf("The request body is %+v\n", payload)
		if payload.Description == "" {
			http.Error(w, "no task description provided", http.StatusBadRequest)
			return
		}

		response := CreateTaskResponse{
			Task: s.list.Create(payload.Description),
		}
		// fmt.Printf("The response is %+v\n", response)

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "error converting create response to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)

	}
}

func (s *Service) GetTaskList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Getting the list of tasks")
		response := ListResponse{
			Tasks: s.list.GetList(),
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
			return
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.Write(jsonResponse)
	}
}

func (s *Service) DeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/tasks/")
		s.list.Delete(id)
	}
}

func (s *Service) CheckTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// fmt.Printf("The url path is %v\n", r.URL.Path)

		id := strings.TrimPrefix(r.URL.Path, "/tasks/")

		var payload CheckPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		taskResponse, err := s.list.Check(id, payload.Completed)
		if err != nil {
			http.Error(w, "error in checking task", http.StatusInternalServerError)
			return
		}
		response := CheckTaskResponse{
			Task: taskResponse,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "error converting create response to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}
