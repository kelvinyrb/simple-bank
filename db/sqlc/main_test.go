package db

// Tell the Go formatter to keep the github import with _ prefix
// because we're not calling any functions from that package.
import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/kelvinyrb/simple-bank/util"
	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	// Don't use := synx here because we already declared the variables
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
