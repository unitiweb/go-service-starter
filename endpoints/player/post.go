package player

import (
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"github.com/unitiweb/go-service-starter/utils"
	"net/http"
)

// The main post struct
// Implements the EndpointInterface
type post struct{}

func PostInit() {
	e := post{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (pt post) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodPost,
		Path:   "/players",
	}
}

// Validate the endpoints input data
func (pt post) Validate(r *http.Request, data *server.Data) []string {
	return validateInput(data)
}

// Add Endpoint Middleware (if any)
func (pt post) Middleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	// Nothing to do here
}

// Handle the endpoint request
func (pt post) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Create a new player and populate the data
	np := db.Player{
		PlayerName: utils.ToStringPointer(data.Body["playerName"]),
	}

	// Store the player in the database
	db.Conn.Create(&np)

	return np, err
}
