package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (s *Service) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payload CreatePayload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if payload.Description == "" {
			http.Error(w, "no task description provided", http.StatusBadRequest)
			return
		}

		task, err := s.Insert(payload.Description)
		if err != nil {
			http.Error(w, "Error inserting data in DB", http.StatusInternalServerError)
			return
		}

		response := CreateTaskResponse{
			Task: task,
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

func (s *Service) GetTaskList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Getting the list of tasks")
		tasks, err := s.GetAllTasks()
		if err != nil {
			http.Error(w, "Error receiving data from DB", http.StatusInternalServerError)
			return
		}
		response := ListResponse{
			Tasks: tasks,
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

		_, err := s.TaskDB.Exec("DELETE FROM tasks WHERE ID = ?", id)

		if err != nil {
			http.Error(w, "Error deleting data", http.StatusInternalServerError)
			return
		}
		fmt.Println("Data deleted successfully")

	}
}

func (s *Service) CheckTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := strings.TrimPrefix(r.URL.Path, "/tasks/")

		var payload CheckPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		taskResponse, err := s.Check(id, payload.Completed)
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
