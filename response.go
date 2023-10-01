package httprw

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Response(
	w http.ResponseWriter,
	responseBody any,
	responseHeaders map[string]string,
	statusCode int,
) {
	buffer, err := makeBuffer(responseBody)
	if err != nil {
		return
	}
	response(w, buffer, responseHeaders, statusCode)
}

/*
It is expected that, value for "data" parameter will be a type of struct.
Where the struct members may contain `json` tag.
*/
func makeBuffer(data any) (io.Reader, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(data)
	return buffer, err
}

func response(
	writer http.ResponseWriter,
	responseBodyBuffer io.Reader,
	responseHeaders map[string]string,
	statusCode int,
) error {
	for key, value := range responseHeaders {
		writer.Header().Set(key, value)
	}
	writer.WriteHeader(statusCode)
	_, err := io.Copy(writer, responseBodyBuffer)
	return err
}
