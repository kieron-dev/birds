package birds_db_test

import (
	"github.com/kieron-pivotal/birdpedia/birds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BirdsDb", func() {
	BeforeEach(func() {
		conn.Query("DELETE FROM birds")
		conn.Query("INSERT INTO birds (species, description) VALUES ($1, $2)", "species", "desc")
	})

	AfterEach(func() {
		conn.Query("DELETE FROM birds")
	})

	It("can get birds", func() {
		list, err := store.GetBirds()
		Expect(err).NotTo(HaveOccurred())
		Expect(list).To(HaveLen(1))

		conn.Query("INSERT INTO birds (species, description) VALUES ($1, $2)", "bird", "type")

		list, err = store.GetBirds()
		Expect(err).NotTo(HaveOccurred())
		Expect(list).To(HaveLen(2))
		Expect(list[1]).To(Equal(&birds.Bird{Species: "bird", Description: "type"}))
	})

	It("can create a bird", func() {
		err := store.CreateBird(&birds.Bird{
			Species:     "Frigate bird",
			Description: "Pteradactyl",
		})
		Expect(err).NotTo(HaveOccurred())

		res, err := conn.Query("SELECT COUNT(*) FROM birds")
		Expect(err).NotTo(HaveOccurred())

		var count int
		for res.Next() {
			err := res.Scan(&count)
			Expect(err).NotTo(HaveOccurred())
			Expect(count).To(Equal(2))
			res.Close()
		}
	})

})
