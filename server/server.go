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
	//params := &models.Movie{}

	params := r.URL.Query()
	for i, param := range params {
		fmt.Println(i)
		fmt.Println(param)
	}
	//fmt.Println(params.Get("language"))

	//foo, _ := json.Marshal(params)
	//fmt.Println(string(foo))

	url := keys.PATH_API_TMD + "discover/movie?" +
		"api_key=912f66ded3f67606bc9ca4503e68c8c1&language=en-US&sort_by=popularity.desc&include_adult=false&include_video=false&page=1"

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
