package server

import (
	"fmt"
	"go-project-layout/internal/middleware"
	"go-project-layout/internal/routes"
	"log"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	routes.Health(mux)
}

func SetupMiddleware(handler http.Handler) http.Handler {
	panicHandler := middleware.RecoverPanic(handler)
	corsHandler := middleware.CorsMiddleware(panicHandler)
	loggingHandler := middleware.LoggerMiddleware(corsHandler)
	return loggingHandler
}

func StartServer(handler http.Handler, port string) {
	fmt.Printf("Server is starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
