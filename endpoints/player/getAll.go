package player

import (
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"net/http"
)

// The main put struct
// Implements the EndpointInterface
type getAll struct{}

func GetAllInit() {
	e := getAll{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (g getAll) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodGet,
		Path:   "/players",
	}
}

// Validate the endpoints input data
func (g getAll) Validate(r *http.Request, data *server.Data) []string {
	var v []string
	return v
}

// Handle the endpoint request
func (g getAll) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get all players
	all := db.Player{}.FindAll()

	return all, err
}
