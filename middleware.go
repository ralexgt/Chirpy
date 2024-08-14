package main

import (
	"log"
	"net/http"
	"time"
)

// type wrappedWriter struct {
// 	http.ResponseWriter
// 	statusCode int
// }

// func (w *wrappedWriter) wrappedHeader(statuscCode int) {
// 	w.ResponseWriter.WriteHeader(statuscCode)
// 	w.statusCode = statuscCode
// }

// type Middleware func(http.Handler) http.Handler

// func CreateMiddlewareStack(middlewares ...Middleware) Middleware {
// 	return func(next http.Handler) http.Handler {
// 		for i := len(middlewares) - 1; i >= 0; i-- {
// 			middleware := middlewares[i]
// 			next = middleware(next)
// 		}
// 		return next
// 	}
// }

func (cfg *apiConfig) middlewareVisitsAndLog(next http.Handler) http.Handler {
	return  middlewareLogging(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	}))
}

func middlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// wrapped := &wrappedWriter{
		// 	ResponseWriter: w,
		// 	statusCode: http.StatusOK,
		// }

		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}