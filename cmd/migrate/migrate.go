package main

import (
	dbmysql "github.com/uiansol/task-accounter.git/internal/infrastructure/db/mysql"
	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi"
)

func main() {
	restapi.LoadEnvs()
	db := restapi.ConnectToMysql()
	db.AutoMigrate(&dbmysql.User{})
}
