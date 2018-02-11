package testing

import (
	"testing"

	"github.com/boltdb/bolt"
)

// TruncateDB will delete all main bucket from the database
func TruncateDB(t testing.TB, db *bolt.DB) {

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.ForEach(func(key []byte, bucket *bolt.Bucket) error {
			return tx.DeleteBucket(key)
		})
		return err
	})

	if err != nil {
		t.Fatal(err)
	}

}
