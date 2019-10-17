package server

import (
	"github.com/unitiweb/go-service-starter/utils"
	"net/http"
)

type EndpointConfig struct {
	Method string
	Path string
}

type EndpointInterface interface {
	Config() EndpointConfig
	Validate(r *http.Request, data *Data) []string
	Handle(r *http.Request, data *Data) (interface{}, error)
	//Middleware(w http.ResponseWriter, r *http.Request, next http.Handler) // Not implemented yet
}

type Data struct {
	UrlParams map[string]interface{}
	Body      map[string]interface{}
}

// A slice that will hold all configured endpoints
var endpoints []EndpointInterface

// Add an endpoint to the server
func AddEndpoint(e EndpointInterface) {
	endpoints = append(endpoints, e)
}

// Add endpoints to the server
func loadEndpoints() {
	// Iterate through the endpoints and add to server
	for _, e := range endpoints {
		// Add the handler to the router
		c := e.Config()
		handler := buildHandler(e)
		router.HandleFunc(c.Path, handler).Methods(http.MethodOptions, c.Method)
	}
}

// Build the endpoint handler
func buildHandler(e EndpointInterface) func(w http.ResponseWriter, r *http.Request) {

	// Build the endpoint handler function
	return func(w http.ResponseWriter, r *http.Request) {

		var data Data
		getBody(r, &data)

		// If a validator exists then validate
		valid := checkValidity(w, r, e, &data)
		if valid == false {
			return
		}

		// Run the handler and display any errors if they exist
		res, valid := runHandler(w, r, e, &data)
		if valid == false {
			return
		}

		// At this point there were not validation errors or handler errors so return the response
		ToJson(w, res)
	}
}

// Get the post, put, or patch body data and populate the Data struct on respected method types
func getBody(r *http.Request, d *Data) {
	if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch {
		err := utils.GetPostBody(r, &d.Body)
		if err != nil {
			panic(err.Error())
		}
	}
}

// Run the validation function if it exists so input validation can be validated
// If validation errors exist then handle the errors and return false
func checkValidity(w http.ResponseWriter, r *http.Request, e EndpointInterface, d *Data) bool {
	if e.Validate != nil {
		v := e.Validate(r, d)
		if len(v) != 0 {
			e := FindError("ValidationError")
			e.Data = v
			HandleError(w, e)
			return false
		}
	}

	return true
}

func runHandler(w http.ResponseWriter, r *http.Request, e EndpointInterface, data *Data) (interface{}, bool) {
	res, err := e.Handle(r, data)
	if err != nil {
		e := FindError(err.Error())
		HandleError(w, e)
		return res, false
	}

	return res, true
}
