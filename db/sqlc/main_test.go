package db

// Tell the Go formatter to keep the github import with _ prefix
// because we're not calling any functions from that package.
import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelvinyrb/simple-bank/util"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// )

var testStore Store

func TestMain(m *testing.M) {
	// var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	// Don't use := synx here because we already declared the variables
	// testDB, err = sql.Open(config.DBDriver, config.DBSource)
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	// testQueries = New(testDB)
	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
