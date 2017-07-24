package main

import (
	"movies/server"
	"net/http"
)

func main() {
	// Get the mux router object
	router := server.RegisterHandlers()
	// Create the Server
	ser := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	// Running the Server
	ser.ListenAndServe()
}
