package utils

import (
	"strings"
)

// Parse the errors and put into a slice to be returned
func ParseValidationError(verr error) []string {
	var err []string

	if verr != nil {
		errStr := verr.Error()
		errsArray := strings.Split(errStr, ";")
		for _, err1 := range errsArray {
			err = append(err, strings.Trim(err1, " "))
		}
	}

	return err
}
