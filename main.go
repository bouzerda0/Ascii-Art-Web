package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

// main initializes the HTTP server, registers handlers, and starts the service.
func main() {
	// Register the home page handler for the root path
	http.HandleFunc("/", handlers.HomePageHandler)
	
	// Register the handler for processing ASCII art generation requests
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	// Serve static assets (CSS, images, etc.) from the "assets" directory
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	port := 8080
	fmt.Printf("Server successfully started at http://localhost:%d\n", port)
	
	// Listen and serve on the specified port; log error if startup fails
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error: Could not start the server: %v\n", err)
	}
}