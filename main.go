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

	// middlewareStack := CreateMiddlewareStack(
	// 	apiCfg.middlewareVisitsInc,
	// 	middlewareLogging,
	// )

	router := http.NewServeMux()
	fileServerHandler := apiCfg.middlewareVisitsAndLog((http.StripPrefix("/app", http.FileServer(http.Dir(filepathToRoot)))))

	// app users routes
	router.Handle("/app/*", fileServerHandler)

	// admin routes
	router.Handle("GET /admin/metrics", middlewareLogging(http.HandlerFunc(apiCfg.handlerVisits)))

	// api routes
	router.Handle("GET /api/healthz", middlewareLogging(http.HandlerFunc(handlerReadiness)))
	router.Handle("GET /api/reset", middlewareLogging(http.HandlerFunc(apiCfg.handlerReset)))
	router.Handle("POST /api/validate_chirp", middlewareLogging(http.HandlerFunc(handlerValidate)))

	server := &http.Server{
		Addr: "127.0.0.1:" + port,
		Handler: router,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathToRoot, port)
	log.Fatal(server.ListenAndServe())
}

