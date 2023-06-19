package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/davidroossien/mysimplebank/util"
	_ "github.com/lib/pq"
)

// global variables that are available to other tests
var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
