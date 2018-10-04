package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/kieron-pivotal/birdpedia/birds"
	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/birds/storage/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BirdHandlers", func() {

	var (
		store       *fakes.FakeStore
		birdHandler handlers.Handler
		testBirds   []*birds.Bird
	)

	BeforeEach(func() {
		store = new(fakes.FakeStore)
		birdHandler = handlers.NewHandler(store)
		testBirds = []*birds.Bird{
			{Species: "Blackbird", Description: "Black with wings"},
			{Species: "Robin", Description: "Has a red breast"},
		}
	})

	Context("bird handler", func() {
		It("returns birds JSON on a GET to /bird", func() {
			store.GetBirdsReturns(testBirds, nil)

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

		It("passes POST vars to create bird", func() {
			data := url.Values{}
			data.Add("species", "foo")
			data.Add("description", "bar")
			req, err := http.NewRequest("POST", "", strings.NewReader(data.Encode()))
			Expect(err).NotTo(HaveOccurred())

			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			recorder := httptest.NewRecorder()
			hf := http.HandlerFunc(birdHandler.CreateBird)
			hf.ServeHTTP(recorder, req)

			Expect(store.CreateBirdCallCount()).To(Equal(1))
			bird := store.CreateBirdArgsForCall(0)
			Expect(bird.Species).To(Equal("foo"))
			Expect(bird.Description).To(Equal("bar"))

			Expect(recorder.Code).To(Equal(http.StatusFound))
		})

	})
})
