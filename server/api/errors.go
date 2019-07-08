package api

import (
	"fmt"
	"net/http"
)

func GetHTTPErrorStatusCode(err error) int {
	return http.StatusInternalServerError
}

func MakeErrorHandler(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusCode := GetHTTPErrorStatusCode(err)
		text := fmt.Sprintf("Status-Code: %d, err: %s", statusCode, err.Error())
		http.Error(w, text, statusCode)
	}
}

func NotFound(response http.ResponseWriter) error {
	response.WriteHeader(http.StatusNotFound)
	return nil
}

func Unauthorized(response http.ResponseWriter) error {
	response.WriteHeader(http.StatusUnauthorized)
	return nil
}

func Forbidden(response http.ResponseWriter) error {
	response.WriteHeader(http.StatusForbidden)
	return nil
}
