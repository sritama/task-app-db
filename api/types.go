package api

type ListResponse struct {
	Tasks []Task `json:"tasks"`
}

type CreatePayload struct {
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	Task *Task `json:"Task"`
}

type CheckPayload struct {
	Completed bool `json:"completed"`
}

type CheckTaskResponse struct {
	Task *Task `json:"Task"`
}
