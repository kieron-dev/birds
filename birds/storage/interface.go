package storage

import "github.com/kieron-pivotal/birdpedia/birds"

//go:generate counterfeiter -o fakes/store.go . Store

type Store interface {
	CreateBird(bird *birds.Bird) error
	GetBirds() ([]*birds.Bird, error)
}
