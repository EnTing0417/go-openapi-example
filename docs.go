// +build !release

package main

import (
	"docs"
	"github.com/swaggo/http-swagger"
)

// @title Authentication Service API
// @version 1.0
// @description This is a sample authentication service API using Swagger.
// @BasePath /api/v1
func init() {
	httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"), // The location of the Swagger JSON file
	)
}
