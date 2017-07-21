package server

import (
	"fmt"
	"io/ioutil"
	"movies/keys"
	"net/http"

	"github.com/gorilla/mux"
)

// PathPrefix for server api
const PathPrefix = "/v1/movies"

// RegisterHandlers create all server api handlers
func RegisterHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix, ListMoviesDiscover).Methods("GET")
	http.Handle(PathPrefix, r)

	return r
}

// badRequest is handled by setting the status code in the reply to StatusBadRequest.
type badRequest struct{ error }

// ListMovies get a list for dicover movies
func ListMoviesDiscover(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	url := keys.PATH_API_TMD + "discover/movie?" + "api_key=" + keys.API_KEY

	for param, value := range params {
		url += "&" + param + "=" + value[0]
		//fmt.Println(param)
		//fmt.Println(value[0])
	}
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
