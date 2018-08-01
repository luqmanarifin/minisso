package service

import (
	"encoding/json"
	"net/http"

	"github.com/bukalapak/packen/response"
)

// Decode body to a struct
func Decode(r *http.Request, object interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	return decoder.Decode(&object)
}

// handle response both success or failed
func HandleResponse(w http.ResponseWriter, response interface{}, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	switch status {
	case 200:
		WriteSuccess(w, response, status)
		break
	case 201:
		WriteSuccess(w, response, status)
		break
	default:
		WriteError(w, err, status)
		break
	}

}

// build error response and write it
func WriteError(w http.ResponseWriter, err string, status int) {
	errCust := response.CustomError{
		Message:  err,
		HTTPCode: status,
	}
	errs := []error{errCust}
	res := response.BuildError(errs)
	response.Write(w, res, status)
}

// build success response and write it
func WriteSuccess(w http.ResponseWriter, data interface{}, status int) {
	res := response.BuildSuccess(data, response.MetaInfo{HTTPStatus: status})
	response.Write(w, res, status)
}
