package dbmysql

import "github.com/uiansol/task-accounter.git/internal/domain/entities"

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (t *TaskRepository) Save(task entities.Task) (*entities.Task, error) {
	return nil, nil
}

func (t *TaskRepository) FindByID(id string) (*entities.Task, error) {
	return nil, nil
}

func (t *TaskRepository) FindAll() ([]*entities.Task, error) {
	return nil, nil
}

func (t *TaskRepository) FindByUserID(userID string) ([]*entities.Task, error) {
	return nil, nil
}

func (t *TaskRepository) Delete(id string) error {
	return nil
}
