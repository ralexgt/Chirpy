package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) middlewareVisitsInc(next http.Handler) http.Handler {
	return  middlewareLog(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	}))
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}