package main

import (
	"fmt"
	"movies/database"
	"movies/server"
	"net/http"
)

func main() {
	fmt.Println("Initializing...")
	// Get the mux router object
	router := server.RegisterHandlers()
	// Create the Server
	ser := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	database.Connect()
	// Running the Server
	ser.ListenAndServe()
}
