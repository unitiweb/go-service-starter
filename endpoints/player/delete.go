package player

import (
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"github.com/unitiweb/go-service-starter/utils"
	"net/http"
)

// The main post struct
// Implements the EndpointInterface
type del struct{}

func DeleteInit() {
	e := del{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (d del) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodDelete,
		Path:   "/players/{id}",
	}
}

// Validate the endpoints input data
func (d del) Validate(r *http.Request, data *server.Data) []string {
	var v []string
	return v
}

// Add Endpoint Middleware (if any)
func (d del) Middleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	// Nothing to do here
}

// Handle the endpoint request
func (d del) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get id as Int from Url Query
	id := utils.GetUrlQueryInt(r, "id")
	p := db.Player{}.Find(id)
	if p.Id == nil {
		err = server.ThrowError("PlayerNotFound")
		return p, err
	}

	// Store the player in the database
	db.Conn.Delete(&p)

	return p, err
}
