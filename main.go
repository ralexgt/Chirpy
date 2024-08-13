package main

import (
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr: "localhost:8080",
		Handler: serverMux,
	}
	server.ListenAndServe()
}