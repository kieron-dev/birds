package db

import (
	"database/sql"

	"github.com/kieron-pivotal/birdpedia/birds"
	"github.com/kieron-pivotal/birdpedia/birds/storage"
)

type Store struct {
	db *sql.DB
}

func NewDBStore(db *sql.DB) storage.Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateBird(bird *birds.Bird) error {
	_, err := s.db.Query("INSERT INTO birds(species, description) VALUES ($1, $2)",
		bird.Species, bird.Description)

	return err
}

func (s *Store) GetBirds() ([]*birds.Bird, error) {
	rows, err := s.db.Query("SELECT species, description FROM birds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*birds.Bird{}
	for rows.Next() {
		bird := new(birds.Bird)
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		list = append(list, bird)
	}
	return list, nil
}
