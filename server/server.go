package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PathPrefix for server api
const PathPrefix = "/v1/movies/"

// RegisterHandlers create all server api handlers
func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix, ListMovies).Methods("GET")
	http.Handle(PathPrefix, r)
}

// ListMovies get a list for dicover movies
func ListMovies(w http.ResponseWriter, r *http.Request) {

}
