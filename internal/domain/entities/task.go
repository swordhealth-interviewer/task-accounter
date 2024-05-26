package entities

import (
	"errors"
	"time"
)

type TaskStatus string

const (
	Open                  TaskStatus = "open"
	Closed                TaskStatus = "closed"
	summaryMaxLength                 = 2500
	errorSummaryMaxLength            = "summary must have a maximum of 2500 characters"
	errorTitleRequired               = "title is required"
	errorSummaryRequired             = "summary is required"
)

type Task struct {
	ID      string
	Title   string
	Summary string
	Owner   string
	Status  TaskStatus
	DoneAt  time.Time
}

func NewTask(title string, summary string, owner string) (Task, error) {
	err := ValidateTaskParameters(title, summary)
	if err != nil {
		return Task{}, err
	}

	task := Task{
		Title:   title,
		Summary: summary,
		Owner:   owner,
		Status:  Open,
	}

	return task, nil
}

func ValidateTaskParameters(title string, summary string) error {
	if title == "" {
		return errors.New(errorTitleRequired)
	}

	if summary == "" {
		return errors.New(errorSummaryRequired)
	}

	if len(summary) > summaryMaxLength {
		return errors.New(errorSummaryMaxLength)
	}

	return nil
}
