package xio

import (
	"io"
	"testing"

	"github.com/kiliandbigblue/octoback/internal/x/testhelper"
)

func TestTape_Write(t *testing.T) {
	file, clean := testhelper.CreateTempFile(t, "12345")
	defer clean()

	Tape := &Tape{file}

	Tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
