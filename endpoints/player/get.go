package player

import (
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"github.com/unitiweb/go-service-starter/utils"
	"net/http"
)

// The main put struct
// Implements the EndpointInterface
type get struct{}

func GetInit() {
	e := get{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (g get) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodGet,
		Path:   "/players/{id}",
	}
}

// Validate the endpoints input data
func (g get) Validate(r *http.Request, data *server.Data) []string {
	var v []string
	return v
}

// Handle the endpoint request
func (g get) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get id as Int from Url Query
	id := utils.GetUrlQueryInt(r, "id")
	p := db.Player{}.Find(id)
	if p.Id == nil {
		err = server.ThrowError(NotFound)
		return p, err
	}

	return p, err
}
