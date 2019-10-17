package course

import (
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"net/http"
)

// The main put struct
// Implements the EndpointInterface
type getAllCourses struct{}

func GetAllInit() {
	e := getAllCourses{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (g getAllCourses) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodGet,
		Path:   "/courses",
	}
}

// Validate the endpoints input data
func (g getAllCourses) Validate(r *http.Request, data *server.Data) []string {
	var v []string
	// Nothing to do here
	return v
}

// Handle the endpoint request
func (g getAllCourses) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get all courses
	all := db.Course{}.FindAll()

	return all, err
}
