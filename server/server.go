package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

const PathPrefix = "/v1/movies/"

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix, ListMovies).Methods("GET")
	http.Handle(PathPrefix, r)
}

func ListMovies(w http.ResponseWriter, r *http.Request) {

}
