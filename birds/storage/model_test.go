package storage_test

import (
	"github.com/kieron-pivotal/birdpedia/birds"
	"github.com/kieron-pivotal/birdpedia/birds/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storage", func() {
	var (
		birdStorage *storage.Birds
	)

	BeforeEach(func() {
		birdStorage = new(storage.Birds)
		birdStorage.InitList()
	})

	It("can get the bird list", func() {
		Expect(birdStorage.GetList()).To(Equal([]birds.Bird{}))
	})

	It("can add a bird to the list", func() {
		newBird := birds.Bird{
			Species:     "robin",
			Description: "reb bread",
		}
		birdStorage.Add(newBird)
		Expect(birdStorage.GetList()).To(Equal([]birds.Bird{newBird}))
	})

})
