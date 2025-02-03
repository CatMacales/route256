package http_server

import (
	"bytes"
	"net/http"
)

func GetErrorResponse(w http.ResponseWriter, handlerName string, err error, statusCode int) {
	w.WriteHeader(statusCode)
	buf := bytes.NewBufferString(handlerName)
	buf.WriteString(": ")
	buf.WriteString(err.Error())
	buf.WriteString("\n")
	_, _ = w.Write(buf.Bytes())
}
