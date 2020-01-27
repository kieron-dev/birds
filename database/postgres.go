package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ConnectionString string

func NewPostgresConnection(cstr ConnectionString) (*sql.DB, error) {
	return sql.Open("postgres", string(cstr))
}
