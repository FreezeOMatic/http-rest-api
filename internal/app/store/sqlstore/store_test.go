package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=Vcnpckszvwjj_851 dbname=testdb sslmode=disable"
	}

	os.Exit(m.Run())
}
