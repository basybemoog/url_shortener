package response

import (
	"fmt"
	"github.com/go-playground/validator"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "ok"
	StatusError = "error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func ERROR(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMessages []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMessages = append(errMessages, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errMessages = append(errMessages, fmt.Sprintf("field %s is not a valid URL", err.Field()))
		default:
			errMessages = append(errMessages, fmt.Sprintf("field %s is not valid"), err.Field())
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMessages, ", "),
	}

}
