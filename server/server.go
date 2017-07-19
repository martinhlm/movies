package server

import "net/http"

func RegisterHandlers(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
}
