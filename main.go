package main

import (
	"movies/server"
	"net/http"
)

func main() {
	//url := "https://api.themoviedb.org/3/genre/movie/list?language=en-US&api_key=912f66ded3f67606bc9ca4503e68c8c1"

	//payload := strings.NewReader("{}")

	//req, _ := http.NewRequest("GET", url, payload)

	//res, _ := http.DefaultClient.Do(req)

	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

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
