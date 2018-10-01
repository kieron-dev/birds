package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/kieron-pivotal/birdpedia/birds"
	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/birds/storage/memory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BirdHandlers", func() {

	var (
		birdStorage *memory.Store
		birdHandler handlers.Handler
		testBirds   []*birds.Bird
	)

	BeforeEach(func() {
		birdStorage = new(memory.Store)
		birdHandler = handlers.NewHandler(birdStorage)
		testBirds = []*birds.Bird{
			{Species: "Blackbird", Description: "Black with wings"},
			{Species: "Robin", Description: "Has a red breast"},
		}
	})

	Context("bird handler", func() {
		It("returns birds JSON on a GET to /bird", func() {
			for _, b := range testBirds {
				birdStorage.CreateBird(b)
			}

			req, err := http.NewRequest("GET", "", nil)
			Expect(err).NotTo(HaveOccurred())

			recorder := httptest.NewRecorder()
			hf := http.HandlerFunc(birdHandler.GetBirds)
			hf.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))

			contentType := recorder.Header().Get("Content-Type")
			Expect(contentType).To(Equal("application/json"))

			var respBirds []*birds.Bird
			err = json.Unmarshal(recorder.Body.Bytes(), &respBirds)
			Expect(err).NotTo(HaveOccurred())
			Expect(respBirds).To(Equal(testBirds))
		})

		It("redirects to assets on successful POST to /bird", func() {
			data := url.Values{}
			data.Add("species", "foo")
			data.Add("description", "bar")
			req, err := http.NewRequest("POST", "", strings.NewReader(data.Encode()))
			Expect(err).NotTo(HaveOccurred())

			recorder := httptest.NewRecorder()
			hf := http.HandlerFunc(birdHandler.CreateBird)
			hf.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusFound))
		})

		It("will return a POSTed bird in a subsequent GET", func() {
			data := url.Values{}
			data.Add("species", "bluebird")
			data.Add("description", "speedboat")

			req, err := http.NewRequest("POST", "", strings.NewReader(data.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			Expect(err).NotTo(HaveOccurred())

			recorder := httptest.NewRecorder()
			hf := http.HandlerFunc(birdHandler.CreateBird)
			hf.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusFound))

			req, err = http.NewRequest("GET", "", nil)
			Expect(err).NotTo(HaveOccurred())

			recorder = httptest.NewRecorder()
			hf = http.HandlerFunc(birdHandler.GetBirds)
			hf.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))

			expectedBirds := []birds.Bird{
				{Species: "bluebird", Description: "speedboat"},
			}
			expectedJSON, err := json.Marshal(expectedBirds)
			Expect(err).NotTo(HaveOccurred())

			Expect(recorder.Body.Bytes()).To(Equal(expectedJSON))
		})
	})
})
