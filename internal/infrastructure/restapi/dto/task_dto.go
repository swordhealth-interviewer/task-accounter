package dto

type TaskCreateRequest struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

type TaskCreateResponse struct {
	ID string `json:"id"`
}
