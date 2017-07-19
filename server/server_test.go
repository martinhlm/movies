package server_test

import (
	"app/server"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	var (
		router *mux.Router
		req    *http.Request
		resp   *httptest.ResponseRecorder
		prefix string
	)

	Describe("Requests controller", func() {

		Context("When client send request to path prefix", func() {
			BeforeEach(func() {
				router = mux.NewRouter()
				prefix = "/v1/movies"
			})

			Context("Basic request", func() {
				It("returns status code of StatusOK (200)", func() {
					server.RegisterHandlers()

					req, _ = http.NewRequest(http.MethodGet, prefix, nil)
					resp = httptest.NewRecorder()
					router.ServeHTTP(resp, req)

					Expect(resp.Code).To(Equal(200))
				})
			})

		})
	})

})
