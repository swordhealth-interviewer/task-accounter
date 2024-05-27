package restapi

import "github.com/uiansol/task-accounter.git/internal/infrastructure/restapi/handlers"

func configHandlers() *AppHandlers {
	loginHandler := handlers.NewLoginHandler()
	pingHandler := handlers.NewPingHandler()

	return &AppHandlers{
		loginHandler: loginHandler,
		pingHandler:  pingHandler,
	}
}
