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

	users := []string{"tech-1", "tech-2", "manager-1"}
	passwords := []string{"tech-1", "tech-2", "manager-1"}
	roles := []string{"techinician", "technician", "manager"}

	for i := 0; i < 3; i++ {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(passwords[i]), bcrypt.DefaultCost)

		user := dbmysql.User{
			ID:       uuid.New(),
			Username: users[i],
			Email:    users[i] + "@task-accounter.com",
			Role:     roles[i],
			Password: string(passwordHash),
		}

		db.Create(&user)
	}
}
