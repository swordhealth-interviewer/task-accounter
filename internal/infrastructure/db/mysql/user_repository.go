package dbmysql

import (
	"github.com/google/uuid"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Username string    `json:"username" gorm:"unique"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}

type UserRepository struct {
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindByUsername(username string) (*entities.User, string, error) {
	return nil, "", nil
}
