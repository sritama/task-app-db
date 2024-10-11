package api

import (
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   int64  `json:"createdAt"`
}

func (s *Service) Insert(description string) (*Task, error) {

	newTask := Task{
		ID:          uuid.NewString(),
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now().Unix(),
	}

	_, err := s.TaskDB.Exec("INSERT INTO tasks (ID, DESCRIPTION, CHECKED, CREATED_AT) VALUES (?, ?, ?, ?)", newTask.ID, newTask.Description, 0, newTask.CreatedAt)

	if err != nil {
		fmt.Printf("Error inserting data:\n %s", err)
		return nil, err
	}

	fmt.Println("Data inserted successfully")

	return &newTask, nil
}

func (s *Service) GetAllTasks() ([]Task, error) {

	rows, err := s.TaskDB.Query("SELECT * FROM TASKS")
	if err != nil {
		fmt.Printf("error in querying data from database\n %s", err)
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		var fetchedTask Task
		if err := rows.Scan(&fetchedTask.ID, &fetchedTask.Description, &fetchedTask.Completed, &fetchedTask.CreatedAt); err != nil {
			fmt.Printf("error in querying data from database\n %s", err)
			return nil, err
		}
		// fmt.Printf("Fetched Task: %v\n", Task)

		tasks = append(tasks, fetchedTask)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CreatedAt < tasks[j].CreatedAt
	})

	return tasks, nil
}

func (s *Service) Check(id string, completed bool) (*Task, error) {

	var checked int
	if completed {
		checked = 1
	} else {
		checked = 0
	}

	_, err := s.TaskDB.Exec("UPDATE tasks set CHECKED = ? where ID = ?", checked, id)

	if err != nil {
		fmt.Printf("Error updating data:\n %s", err)
		return nil, err
	}

	fmt.Println("Data updated successfully")

	rows, err := s.TaskDB.Query("SELECT * from tasks where ID = ?", id)
	if err != nil {
		fmt.Printf("Error querying data after update:\n %s", err)
		return nil, err
	}
	defer rows.Close()

	var updatedTask Task
	for rows.Next() {
		if err := rows.Scan(&updatedTask.ID, &updatedTask.Description, &updatedTask.Completed, &updatedTask.CreatedAt); err != nil {
			fmt.Printf("error in querying data after update\n %s", err)
			return nil, err
		}
	}

	fmt.Println("Data fetched successfully after update")

	return &updatedTask, nil

}
