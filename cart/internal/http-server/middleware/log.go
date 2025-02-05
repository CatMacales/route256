package middleware

import (
	"log"
	"net/http"
)

type Logger struct {
	handler http.Handler
}

func NewLogger(handlerToWrap http.Handler) http.Handler {
	return &Logger{handler: handlerToWrap}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("request got: method=%s, path=%s, remote_addr=%s, user_agent=%s", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
	l.handler.ServeHTTP(w, r)
}
