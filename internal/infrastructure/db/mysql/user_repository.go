package dbmysql

import (
	"errors"

	"github.com/google/uuid"
	"github.com/uiansol/task-accounter.git/internal/domain/entities"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	Username string    `json:"username" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Role     string    `json:"role"`
	Password string    `json:"password"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindByUsername(username string) (*entities.User, string, error) {
	var userFound User
	ur.db.Where("username=?", username).Find(&userFound)
	if userFound.ID == uuid.Nil {
		return nil, "", errors.New("User not found")
	}

	user := entities.User{
		ID:    userFound.ID.String(),
		Name:  userFound.Username,
		Email: userFound.Email,
		Role:  entities.UserRole(userFound.Role),
	}

	return &user, userFound.Password, nil
}
