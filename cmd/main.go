package main

import (
	"log"
	"net/http"

	"github.com/cslemes/heroes-cube-web/internal/database"
	"github.com/cslemes/heroes-cube-web/internal/handlers"
	"github.com/cslemes/heroes-cube-web/internal/templates"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	dbName := "heroes.db"

	if err := templates.Init(); err != nil {
		log.Fatal(err)
	}

	db, err := database.Open(dbName)
	//db, err := database.Open("./heroes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Serve static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.Get("/", handlers.Home())
	r.Get("/characters", handlers.Characters(db))

	// Start server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
