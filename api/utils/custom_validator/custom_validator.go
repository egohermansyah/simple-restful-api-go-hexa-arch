package custom_validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"simple-restful-api-go-hexa-arch/api/utils/response"
	"simple-restful-api-go-hexa-arch/api/utils/strcase"
)

type BodyRequestValidator struct {
	Validator *validator.Validate
}

func (v *BodyRequestValidator) Validate(schema interface{}) error {
	return v.Validator.Struct(schema)
}

func BuildErrorBodyRequestValidator(err error) response.Errors {
	var errors []string
	for _, __ := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("%s : %s", strcase.CamelCaseToSnakeCase(__.Field()), __.Tag()))
	}
	return response.Errors{UserErrors: errors, EngineerErrors: errors}
}
