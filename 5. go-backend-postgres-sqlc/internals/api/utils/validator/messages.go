package validator

import "github.com/go-playground/validator/v10"

func getMsgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Please provide a valid email address"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters long"
	case "max":
		return fe.Field() + " cannot exceed " + fe.Param() + " characters"
	case "gte":
		return fe.Field() + " must be greater than or equal to " + fe.Param()
	case "lte":
		return fe.Field() + " must be less than or equal to " + fe.Param()
	case "gt":
		return fe.Field() + " must be greater than " + fe.Param()
	case "lt":
		return fe.Field() + " must be less than " + fe.Param()
	}
	
	return fe.Field() + " is invalid" // Default fallback
}