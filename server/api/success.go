package api

import "net/http"

func SuccessEmpty(response http.ResponseWriter) error {
	response.WriteHeader(http.StatusNoContent)
	return nil
}

func SuccessOK(response http.ResponseWriter) error {
	response.WriteHeader(http.StatusOK)
	return nil
}
