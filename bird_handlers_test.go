package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BirdHandlers", func() {

	var (
		mockServer *httptest.Server
		birdsURL   string
		testBirds  []Bird
	)

	BeforeEach(func() {
		router := newRouter()
		mockServer = httptest.NewServer(router)
		birdsURL = mockServer.URL + "/bird"
		testBirds = []Bird{
			{Species: "Blackbird", Description: "Black with wings"},
			{Species: "Robin", Description: "Has a red breast"},
		}
		birds = []Bird{}
	})

	Context("bird handler", func() {
		It("returns birds JSON on a GET to /birds", func() {
			birds = testBirds

			resp, err := http.Get(birdsURL)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			contentType := resp.Header.Get("Content-Type")
			Expect(contentType).To(Equal("application/json"))

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			var respBirds []Bird
			err = json.Unmarshal(body, &respBirds)
			Expect(err).NotTo(HaveOccurred())
			Expect(respBirds).To(Equal(testBirds))
		})

		It("redirects to assets on successful POST to /birds", func() {
			resp, err := http.Post(birdsURL, "", nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(resp.Header.Get("Content-Type")).To(Equal("text/html; charset=utf-8"))
		})

		It("will return a POSTed bird in a subsequent GET", func() {
			resp, err := http.PostForm(birdsURL,
				url.Values{"species": {"bluebird"}, "description": {"speedboat"}})
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			resp, err = http.Get(birdsURL)

			expectedBirds := []Bird{
				{Species: "bluebird", Description: "speedboat"},
			}
			expectedJSON, err := json.Marshal(expectedBirds)
			Expect(err).NotTo(HaveOccurred())

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			Expect(body).To(Equal(expectedJSON))
		})
	})
})
