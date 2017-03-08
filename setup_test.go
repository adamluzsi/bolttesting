package testing_test

import (
	"os"
	"testing"

	. "github.com/LxDB/testing"
	"github.com/boltdb/bolt"
)

func TestSetUp(t *testing.T) {
	db, teardown := SetUp(t)

	path := db.Path()

	t.Log("execute sample bucket creation in the new bolt database")
	err := db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(`test`))
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	teardown()

	t.Log("Check file existence after teardown executed")
	if _, err := os.Stat(path); err == nil {
		t.Fatal("teardown failed to clean up the contents")
	} else {
		t.Log("File is cleaned up successfully!")
	}

}
