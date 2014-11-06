package json

import (
	j "encoding/json"
	"net/http"
	"strconv"
)

type Wr struct {
	W http.ResponseWriter
	R *http.Request
}

type JsonWriter struct {
	Rw       Wr
	Status   int
	Error    error
	Body     interface{}
	BodyOnly bool
}

func WriteJson(jw JsonWriter) {
	// Set content-type to JSON
	jw.Rw.W.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Figure out the error
	var errs string
	if jw.Error != nil {
		errs = "\"" + jw.Error.Error() + "\""
	} else {
		// We set it to "null" instead of "nil" because
		// "nil" is not proper JSON.
		errs = "null"
	}

	// Figure out body JSON
	body, err := j.Marshal(jw.Body)
	if err != nil {
		// If we have an error, report it.
		jw.Status = 500
		jw.Error = err
		body = []byte("null")
	}

	// Check if we only want to return the body content
	v := jw.Rw.R.URL.Query()
	only, err := strconv.ParseBool(v.Get("body"))
	if err != nil {
		only = false
	}

	// Only show body of JSON if that's what we want.
	var json string
	if only {
		json = string(body)
	} else {
		// Put together full JSON string
		json = "{\"Status\": " + strconv.Itoa(jw.Status) + ", \"Error\": " + errs + ", \"Body\": " + string(body) + "}"
	}

	// Write JSON
	jw.Rw.W.Write([]byte(json))
}
