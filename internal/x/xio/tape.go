package xio

import (
	"os"

	"github.com/pkg/errors"
)

type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	if err := t.File.Truncate(0); err != nil {
		return 0, errors.Wrapf(err, "failed to truncate File %s", t.File.Name())
	}

	if _, err := t.File.Seek(0, 0); err != nil {
		return 0, errors.Wrapf(err, "failed to seek File %s", t.File.Name())
	}

	n, err = t.File.Write(p)
	return n, errors.Wrapf(err, "failed to write to File %s", t.File.Name())
}
