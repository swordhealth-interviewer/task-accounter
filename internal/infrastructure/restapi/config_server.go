package restapi

import "github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/handlers"

func configHandlers() *AppHandlers {
	pingHandler := handlers.NewPingHandler()

	return &AppHandlers{
		pingHandler: pingHandler,
	}
}
