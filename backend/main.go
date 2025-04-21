package main

import (
	"backend/api"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Configurable constants, can be loaded from environment variables
var (
	port      = ":8080"
	staticDir = "static"
)

func init() {
	api.GenOrLoadKey()

	// Get environment variables, using default values if not set
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = ":" + envPort
	}

	envStaticDir := os.Getenv("STATIC_DIR")
	if envStaticDir != "" {
		staticDir = envStaticDir
	}

	// this will be used later on in docker build, in dev it's not used
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		if err := os.Mkdir(staticDir, 0755); err != nil {
			log.Fatalf("Failed to create static directory: %v", err)
		}
	}
}

func main() {
	// File server
	fs := http.FileServer(http.Dir(staticDir))

	// Handlers
	http.Handle("/", http.StripPrefix("/", fs)) // Serves files from the static directory
	http.HandleFunc("/api", api.Router)

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
