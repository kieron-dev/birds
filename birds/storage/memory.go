package storage

import "github.com/kieron-pivotal/birdpedia/birds"

type Memory struct {
	list []*birds.Bird
}

func (m *Memory) GetBirds() ([]*birds.Bird, error) {
	return m.list, nil
}

func (m *Memory) CreateBird(bird *birds.Bird) error {
	m.list = append(m.list, bird)
	return nil
}
