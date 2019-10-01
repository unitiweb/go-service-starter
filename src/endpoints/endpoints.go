package endpoints

import (
	"github.com/unitiweb/go-microservice/src/endpoints/healthCheck"
	"github.com/unitiweb/go-microservice/src/server"
)

// Add the endpoints to the server
func AddAll() {
	server.AddEndpoint(healthCheck.Endpoint())
}
