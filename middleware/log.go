package middleware

import (
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type Middleware struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *Middleware {
	return &Middleware{w, http.StatusOK}
}

func (m *Middleware) WriteHeader(code int) {
	m.statusCode = code
	m.ResponseWriter.WriteHeader(code)
}

func (m Middleware) LoggingHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		h.ServeHTTP(w, r)
		t2 := time.Now()

		lrw := NewLoggingResponseWriter(w)
		statusCode := lrw.statusCode
		log.Printf("[%s]  %v %v %v %v", r.Method, r.URL.String(),clientIP(r), statusCode,t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}


func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
