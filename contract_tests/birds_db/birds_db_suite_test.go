package birds_db_test

import (
	"database/sql"
	"log"

	"github.com/kieron-pivotal/birdpedia/birds/storage"
	"github.com/kieron-pivotal/birdpedia/birds/storage/db"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBirdsDb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BirdsDb Suite")
}

var (
	store storage.Store
	conn  *sql.DB
)

const connString = "host=/var/run/postgresql dbname=bird_encyclopedia sslmode=disable"

var _ = BeforeSuite(func() {
	var err error
	conn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	store = db.NewDBStore(conn)
})

var _ = AfterSuite(func() {
	conn.Close()
})
