package v1

import (
	"os"

	"github.com/pkg/errors"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	if err := t.file.Truncate(0); err != nil {
		return 0, errors.Wrapf(err, "failed to truncate file %s", t.file.Name())
	}

	if _, err := t.file.Seek(0, 0); err != nil {
		return 0, errors.Wrapf(err, "failed to seek file %s", t.file.Name())
	}

	n, err = t.file.Write(p)
	return n, errors.Wrapf(err, "failed to write to file %s", t.file.Name())
}
