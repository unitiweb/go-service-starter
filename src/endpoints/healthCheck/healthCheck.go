package healthCheck

import (
	"github.com/unitiweb/go-microservice/src/server"
	"net/http"
)

type success struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Endpoint() server.Endpoint {
	return server.Endpoint{
		Method:   http.MethodPost,
		Path:     "/health-check",
		Validate: validate,
		Handle:   handle,
		Config:   config,
	}
}

func config() {
	server.AddError("ItemNotFound", 404, "The requested item could not be found")
	//server.AddMiddleware(func (w http.ResponseWriter, r *http.Request, next http.Handler) {
	//	fmt.Println("Just TEsting middleware")
	//	next.ServeHTTP(w, r)
	//})
}

func validate(r *http.Request) []string {
	var err []string

	//err = append(err, "Input had errors")

	return err
}

func handle(r *http.Request) (interface{}, error) {
	var res success
	var err error

	show := false

	if show == true {
		err = server.ThrowError("ItemNotFound")
		return res, err
	}

	res = success{
		"Success",
		"Service is running and healthy",
	}

	return res, err
}
