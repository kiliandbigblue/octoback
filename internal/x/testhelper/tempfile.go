package testhelper

import (
	"os"
	"testing"
)

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	_, _ = tmpfile.WriteString(initialData)
	_, _ = tmpfile.Seek(0, 0)

	removeFile := func() {
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
