package testing

import (
	"path/filepath"
	"testing"

	"github.com/boltdb/bolt"
)

// CreateTestDB will create a database for the current test
func CreateTestDB(t testing.TB) (db *bolt.DB, teardown func()) {

	tmpdir, td := TempDir(t)
	path := filepath.Join(tmpdir, "test.db")
	database, err := bolt.Open(path, 0600, nil)

	if err != nil {
		t.Fatal(err)
	}

	teardown = func() {
		database.Close()
		td()
	}

	return database, teardown
}
