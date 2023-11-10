package v1

import (
	"encoding/json"
	"io"

	"github.com/bufbuild/protovalidate-go"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
)

type FileSystemGroceryStore struct {
	database io.ReadWriteSeeker
	v        *protovalidate.Validator
}

func NewFileSystemGroceryStore(database io.ReadWriteSeeker) *FileSystemGroceryStore {
	v, _ := protovalidate.New()
	return &FileSystemGroceryStore{
		database: database,
		v:        v,
	}
}

func (f *FileSystemGroceryStore) GroceryLists() ([]*models.GroceryList, error) {
	_, _ = f.database.Seek(0, 0)
	return NewGroceryLists(f.database)
}

func (f *FileSystemGroceryStore) GroceryList(id string) (*models.GroceryList, error) {
	gls, err := f.GroceryLists()
	if err != nil {
		return nil, err
	}

	for _, gl := range gls {
		if gl.Id == id {
			return gl, nil
		}
	}

	return nil, ErrNoSuchEntity
}

func (f *FileSystemGroceryStore) SetGroceryList(r *models.GroceryList) error {
	if err := f.v.Validate(r); err != nil {
		return &StoreValidationError{Err: err}
	}

	gls, err := f.GroceryLists()
	if err != nil {
		return err
	}

	var found bool
	for i, gl := range gls {
		if gl.Id == r.GetId() {
			gls[i] = r
			found = true
			break
		}
	}
	if !found {
		gls = append(gls, r)
	}

	_, _ = f.database.Seek(0, 0)
	_ = json.NewEncoder(f.database).Encode(gls)

	return nil
}

func (f *FileSystemGroceryStore) DeleteGroceryList(id string) error {
	gls, err := f.GroceryLists()
	if err != nil {
		return err
	}

	var found bool
	for i, gl := range gls {
		if gl.Id == id {
			gls = append(gls[:i], gls[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return ErrNoSuchEntity
	}

	_, _ = f.database.Seek(0, 0)
	_ = json.NewEncoder(f.database).Encode(gls)

	return nil
}
