package v1

import (
	"encoding/json"
	"os"
	"slices"

	"github.com/bufbuild/protovalidate-go"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	"github.com/pkg/errors"
)

type FileSystemGroceryStore struct {
	database *json.Encoder
	data     GroceryLists
	v        *protovalidate.Validator
}

func NewFileSystemGroceryStore(file *os.File) (*FileSystemGroceryStore, error) {
	v, _ := protovalidate.New()

	if err := initializeFileSystemStoreFile(file); err != nil {
		return nil, errors.Wrapf(err, "failed to initialize file %s", file.Name())
	}

	data, err := NewGroceryLists(file)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load grocery lists from file %s", file.Name())
	}

	return &FileSystemGroceryStore{
		database: json.NewEncoder(&tape{file}),
		data:     data,
		v:        v,
	}, nil
}

// GroceryLists returns all grocery lists.
//
//nolint:unparam //error is always nil.
func (f *FileSystemGroceryStore) GroceryLists() ([]*models.GroceryList, error) {
	return f.data, nil
}

func (f *FileSystemGroceryStore) GroceryList(id string) (*models.GroceryList, error) {
	gl := f.data.Find(id)
	if gl == nil {
		return nil, ErrNoSuchEntity
	}
	return gl, nil
}

func (f *FileSystemGroceryStore) SetGroceryList(r *models.GroceryList) error {
	if err := f.v.Validate(r); err != nil {
		return &StoreValidationError{Err: err}
	}

	gl := f.data.Find(r.GetId())
	switch {
	case gl != nil:
		gl.Name = r.GetName()
		gl.Items = r.GetItems()
	case gl == nil:
		f.data = append(f.data, r)
	}

	if err := f.flush(); err != nil {
		return &StoreInternalError{Err: err}
	}

	return nil
}

func (f *FileSystemGroceryStore) DeleteGroceryList(id string) error {
	var found bool
	for i, gl := range f.data {
		if gl.Id == id {
			f.data[i] = f.data[len(f.data)-1]
			found = true
			break
		}
	}
	if !found {
		return ErrNoSuchEntity
	}
	f.data = f.data[:len(f.data)-1]

	if err := f.flush(); err != nil {
		return &StoreInternalError{Err: err}
	}

	return nil
}

func (f *FileSystemGroceryStore) flush() error {
	slices.SortStableFunc(f.data, func(lhs, rhs *models.GroceryList) int {
		switch {
		case lhs.GetId() < rhs.GetId():
			return -1
		case lhs.GetId() > rhs.GetId():
			return 1
		default:
			return 0
		}
	})

	if err := f.database.Encode(f.data); err != nil {
		return &StoreInternalError{Err: err}
	}
	return nil
}

func initializeFileSystemStoreFile(file *os.File) error {
	if _, err := file.Seek(0, 0); err != nil {
		return errors.Wrapf(err, "failed to seek file %s", file.Name())
	}

	info, err := file.Stat()
	if err != nil {
		return errors.Wrapf(err, "failed to stat file %s", file.Name())
	}

	if info.Size() == 0 {
		if _, err := file.WriteString("[]"); err != nil {
			return errors.Wrapf(err, "failed to write to file %s", file.Name())
		}

		if _, err := file.Seek(0, 0); err != nil {
			return errors.Wrapf(err, "failed to seek file %s", file.Name())
		}
	}

	return nil
}
