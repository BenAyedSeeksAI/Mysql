package apires

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	ERR_MSG_SOMETHING_WENT_WRONG = "Something went wrong. Please try again."
)

type APIError struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

type APIResponse struct {
	Error      []APIError `json:"error"`
	Data       any        `json:"data"`
	StatusCode int        `json:"statusCode"`
}

func NewError(message, description string) *APIError {
	return &APIError{
		Message:     message,
		Description: description,
	}
}

func NewDefaultError() *APIError {
	return &APIError{
		Message:     ERR_MSG_SOMETHING_WENT_WRONG,
		Description: "",
	}
}

func Errors(c *gin.Context, statusCode int, err []APIError) {
	c.JSON(statusCode, &APIResponse{
		Error:      err,
		Data:       []interface{}{},
		StatusCode: statusCode,
	})
}

func Error(c *gin.Context, statusCode int, message, description string) {
	c.JSON(statusCode, &APIResponse{
		Error:      []APIError{*NewError(message, description)},
		Data:       []interface{}{},
		StatusCode: statusCode,
	})
}

func DefaultError(c *gin.Context) {

	c.JSON(http.StatusInternalServerError, &APIResponse{
		Error:      []APIError{*NewDefaultError()},
		Data:       []interface{}{},
		StatusCode: http.StatusInternalServerError,
	})

}

func Success(c *gin.Context, statusCode int, data any) {

	var nonNilData any
	value := reflect.ValueOf(data)
	if data == nil || (value.Kind() == reflect.Pointer && value.IsNil()) {
		nonNilData = map[string]any{}
	} else {
		nonNilData = data
	}

	c.JSON(statusCode, &APIResponse{
		Error:      []APIError{},
		Data:       nonNilData,
		StatusCode: statusCode,
	})

}
