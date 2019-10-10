package player

import (
	"fmt"
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

// Add Endpoint Middleware (if any)
func (g get) Middleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	// Nothing to do here
}

// Handle the endpoint request
func (g get) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get id as Int from Url Query
	id := utils.GetUrlQueryInt(r, "id")
	p := db.Player{}.Find(id)
	fmt.Println("p", p)
	if p.Id == nil || p.DeletedAt != nil {
		err = server.ThrowError("PlayerNotFound")
		return p, err
	}

	return p, err
}
