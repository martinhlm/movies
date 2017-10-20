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
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Connecting db...")
	database.Connect()

	// Running the Server
	fmt.Println("Rock and Roll...")
	ser.ListenAndServe()
}
