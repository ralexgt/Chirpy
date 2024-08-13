package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	const filepathToRoot = "."
	const port = "8080"

	apiCfg := apiConfig{fileserverHits: 0}
	serverMux := http.NewServeMux()
	serverMux.Handle("/app/*", apiCfg.middlewareVisitsInc((http.StripPrefix("/app", http.FileServer(http.Dir(filepathToRoot))))))
	serverMux.Handle("GET /healthz", middlewareLog(http.HandlerFunc(handlerReadiness)))
	serverMux.Handle("GET /metrics", middlewareLog(http.HandlerFunc(apiCfg.handlerVisits)))
	serverMux.Handle("GET /reset", middlewareLog(http.HandlerFunc(apiCfg.handlerReset)))

	server := &http.Server{
		Addr: "127.0.0.1:" + port,
		Handler: serverMux,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathToRoot, port)
	log.Fatal(server.ListenAndServe())
}

