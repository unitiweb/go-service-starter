package utils

import (
	_ "github.com/go-ozzo/ozzo-validation"
	_ "github.com/go-ozzo/ozzo-validation/is"
)

func IsValidString(err []string, v string, r string) {
	//fmt.Println("v", v)
	//validator.
}

func IsValid(err []string, v interface{}, r string) {
	//error := validator.Validate(v)
	//if error != nil {
	//	fmt.Println(error)
	//}
}
