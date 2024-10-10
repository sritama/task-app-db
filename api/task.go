package api

import (
	"fmt"
	"github.com/google/uuid"
	"sort"
	"time"
)

type task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   int64  `json:"createdAt"`
}

type taskLists struct {
	tasks map[string]*task
}

type TaskLists interface {
	Create(description string) *task
	Check(id string, completed bool) (*task, error)
	Delete(id string)
	GetList() []task
}

func newTaskList() taskLists {

	list := taskLists{
		tasks: make(map[string]*task),
	}
	list.AddInitialTasks()
	return list
}
func (l *taskLists) AddInitialTasks() {

	id1 := uuid.NewString()
	l.tasks[id1] = &task{
		ID:          id1,
		Description: "Pick up dry cleaning",
		Completed:   true,
		CreatedAt:   1,
	}

	id2 := uuid.NewString()
	l.tasks[id2] = &task{
		ID:          id2,
		Description: "Grab Coffee",
		Completed:   false,
		CreatedAt:   2,
	}

	id3 := uuid.NewString()
	l.tasks[id3] = &task{
		ID:          id3,
		Description: "Medical appointment",
		Completed:   false,
		CreatedAt:   3,
	}
}

func (l *taskLists) Create(description string) *task {

	id := uuid.NewString()
	newTask := &task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now().Unix(),
	}

	l.tasks[id] = newTask
	fmt.Printf("Sucessfully created new task %v\n", newTask)
	return newTask

}

func (l *taskLists) Check(id string, completed bool) (*task, error) {
	l.tasks[id].Completed = completed
	return l.tasks[id], nil
}

func (l *taskLists) Delete(id string) {
	delete(l.tasks, id)
}

func (l *taskLists) GetList() []task {

	// fmt.Printf("Number of tasks %d\n", len(l.tasks))
	tasks := make([]task, 0)

	for _, t := range l.tasks {
		tasks = append(tasks, *t)
		// fmt.Printf("Task Description %s\n", t.Description)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].CreatedAt < tasks[j].CreatedAt
	})

	return tasks
}
