package utils

import "gopkg.in/go-playground/validator.v8"

var Validator *validator.Validate

func init() {
	config := &validator.Config{TagName: "validate"}
	Validator = validator.New(config)
}
