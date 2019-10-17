package course

import (
	"github.com/unitiweb/go-service-starter/server"
)

type course struct {
	courseName string
	city string
	state string
	zip string
}

const NotFound = "CourseNotFound"

func Init() {
	server.AddError(NotFound, 404, "The requested course could not be found")

	// Initialize Endpoints
	GetInit()
	GetAllInit()
	PostInit()
	PutInit()
	DeleteInit()
}
