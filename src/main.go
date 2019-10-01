package main

import (
	"fmt"
	"github.com/unitiweb/go-microservice/src/endpoints"
	"github.com/unitiweb/go-microservice/src/server"
)

func main() {

	// Load the .env (or server environment) server variables and populate server.Config
	server.LoadEnv()

	//port := os.Getenv("PORT")
	fmt.Println("port", server.Config.Port)

	// Add the endpoints to the server
	endpoints.AddAll()

	// Start the server
	server.Listen()
}
