package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

// @summary Authenticates a user
// @description Authenticates a user based on the provided credentials.
// @tags authentication
// @accept json
// @produce json
// @param credentials body Credentials true "User credentials"
// @success 200 {string} string "Successfully authenticated"
// @failure 400 {string} string "Invalid request"
// @router /auth [post]

type Credentials struct {
	User string `json:user,bson:user`
	Password string `json:password,bson:password`
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	// Implement your authentication logic here

	w.Write([]byte("Successfully authenticated"))
}

func routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/auth", authenticate)

	return router
}
