package dto

import "time"

type TaskCreateRequest struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

type TaskCreateResponse struct {
	ID string `json:"id"`
}

type TaskReadResponse struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	Summary string    `json:"summary"`
	OwerID  string    `json:"owner_id"`
	Status  string    `json:"status"`
	DoneAt  time.Time `json:"done_at"`
}

type TaskReadAllResponse struct {
	Tasks []TaskReadResponse `json:"tasks"`
}

type TaskUpdateRequest struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	CloseTask bool   `json:"close_task"`
}
