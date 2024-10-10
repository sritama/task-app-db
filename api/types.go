package api

type ListResponse struct {
	Tasks []task `json:"tasks"`
}

type CreatePayload struct {
	Description string `json:"description"`
}

type CreateTaskResponse struct {
	Task *task `json:"task"`
}

type CheckPayload struct {
	Completed bool `json:"completed"`
}

type CheckTaskResponse struct {
	Task *task `json:"task"`
}
