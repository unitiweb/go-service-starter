package server

import (
	"net/http"
)

// The structure of an Endpoint
type Endpoint struct {
	Method   string
	Path     string
	Validate func(r *http.Request) []string
	Handle   func(r *http.Request) (interface{}, error)
	Config   func()
}

// A slice that will hold all configured endpoints
var endpoints []Endpoint

// Add an endpoint to the server
func AddEndpoint(e Endpoint) {
	endpoints = append(endpoints, e)
}

// Add endpoints to the server
func loadEndpoints() {

	// Iterate through the endpoints and add to server
	for _, e := range endpoints {

		// Run config if one exists
		if e.Config != nil {
			e.Config()
		}

		// Build the endpoint handler function
		f := func(w http.ResponseWriter, r *http.Request) {

			// If a validator exists then validate
			if e.Validate != nil {
				v := e.Validate(r)
				if len(v) != 0 {
					e := FindError("ValidationError")
					e.Data = v
					HandleError(w, e)
					return
				}
			}

			// Run tha handler and display any errors if they exist
			res, err := e.Handle(r)
			if err != nil {
				e := FindError(err.Error())
				HandleError(w, e)
				return
			}

			// At this point there were not validation errors or handler errors so return the response
			ToJson(w, res)
		}

		// Add the handler to the router
		router.HandleFunc(e.Path, f).Methods(http.MethodOptions, e.Method)
	}
}
