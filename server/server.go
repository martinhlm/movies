package server

import (
	"io/ioutil"
	"log"
	"movies/keys"
	"net/http"

	"github.com/gorilla/mux"
)

// PathPrefix for server api
const PathPrefix = "/v1/movies"

// RegisterHandlers create all server api handlers
func RegisterHandlers() http.Handler {
	return NewRouter()
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(PathPrefix, ListMoviesDiscover).Methods(http.MethodGet)
	router.HandleFunc(PathPrefix+"/{movie_id}", GetMovie).Methods(http.MethodGet)

	return router
}

// badRequest is handled by setting the status code in the reply to StatusBadRequest.
type badRequest struct{ error }

// notFound is handled by setting the status code in the reply to StatusNotFound.
type notFound struct{ error }

// errorHandler wraps a function returning an error by handling the error and returning a http.Handler.
// If the error is of the one of the types defined above, it is handled as described for every type.
// If the error is of another type, it is considered as an internal error and its message is logged.
func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

// ListMovies get a list for dicover movies
func ListMoviesDiscover(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	url := keys.PATH_API_TMD + "discover/movie?" + "api_key=" + keys.API_KEY

	for param, value := range params {
		url += "&" + param + "=" + value[0]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {

}
