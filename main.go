package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/endpoints"
	"github.com/unitiweb/go-service-starter/env"
	"github.com/unitiweb/go-service-starter/server"
)

type Player struct {
	gorm.Model
	PlayerName string
}

func main() {
	// Load Environmet Variables
	env.LoadEnv()

	// Create the database connection
	db.Connect()
	defer db.Close()

	// Migrate the database if any
	db.Migrate()

	// Add an error handler to the server
	server.AddError("RequestBodyError", 500, "Error reading request body")

	// Add the endpoints to the server
	endpoints.AddAll()

	// Start the server
	server.Listen(
		server.Config{
			Port: env.Get("PORT"),
		},
	)
}
