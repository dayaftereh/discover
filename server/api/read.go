package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// ReadJSON from the request into the given object
func ReadJSON(request *http.Request, v interface{}) error {
	// get the content-type
	contentType := request.Header.Get("Content-Type")

	// check for json content-type
	if !strings.EqualFold(contentType, "application/json") {
		return errors.Errorf("unable to decoder json from received request, because unexpected content-type [ %s ]", contentType)
	}

	// create the decoder for the body
	decoder := json.NewDecoder(request.Body)

	// decode for given object
	err := decoder.Decode(v)
	if err != nil {
		return errors.Wrapf(err, "unable to decoder json from received request")
	}
	return nil
}
