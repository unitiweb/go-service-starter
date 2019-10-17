package endpoints

import (
	"github.com/unitiweb/go-service-starter/endpoints/course"
	"github.com/unitiweb/go-service-starter/endpoints/healthCheck"
	"github.com/unitiweb/go-service-starter/endpoints/player"
)

// Add the endpoints to the server
func AddAll() {
	// Load the healthCheck endpoints
	healthCheck.Init()

	// Load the player endpoints
	player.Init()
	course.Init()
}
