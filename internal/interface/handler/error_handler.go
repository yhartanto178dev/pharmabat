package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Custom error type untuk domain
type DomainError struct {
	Code    string
	Message string
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Error definitions
var (
	ErrValidation = &DomainError{Code: "VALIDATION_ERROR", Message: "Validation failed"}
	ErrNotFound   = &DomainError{Code: "NOT_FOUND", Message: "Resource not found"}
	ErrDuplicate  = &DomainError{Code: "DUPLICATE_ENTRY", Message: "Duplicate resource"}
)

// Custom HTTP Error Handler
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	response := map[string]interface{}{
		"error":   "Internal Server Error",
		"message": "Something went wrong",
	}

	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		response["error"] = http.StatusText(code)
		response["message"] = e.Message
	case *DomainError:
		code = mapDomainErrorToHTTPStatus(e)
		response["error"] = e.Code
		response["message"] = e.Message
	}

	if !c.Response().Committed {
		c.JSON(code, response)
	}
}

func mapDomainErrorToHTTPStatus(err *DomainError) int {
	switch err.Code {
	case ErrValidation.Code:
		return http.StatusBadRequest
	case ErrNotFound.Code:
		return http.StatusNotFound
	case ErrDuplicate.Code:
		return http.StatusConflict
	default:
		return http.StatusUnprocessableEntity
	}
}
