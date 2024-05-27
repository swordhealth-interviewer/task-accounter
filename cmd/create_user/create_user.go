package main

import (
	"github.com/google/uuid"
	dbmysql "github.com/uiansol/task-accounter.git/internal/infrastructure/db/mysql"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	restapi.LoadEnvs()
	db := restapi.ConnectToMysql()

	username := "manager-2"
	password := "manager-2"
	role := "manager"

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := dbmysql.User{
		ID:       uuid.New(),
		Username: username,
		Password: string(passwordHash),
		Role:     role,
	}

	db.Create(&user)
}
