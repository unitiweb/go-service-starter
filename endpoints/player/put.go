package player

import (
	"github.com/unitiweb/go-service-starter/db"
	"github.com/unitiweb/go-service-starter/server"
	"github.com/unitiweb/go-service-starter/utils"
	"net/http"
)

// The main put struct
// Implements the EndpointInterface
type put struct{}

func PutInit() {
	e := put{}
	server.AddEndpoint(e)
}

// Get endpoint configuration
func (pt put) Config() server.EndpointConfig {
	return server.EndpointConfig{
		Method: http.MethodPut,
		Path:   "/players/{id}",
	}
}

// Validate the endpoints input data
func (pt put) Validate(r *http.Request, data *server.Data) []string {
	return validateInput(data)
}

// Handle the endpoint request
func (pt put) Handle(r *http.Request, data *server.Data) (interface{}, error) {
	var err error

	// Get id as Int from Url Query
	id := utils.GetUrlQueryInt(r, "id")
	p := db.Player{}.Find(id)
	if p.Id == nil {
		err = server.ThrowError(NotFound)
		return p, err
	}

	// Update the player and populate the data
	p.PlayerName = utils.ToStringPointer(data.Body["playerName"])

	// Store the player in the database
	db.Conn.Save(&p)

	return p, err
}
