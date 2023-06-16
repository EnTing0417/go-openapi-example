package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/swaggo/http-swagger"
	
)

// @title Authentication Service API
// @version 1.0
// @description This is a sample authentication service API using Swagger.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/auth", authenticate)
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// @summary Authenticates a user
// @description Authenticates a user based on the provided credentials.
// @tags authentication
// @accept json
// @produce json
// @param credentials body Credentials true "User credentials"
// @success 200 {string} string "Successfully authenticated"
// @failure 400 {string} string "Invalid request"
// @router /auth [post]
func authenticate(w http.ResponseWriter, r *http.Request) {
	// Implement your authentication logic here

	w.Write([]byte("Successfully authenticated"))
}