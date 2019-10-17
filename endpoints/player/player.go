package player

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/unitiweb/go-service-starter/server"
	"github.com/unitiweb/go-service-starter/utils"
)

type player struct {
	playerName string
}

const NotFound = "PlayerNotFound"

func Init() {
	// Add the player not found error
	server.AddError(NotFound, 404, "The requested player could not be found")

	// Initialize Endpoints
	GetInit()
	GetAllInit()
	PostInit()
	PutInit()
	DeleteInit()
}

func validateInput(data *server.Data) []string {
	p := player{
		playerName: data.Body["playerName"].(string),
	}

	err := validation.ValidateStruct(&p,
		validation.Field(&p.playerName, validation.Required, validation.Length(5, 256)),
	)

	return utils.ParseValidationError(err)
}
