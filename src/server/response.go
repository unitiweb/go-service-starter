package server

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response structure that will wrap the response structure
type Response struct {
	Data interface{} `json:"data"`
}

func ToJson(w http.ResponseWriter, d interface{}) {
	b, err := json.Marshal(
		Response{
			Data: d,
		},
	)
	if err != nil {
		io.WriteString(w, "Error could not be converted to json")
	}

	io.WriteString(w, string(b))
}
