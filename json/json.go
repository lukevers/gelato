package json

import (
	j "encoding/json"
	"net/http"
	"strconv"
)

type JsonWriter struct {
	Status int
	Rw     http.ResponseWriter
	Error  error
	Body   interface{}
}

func WriteJson(jw JsonWriter) {
	// Set content-type to json
	jw.Rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Figure out the error
	var errs string
	if jw.Error != nil {
		errs = "\"" + jw.Error.Error() + "\""
	} else {
		// We set it to "null" instead of "nil" because
		// "nil" is not proper JSON.
		errs = "null"
	}

	// Figure out body json
	body, err := j.Marshal(jw.Body)
	if err != nil {
		// If we have an error, report it.
		jw.Status = 500
		jw.Error = err
		body = []byte("null")
	}

	// Put together json string
	json := "{\"Status\": " + strconv.Itoa(jw.Status) + ", \"Error\": " + errs + ", \"Body\": " + string(body) + "}"

	// Write json
	jw.Rw.Write([]byte(json))
}
