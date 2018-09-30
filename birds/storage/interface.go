package storage

import "github.com/kieron-pivotal/birdpedia/birds"

type Store interface {
	CreateBird(bird *birds.Bird) error
	GetBirds() ([]*birds.Bird, error)
}
