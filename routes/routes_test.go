package routes_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/routes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("web server", func() {

	Context("routing", func() {
		var (
			mockServer *httptest.Server
		)

		BeforeEach(func() {
			router := routes.NewRouter(handlers.Handler{})
			mockServer = httptest.NewServer(router)
		})

		It("serves static files on /assets/", func() {
			resp, err := http.Get(mockServer.URL + "/assets/")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			contentType := resp.Header.Get("Content-Type")
			Expect(contentType).To(Equal("text/html; charset=utf-8"))
		})

		It("returns 405 for POSTs to /hello", func() {
			resp, err := http.Post(mockServer.URL+"/hello", "", nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusMethodNotAllowed))

			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			Expect(string(b)).To(Equal(""))
		})
	})

})
