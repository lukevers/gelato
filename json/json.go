package json

import (
	"net/http"
	"strconv"
)

type JsonWriter struct {
	Status int
	Rw     http.ResponseWriter
	Error  error
	Body   string //interface{}
}

func WriteJson(jw JsonWriter) {
	// Set content-type to json
	jw.Rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Figure out the error
	var err string
	if jw.Error != nil {
		err = "\"" + jw.Error.Error() + "\""
	} else {
		err = "nil"
	}

	// Figure out body json
	body := jw.Body

	// Create json
	json := "{\"status\": "+strconv.Itoa(jw.Status)+", \"error\": "+err+", \"body\": "+body+"}"

	// Write json
	jw.Rw.Write([]byte(json))
}
