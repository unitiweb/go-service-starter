package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Create the router
var router = mux.NewRouter().StrictSlash(true)

// A catch all for 404 route not found error
func catchAllNonRoutes() {
	f := func (w http.ResponseWriter, r *http.Request) {
		e := FindError("RouteNotFound")
		HandleError(w, e)
	}
	router.PathPrefix("/").HandlerFunc(f)
}
