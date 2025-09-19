package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rafrey/dndhub/internal/db"
	"github.com/rafrey/dndhub/internal/franchise"
	"github.com/rafrey/dndhub/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	conn := db.Connect()

	// modules
	userRepo := user.NewRepository(conn)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	franchiseRepo := franchise.NewRepository(conn)
	franchiseService := franchise.NewService(franchiseRepo)
	franchiseHandler := franchise.NewHandler(franchiseService)

	// root router
	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// public
	r.Mount("/users", userHandler.Routes())

	// protected
	r.Group(func(protected chi.Router) {
		protected.Use(user.AuthMiddleware)
		protected.Mount("/franchise", franchiseHandler.Routes())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("DnDHub API running on :%s\n", port)
	http.ListenAndServe(":"+port, r)
}
