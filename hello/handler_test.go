package hello_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/kieron-pivotal/birdpedia/hello"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {

	Context("routing", func() {

		It("responds to GET on /hello", func() {
			req, err := http.NewRequest("GET", "", nil)
			Expect(err).NotTo(HaveOccurred())

			recorder := httptest.NewRecorder()
			hf := http.HandlerFunc(hello.Handler)
			hf.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(string(recorder.Body.Bytes())).To(Equal("Hello World!"))
		})

	})
})
