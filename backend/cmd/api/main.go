package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/config"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/handler"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/repository"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/service"
)

func main() {
	// init db and repositories
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	convRepo := repository.NewConversationRepository(db)

	// init services
	convService := service.NewConversationService(convRepo)
	
	// init handlers
	convHandler := handler.NewConversationHandler(convService)

	// setup router
	r := chi.NewRouter()
}

