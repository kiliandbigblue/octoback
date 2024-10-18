package store

import (
	"context"
	"database/sql"

	"github.com/bufbuild/protovalidate-go"
	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	"github.com/kiliandbigblue/octoback/internal/database"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrNoSuchEntity = errors.New("no such entity")

type StoreValidationError struct {
	Err error
}

func (e *StoreValidationError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return "store validation error"
}

type StoreInternalError struct {
	Err error
}

func (e *StoreInternalError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return "store internal error"
}

type GroceryStore struct {
	v  *protovalidate.Validator
	db *database.Queries
}

func NewGroceryStore(db database.DBTX) *GroceryStore {
	v, _ := protovalidate.New()

	return &GroceryStore{
		db: database.New(db),
		v:  v,
	}
}

// GroceryLists returns all grocery lists.
func (f *GroceryStore) GroceryLists(ctx context.Context) ([]*models.GroceryList, error) {
	groceryLists, err := f.db.ListGroceryLists(ctx)
	if err != nil {
		return nil, &StoreInternalError{Err: err}
	}

	glis := make([]*models.GroceryList, 0, len(groceryLists))
	for _, gl := range groceryLists {
		glis = append(glis, f.translate(gl))
	}

	return glis, nil
}

// GroceryList returns a grocery list by its ID.
func (gs *GroceryStore) GroceryList(ctx context.Context, id int64) (*models.GroceryList, error) {
	gl, err := gs.db.GetGroceryList(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoSuchEntity
		}
		return nil, &StoreInternalError{Err: err}
	}
	return gs.translate(gl), nil
}

func (gs *GroceryStore) CreateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error) {
	if err := gs.v.Validate(r); err != nil {
		return nil, &StoreValidationError{Err: err}
	}

	created, err := gs.db.CreateGroceryList(ctx, r.GetName())
	if err != nil {
		return nil, &StoreInternalError{Err: err}
	}

	return gs.translate(created), nil
}

func (gs *GroceryStore) UpdateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error) {
	if err := gs.v.Validate(r); err != nil {
		return nil, &StoreValidationError{Err: err}
	}

	updated, err := gs.db.UpdateGroceryList(ctx, database.UpdateGroceryListParams{
		ID:   r.GetId(),
		Name: r.GetName(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoSuchEntity
		}
	}

	return gs.translate(updated), nil
}

func (f *GroceryStore) DeleteGroceryList(ctx context.Context, id int64) error {
	if _, err := f.db.DeleteGroceryList(ctx, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoSuchEntity
		}
		return &StoreInternalError{Err: err}
	}

	return nil
}

func (f *GroceryStore) translate(r database.GroceryList) *models.GroceryList {
	return &models.GroceryList{
		Id:         r.ID,
		Name:       r.Name,
		CreateTime: timestamppb.New(r.CreatedAt),
	}
}
