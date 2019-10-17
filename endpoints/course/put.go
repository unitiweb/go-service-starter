package course

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"github.com/unitiweb/go-service-starter/utils"
	"net/http"
)

// The main put struct
// Implements the EndpointInterface
type coursePut struct{}

func PutInit() {
	e := coursePut{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (pt coursePut) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodPut,
		Path:   "/courses/{id}",
	}
}

// Validate the endpoints input data
func (pt coursePut) Validate(r *http.Request, data *server.Data) []string {
	c := course{}

	var f []*validation.FieldRules

	if data.Body["courseName"] != nil {
		c.courseName = data.Body["courseName"].(string)
		f = append(f, validation.Field(&c.courseName, validation.Required))
	}

	if data.Body["city"] != nil {
		c.city = data.Body["city"].(string)
		f = append(f, validation.Field(&c.city, validation.Required))
	}

	if data.Body["state"] != nil {
		c.state = data.Body["state"].(string)
		f = append(f, validation.Field(&c.state, validation.Required, validation.Length(2, 2)))
	}

	if data.Body["zip"] != nil {
		c.zip = data.Body["zip"].(string)
		f = append(f, validation.Field(&c.zip, validation.Required))
	}

	err := validation.ValidateStruct(&c)

	return utils.ParseValidationError(err)
}

// Handle the endpoint request
func (pt coursePut) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get id as Int from Url Query
	id := utils.GetUrlQueryInt(r, "id")
	c := db.Course{}.Find(id)
	if c.Id == nil {
		err = server.ThrowError(NotFound)
		return c, err
	}

	// Update the course and populate the data
	if data.Body["courseName"] != nil {
		c.CourseName = utils.ToStringPointer(data.Body["courseName"])
	}

	if data.Body["city"] != nil {
		c.City = utils.ToStringPointer(data.Body["city"])
	}

	if data.Body["state"] != nil {
		c.State = utils.ToStringPointer(data.Body["state"])
	}

	if data.Body["zip"] != nil {
		c.Zip = utils.ToStringPointer(data.Body["zip"])
	}

	// Store the player in the database
	db.Conn.Save(&c)

	return c, err
}
