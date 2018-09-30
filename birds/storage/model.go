package storage

import "github.com/kieron-pivotal/birdpedia/birds"

type Birds struct {
	list []birds.Bird
}

func (b *Birds) InitList(initialBirds ...birds.Bird) {
	b.list = []birds.Bird{}
	b.list = append(b.list, initialBirds...)
}

func (b *Birds) GetList() []birds.Bird {
	return b.list
}

func (b *Birds) Add(bird birds.Bird) {
	b.list = append(b.list, bird)
}
