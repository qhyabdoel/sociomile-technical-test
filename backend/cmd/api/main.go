package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/config"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/handler"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/middleware"
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

	// public routes
	r.Post("/channel/webhook", convHandler.HandleWebhook)

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware("your-secret-key"))

		// conversations routes
		r.Get("/conversations", convHandler.List)
		r.Get("/conversations/{id}", convHandler.GetDetail)
		r.Post("/conversations/{id}/messages", convHandler.Reply)

		// Ticket Routes (Hanya Admin) [cite: 80, 95]
		r.Group(func(r chi.Router) {
			r.Use(middleware.RoleMiddleware("admin"))
			r.Patch("/tickets/{id}/status", ticketHandler.UpdateStatus)
		})
	})

	http.ListenAndServe(":8080", r)
}
