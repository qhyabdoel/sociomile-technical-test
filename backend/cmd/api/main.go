package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/config"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/handler"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/middleware"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/repository"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/service"

	_ "github.com/qhyabdoel/sociomile-technical-test/backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Sociomile Technical Test API
// @version         1.0
// @description     API for multi-tenant conversation and ticketing system.
// @host            localhost:8080
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// init db and repositories
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	convRepo := repository.NewConversationRepository(db)
	msgRepo := repository.NewMessageRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	tenantRepo := repository.NewTenantRepository(db)
	userRepo := repository.NewUserRepository(db)

	// init services
	convService := service.NewConversationService(convRepo, msgRepo, tenantRepo)
	ticketService := service.NewTicketService(ticketRepo, convRepo)

	// init handlers
	jwtSecret := os.Getenv("JWT_SECRET")
	authHandler := handler.NewAuthHandler(userRepo, jwtSecret)
	convHandler := handler.NewConversationHandler(convService)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// setup router
	r := chi.NewRouter()

	// public routes
	r.Post("/channel/webhook", convHandler.HandleWebhook)
	r.Post("/login", authHandler.Login)

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(jwtSecret))

		// conversations routes
		r.Get("/conversations", convHandler.List)
		r.Get("/conversations/{id}", convHandler.GetDetail)
		r.Post("/conversations/{id}/messages", convHandler.Reply)

		// ticket routes

		// agent and admin can create ticket
		r.Post("/tickets", ticketHandler.Create)

		// only admin can update status
		r.Group(func(r chi.Router) {
			r.Use(middleware.RoleMiddleware("admin"))
			r.Patch("/tickets/{id}/status", ticketHandler.UpdateStatus)
		})
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	http.ListenAndServe(":8080", r)
}
