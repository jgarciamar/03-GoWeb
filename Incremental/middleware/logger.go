package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// This logger is able to log the following:
// - The request method
// - The time of the request
// - The request path
// - The size of bytes of the request
type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Before

		requestTime := time.Now().Format(time.RFC3339)
		method := r.Method
		urlPath := r.URL.Path
		size := r.ContentLength

		fmt.Printf("REQUEST: %s, TIME: %s, PATH: %s, SIZE: %d\n", method, requestTime, urlPath, size)

		next.ServeHTTP(w, r)

		//After

	})
}
