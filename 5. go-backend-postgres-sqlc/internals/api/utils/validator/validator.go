package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate
func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		
		// 1. Try to grab our custom "label" tag first
		label := fld.Tag.Get("label")
		if label != "" { return label }
		
		// 2. If there is no label, fall back to the "json" tag
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" { return "" }

		return name
	})
}

// This takes any struct, runs validation, and returns ONLY the first error.
func ValidateStruct(data any) error {

	err := validate.Struct(data)
	if err == nil { return nil }

	// Assert that the error is of type validator.ValidationErrors
		if ve, ok := err.(validator.ValidationErrors); ok {
			// Return only the first formatted error
			return errors.New(getMsgForTag(ve[0])) 
		}
	
	return errors.New("Invalid data provided.")
}