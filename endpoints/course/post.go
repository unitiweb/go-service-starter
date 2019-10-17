package course

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"github.com/unitiweb/go-service-starter/utils"
	"net/http"
	"strings"
)

// The main post struct
// Implements the EndpointInterface
type coursePost struct{}

func PostInit() {
	e := coursePost{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (pt coursePost) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodPost,
		Path:   "/courses",
	}
}

// Validate the endpoints input data
func (pt coursePost) Validate(r *http.Request, data *server.Data) []string {
	p := course{
		courseName: data.Body["courseName"].(string),
		city: data.Body["city"].(string),
		state: data.Body["state"].(string),
		zip: data.Body["zip"].(string),
	}

	err := validation.ValidateStruct(&p,
		validation.Field(&p.courseName, validation.Required),
		validation.Field(&p.city, validation.Required),
		validation.Field(&p.state, validation.Required, validation.Length(2, 2)),
		validation.Field(&p.zip, validation.Required),
	)

	return utils.ParseValidationError(err)
}

// Handle the endpoint request
func (pt coursePost) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Create a new course and populate the data
	nc := db.Course{
		CourseName: utils.ToStringPointer(data.Body["courseName"]),
		City: utils.ToStringPointer(data.Body["city"]),
		State: utils.ToStringPointer(data.Body["state"]),
		Zip: utils.ToStringPointer(data.Body["zip"]),
	}

	*nc.State = strings.ToUpper(*nc.State)

	// Store the course in the database
	db.Conn.Create(&nc)

	return nc, err
}
