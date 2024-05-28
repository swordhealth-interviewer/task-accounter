package main

import (
	"fmt"
	"os"

	"github.com/uiansol/task-accounter.git/internal/infrastructure/restapi"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "debug" {
		fmt.Println("Starting local server!")
		restapi.StartServer(true)
		return
	} else {
		debug := false
		restapi.MigrateDB(debug)
		restapi.PopulateDB(debug)
		restapi.StartServer(debug)
		return
	}
}
