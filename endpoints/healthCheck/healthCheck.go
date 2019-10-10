package healthCheck

import (
	"github.com/unitiweb/go-service-starter/server"
	"net/http"
)

type healthCheck struct{}

type success struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Init() {
	server.AddError("ItemNotFound", 404, "The requested item could not be found")
	e := healthCheck{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (hc healthCheck) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodGet,
		Path:   "/health-check",
	}
}

// Add Endpoint Middleware (if any)
func (hc healthCheck) Middleware(w http.ResponseWriter, r *http.Request, next http.Handler) {
	// Nothing to do here
}

// Validate the endpoints input data
func (hc healthCheck) Validate(r *http.Request, data *server.Data) []string {
	var err []string
	// Validation will go here
	return err
}

// Handle the endpoint request
func (hc healthCheck) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var res success
	var err error

	//err = server.ThrowError("ItemNotFound")
	//return res, err

	res = success{
		Code:    "Success",
		Message: "Service is running and healthy",
	}

	return res, err
}
