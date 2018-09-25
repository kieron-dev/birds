package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("web server", func() {

	Context("hello handler", func() {
		It("says hello", func() {
			req, err := http.NewRequest("GET", "", nil)
			Expect(err).NotTo(HaveOccurred())

			recorder := httptest.NewRecorder()

			hf := http.HandlerFunc(handler)
			hf.ServeHTTP(recorder, req)

			status := recorder.Code
			Expect(status).To(Equal(http.StatusOK))

			body := recorder.Body.String()
			Expect(body).To(Equal("Hello World!"))
		})
	})

	Context("routing", func() {
		var (
			router     *mux.Router
			mockServer *httptest.Server
		)

		BeforeEach(func() {
			router = newRouter()
			mockServer = httptest.NewServer(router)
		})

		It("responds to GET on /hello", func() {
			resp, err := http.Get(mockServer.URL + "/hello")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(b)).To(Equal("Hello World!"))
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

		It("serves static files on /assets/", func() {
			resp, err := http.Get(mockServer.URL + "/assets/")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			contentType := resp.Header.Get("Content-Type")
			Expect(contentType).To(Equal("text/html; charset=utf-8"))
		})
	})

})
