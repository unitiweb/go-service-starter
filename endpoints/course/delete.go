package course

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
		Path:   "/courses/{id}",
	}
}

// Validate the endpoints input data
func (d del) Validate(r *http.Request, data *server.Data) []string {
	var v []string
	return v
}

// Handle the endpoint request
func (d del) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get id as Int from Url Query
	id := utils.GetUrlQueryInt(r, "id")
	p := db.Course{}.Find(id)
	if p.Id == nil {
		err = server.ThrowError(NotFound)
		return p, err
	}

	// Store the course in the database
	db.Conn.Delete(&p)

	return p, err
}
