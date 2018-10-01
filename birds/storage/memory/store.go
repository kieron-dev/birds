package memory

import "github.com/kieron-pivotal/birdpedia/birds"

type Store struct {
	list []*birds.Bird
}

func (m *Store) GetBirds() ([]*birds.Bird, error) {
	return m.list, nil
}

func (m *Store) CreateBird(bird *birds.Bird) error {
	m.list = append(m.list, bird)
	return nil
}
