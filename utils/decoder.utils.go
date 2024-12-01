package utils

import (
	"dating-sim/res"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func DecodeJson(v any, w http.ResponseWriter, r io.Reader, isStrict bool) error {
	decoder := json.NewDecoder(r)
	if isStrict {
		decoder.DisallowUnknownFields()
	}
	err := decoder.Decode(&v)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			res.ResErrJson(w, http.StatusBadRequest, errors.New(msg))
			return errors.New(msg)
		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "request body contains badly-formed JSON"
			res.ResErrJson(w, http.StatusBadRequest, errors.New(msg))
			return errors.New(msg)
		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			res.ResErrJson(w, http.StatusBadRequest, errors.New(msg))
			return errors.New(msg)
		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			res.ResErrJson(w, http.StatusBadRequest, errors.New(msg))
			return errors.New(msg)
		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg := "request body must not be empty"
			res.ResErrJson(w, http.StatusBadRequest, errors.New(msg))
			return errors.New(msg)
		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		case err.Error() == "http: request body too large":
			msg := "request body must not be larger than 1MB"
			res.ResErrJson(w, http.StatusRequestEntityTooLarge, errors.New(msg))
			return errors.New(msg)
		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Print(err.Error())
			res.ResErrJson(w, http.StatusInternalServerError, errors.New("internal server error"))
			return errors.New("internal server errror")
		}
	}
	return nil
}
