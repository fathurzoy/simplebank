package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// config, err := util.LoadConfig("../..")
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// testDB, err = sql.Open(config.DBDriver, config.DBSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// }

	testQueries = New(conn)

	os.Exit(m.Run())
}