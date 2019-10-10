package server

import (
	"fmt"
	"log"
	"net/http"
)

type ValidateError struct {
	Data []string
}

// Starts the server and listens on the configured port
func Listen() {

	// If the env hasn't been loaded yet, then load it now
	if Config.EnvLoaded == false {
		LoadEnv()
	}

	// Add some system errors
	AddError("RouteNotFound", 404, "The route does not exist")
	AddError("ValidationError", 422, "Input validation error")

	// Add the response middleware
	AddMiddleware(responseMiddleware)

	// Load all the configured endpoints
	loadEndpoints()

	// Load all the configured middleware
	loadMiddleware()

	// A catch all for 404 route not found error
	catchAllNonRoutes()

	// Build and log the address
	address := fmt.Sprintf(":%s", Config.Port)
	logStartMessage(address)

	// Start the server and listen to port
	log.Fatal(http.ListenAndServe(address, router))
} /**/

func logStartMessage(a string) {
	fmt.Println("")
	fmt.Println("--------------------------------------------")
	fmt.Printf("The server is running at %s\n", a)
	fmt.Println("--------------------------------------------")
	fmt.Println("")
}
