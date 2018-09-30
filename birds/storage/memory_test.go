package storage_test

import (
	"github.com/kieron-pivotal/birdpedia/birds"
	"github.com/kieron-pivotal/birdpedia/birds/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storage", func() {
	var (
		birdStorage *storage.Memory
	)

	BeforeEach(func() {
		birdStorage = new(storage.Memory)
	})

	It("can get the bird list", func() {
		list, err := birdStorage.GetBirds()
		Expect(err).NotTo(HaveOccurred())
		Expect(list).To(BeEmpty())
	})

	It("can add a bird to the list", func() {
		newBird := &birds.Bird{
			Species:     "robin",
			Description: "reb bread",
		}
		birdStorage.CreateBird(newBird)
		Expect(birdStorage.GetBirds()).To(Equal([]*birds.Bird{newBird}))
	})

})
