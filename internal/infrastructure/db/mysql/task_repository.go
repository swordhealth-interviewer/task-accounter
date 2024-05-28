package dbmysql

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"gorm.io/gorm"
)

type Task struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`
	Title     string     `json:"title"`
	Summary   string     `json:"summary" gorm:"size:2500"`
	OwnerID   string     `json:"owner_id"`
	Status    string     `json:"status"`
	DoneAt    *time.Time `json:"done_at"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) Create(task entities.Task) (string, error) {
	taskDB := Task{
		ID:      uuid.MustParse(task.ID),
		Title:   task.Title,
		Summary: task.Summary,
		OwnerID: task.OwnerID,
		Status:  string(task.Status),
	}

	result := t.db.Create(&taskDB)
	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", errors.New("task not created")
	}

	return taskDB.ID.String(), nil
}

func (t *TaskRepository) Save(task entities.Task) (*entities.Task, error) {
	taskDB := Task{
		ID: uuid.MustParse(task.ID),
	}

	result := t.db.First(&taskDB)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("task not found")
	}

	taskDB.Title = task.Title
	taskDB.Summary = task.Summary

	if task.Status == entities.Closed {
		taskDB.Status = string(task.Status)
		taskDB.DoneAt = &task.DoneAt
	}

	result = t.db.Save(&taskDB)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("task not saved")
	}

	timeDoneAt := time.Time{}
	if taskDB.DoneAt != nil {
		timeDoneAt = *taskDB.DoneAt
	}

	task = entities.Task{
		ID:      taskDB.ID.String(),
		Title:   taskDB.Title,
		Summary: taskDB.Summary,
		OwnerID: taskDB.OwnerID,
		Status:  entities.TaskStatus(taskDB.Status),
		DoneAt:  timeDoneAt,
	}

	return &task, nil
}

func (t *TaskRepository) FindByID(id string) (*entities.Task, error) {
	taskDB := Task{
		ID: uuid.MustParse(id),
	}

	result := t.db.First(&taskDB)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("task not found")
	}

	timeDoneAt := time.Time{}
	if taskDB.DoneAt != nil {
		timeDoneAt = *taskDB.DoneAt
	}

	task := entities.Task{
		ID:      taskDB.ID.String(),
		Title:   taskDB.Title,
		Summary: taskDB.Summary,
		OwnerID: taskDB.OwnerID,
		Status:  entities.TaskStatus(taskDB.Status),
		DoneAt:  timeDoneAt,
	}

	return &task, nil
}

func (t *TaskRepository) FindAll() ([]*entities.Task, error) {
	tasksDB := []Task{}

	result := t.db.Find(&tasksDB)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("tasks not found")
	}

	tasks := []*entities.Task{}
	for _, taskDB := range tasksDB {
		timeDoneAt := time.Time{}
		if taskDB.DoneAt != nil {
			timeDoneAt = *taskDB.DoneAt
		}

		task := entities.Task{
			ID:      taskDB.ID.String(),
			Title:   taskDB.Title,
			Summary: taskDB.Summary,
			OwnerID: taskDB.OwnerID,
			Status:  entities.TaskStatus(taskDB.Status),
			DoneAt:  timeDoneAt,
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *TaskRepository) FindByUserID(userID string) ([]*entities.Task, error) {
	tasksDB := []Task{}

	result := t.db.Where("owner_id = ?", userID).Find(&tasksDB)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("tasks not found")
	}

	tasks := []*entities.Task{}
	for _, taskDB := range tasksDB {
		timeDoneAt := time.Time{}
		if taskDB.DoneAt != nil {
			timeDoneAt = *taskDB.DoneAt
		}

		task := entities.Task{
			ID:      taskDB.ID.String(),
			Title:   taskDB.Title,
			Summary: taskDB.Summary,
			OwnerID: taskDB.OwnerID,
			Status:  entities.TaskStatus(taskDB.Status),
			DoneAt:  timeDoneAt,
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *TaskRepository) Delete(id string) error {
	taskDB := Task{
		ID: uuid.MustParse(id),
	}

	result := t.db.Delete(&taskDB)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("task not deleted")
	}

	return nil
}
