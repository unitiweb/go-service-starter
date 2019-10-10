package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetPostBody(r *http.Request, b interface{}) error {
	// Get the body as a string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// Unmarshal the body string into the pb struct
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Println("Json Error:", err)
		return err
	}

	return nil
}
