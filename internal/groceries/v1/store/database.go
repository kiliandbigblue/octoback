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
func (gs *GroceryStore) GroceryLists(ctx context.Context) ([]*models.GroceryList, error) {
	slis, err := gs.db.ListGroceryLists(ctx)
	if err != nil {
		return nil, &StoreInternalError{Err: err}
	}

	mlis := make([]*models.GroceryList, len(slis))
	for i, sli := range slis {
		sitms, err := gs.db.ListGroceryItemsByGroceryList(ctx, sli.ID)
		if err != nil {
			return nil, &StoreInternalError{Err: err}
		}
		mlis[i] = gs.translate(sli, sitms)
	}

	return mlis, nil
}

// GroceryList returns a grocery list by its ID.
func (gs *GroceryStore) GroceryList(ctx context.Context, id int64) (*models.GroceryList, error) {
	sli, err := gs.db.GetGroceryList(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoSuchEntity
		}
		return nil, &StoreInternalError{Err: err}
	}

	sitms, err := gs.db.ListGroceryItemsByGroceryList(ctx, sli.ID)
	if err != nil {
		return nil, &StoreInternalError{Err: err}
	}

	return gs.translate(sli, sitms), nil
}

func (gs *GroceryStore) CreateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error) {
	if err := gs.v.Validate(r); err != nil {
		return nil, &StoreValidationError{Err: err}
	}

	sli, err := gs.db.CreateGroceryList(ctx, r.GetName())
	if err != nil {
		return nil, &StoreInternalError{Err: err}
	}

	return gs.translate(sli, nil), nil
}

func (gs *GroceryStore) UpdateGroceryList(ctx context.Context, r *models.GroceryList) (*models.GroceryList, error) {
	if err := gs.v.Validate(r); err != nil {
		return nil, &StoreValidationError{Err: err}
	}

	sli, err := gs.db.UpdateGroceryList(ctx, database.UpdateGroceryListParams{
		ID:   r.GetId(),
		Name: r.GetName(),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoSuchEntity
		}
	}

	sitms, err := gs.db.ListGroceryItemsByGroceryList(ctx, sli.ID)
	if err != nil {
		return nil, &StoreInternalError{Err: err}
	}

	return gs.translate(sli, sitms), nil
}

func (gs *GroceryStore) DeleteGroceryList(ctx context.Context, id int64) error {
	if _, err := gs.db.DeleteGroceryList(ctx, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoSuchEntity
		}
		return &StoreInternalError{Err: err}
	}

	return nil
}

func (gs *GroceryStore) CreateGroceryItem(ctx context.Context, groceryList int64, r *models.GroceryItem) (*models.GroceryItem, error) {
	if err := gs.v.Validate(r); err != nil {
		return nil, &StoreValidationError{Err: err}
	}

	sit, err := gs.db.CreateGroceryItem(ctx, database.CreateGroceryItemParams{
		GroceryListID: groceryList,
		Name:          r.GetName(),
		Quantity:      r.GetQuantity(),
		Checked:       r.GetChecked(),
	})
	if err != nil {
		return nil, &StoreInternalError{Err: err}
	}

	return gs.translateGroceryItem(sit), nil
}

func (gs *GroceryStore) translate(sgli database.GroceryList, sglitms []database.GroceryItem) *models.GroceryList {
	mglitms := make([]*models.GroceryItem, len(sglitms))
	for i, sglitm := range sglitms {
		mglitms[i] = gs.translateGroceryItem(sglitm)
	}

	return &models.GroceryList{
		Id:         sgli.ID,
		Name:       sgli.Name,
		Items:      mglitms,
		CreateTime: timestamppb.New(sgli.CreatedAt),
	}
}

func (gs *GroceryStore) translateGroceryItem(it database.GroceryItem) *models.GroceryItem {
	return &models.GroceryItem{
		Id:         it.ID,
		Name:       it.Name,
		Quantity:   it.Quantity,
		Checked:    it.Checked,
		CreateTime: timestamppb.New(it.CreatedAt),
	}
}
