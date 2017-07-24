package server_test

import (
	"movies/server"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	var (
		router *mux.Router
		resp   *httptest.ResponseRecorder
		prefix string
	)

	Describe("Requests controller", func() {

		Describe("When client send request to path prefix", func() {
			BeforeEach(func() {
				router = server.NewRouter()
				resp = httptest.NewRecorder()
			})

			Context("Movie Discover", func() {

				Context("Basic request without params to discover movies", func() {
					It("returns status code of StatusOK (200)", func() {
						prefix = "/v1/movies"
						req, _ := http.NewRequest(http.MethodGet, prefix, nil)
						router.ServeHTTP(resp, req)

						Expect(resp.Code).To(Equal(http.StatusOK))
					})
				})

				Context("Request with params to discover movies", func() {
					It("returns status code of StatusOK", func() {
						prefix = "/v1/movies"
						params := "?language=en-US" +
							"&sort_by=popularity.desc" +
							"&include_adult=false" +
							"&include_video=false" +
							"&page=1"
						req, _ := http.NewRequest(http.MethodGet, prefix+params, nil)
						router.ServeHTTP(resp, req)

						Expect(resp.Code).To(Equal(http.StatusOK))
					})
				})

				Context("Request discover movies post", func() {
					It("then should return not found status", func() {
						prefix = "/v1/movies"
						params := "?lang=ET"
						req, _ := http.NewRequest(http.MethodPost, prefix+params, nil)
						router.ServeHTTP(resp, req)

						Expect(resp.Code).To(Equal(http.StatusNotFound))
					})
				})

			})

			Context("Get single movie", func() {
				Context("Request single movie with ID", func() {
					It("then should return StatusOK", func() {
						id := "324852"
						req, _ := http.NewRequest(http.MethodGet,
							"/v1/movies/"+id, nil)
						router.ServeHTTP(resp, req)

						Expect(resp.Code).To(Equal(http.StatusOK))
					})
				})
			})

		})
	})

})
