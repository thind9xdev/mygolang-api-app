package main

import (
	"log"
	"mygolangapp/routes"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	log.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
