package bolttesting

import (
	"io/ioutil"
	"os"
	"testing"
)

// TempDir return a tempdirectory path and a teardown to clean up that tmpdirectory after the test
func TempDir(t testing.TB) (tmpdir string, teardown func()) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), generateUUID())

	if err != nil {
		t.Fatal(err)
	}

	teardown = func() {
		os.RemoveAll(tmpdir)
	}

	return tmpdir, teardown
}
