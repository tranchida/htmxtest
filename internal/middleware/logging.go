package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	bytes      int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(b)
	lrw.bytes += n
	return n, err
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := r.RemoteAddr
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)
		t := time.Now().Format("02/Jan/2006:15:04:05 -0700")
		log.Printf("%s - - [%s] \"%s %s %s\" %d %d %v",
			ip,
			t,
			r.Method,
			r.URL.RequestURI(),
			r.Proto,
			lrw.statusCode,
			lrw.bytes,
			duration,
		)
	})
}
