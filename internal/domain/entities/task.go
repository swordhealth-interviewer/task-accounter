package entities

import "time"

type TaskStatus string

const (
	Open   TaskStatus = "open"
	Closed TaskStatus = "closed"
)

type Task struct {
	ID      string
	Title   string
	Summary string
	Owner   string
	Status  TaskStatus
	DoneAt  time.Time
}

func NewTask(title string, summary string, owner string) Task {
	return Task{
		Title:   title,
		Summary: summary,
		Owner:   owner,
		Status:  Open,
	}
}
