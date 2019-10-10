package utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUrlQueryInt(r *http.Request, key string) int {
	id, _ := strconv.Atoi(mux.Vars(r)[key])
	return id
}
