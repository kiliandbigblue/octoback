package store

import (
	"errors"
	"testing"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/groceries/v1"
	"github.com/kiliandbigblue/octoback/internal/x/testhelper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemGroceryStore(database)
		assert.NoError(t, err)
	})

	t.Run("grocery lists from a reader", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, `[
			{"Id": "gl_12345671", "Name": "My first grocery list", "Items": [ {"Id": 1, "Name": "Banana", "Quantity": 3}, {"Id": 2, "Name": "Apple", "Quantity": 2} ]},
			{"Id": "gl_12345672", "Name": "My second grocery list", "Items": [ {"Id": 1, "Name": "Orange", "Quantity": 5}, {"Id": 2, "Name": "Pear", "Quantity": 1} ]}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemGroceryStore(database)
		assert.NoError(t, err)

		got, err := store.GroceryLists()
		assert.NoError(t, err)

		want := []*models.GroceryList{
			{
				Id:   "gl_12345671",
				Name: "My first grocery list",
				Items: []*models.GroceryItem{
					{
						Id:       1,
						Name:     "Banana",
						Quantity: 3,
					},
					{
						Id:       2,
						Name:     "Apple",
						Quantity: 2,
					},
				},
			},
			{
				Id:   "gl_12345672",
				Name: "My second grocery list",
				Items: []*models.GroceryItem{
					{
						Id:       1,
						Name:     "Orange",
						Quantity: 5,
					},
					{
						Id:       2,
						Name:     "Pear",
						Quantity: 1,
					},
				},
			},
		}

		assert.Equal(t, got, want)
	})

	t.Run("get grocery list", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, `[
			{"Id": "gl_12345671", "Name": "My first grocery list", "Items": [ {"Id": 1, "Name": "Banana", "Quantity": 3}, {"Id": 2, "Name": "Apple", "Quantity": 2} ]},
			{"Id": "gl_12345672", "Name": "My second grocery list", "Items": [ {"Id": 1, "Name": "Orange", "Quantity": 5}, {"Id": 2, "Name": "Pear", "Quantity": 1} ]}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemGroceryStore(database)

		got, err := store.GroceryList("gl_12345672")

		want := &models.GroceryList{
			Id:   "gl_12345672",
			Name: "My second grocery list",
			Items: []*models.GroceryItem{
				{
					Id:       1,
					Name:     "Orange",
					Quantity: 5,
				},
				{
					Id:       2,
					Name:     "Pear",
					Quantity: 1,
				},
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, got, want)
	})

	t.Run("set grocery list", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, `[
			{"Id": "gl_12345671", "Name": "My first grocery list", "Items": [ {"Id": 1, "Name": "Banana", "Quantity": 3}, {"Id": 2, "Name": "Apple", "Quantity": 2} ]},
			{"Id": "gl_12345672", "Name": "My second grocery list", "Items": [ {"Id": 1, "Name": "Orange", "Quantity": 5}, {"Id": 2, "Name": "Pear", "Quantity": 1} ]}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemGroceryStore(database)

		input := &models.GroceryList{
			Id:   "gl_12345672",
			Name: "My better second grocery list",
			Items: []*models.GroceryItem{
				{
					Id:       1,
					Name:     "Orange",
					Quantity: 5,
				},
				{
					Id:       2,
					Name:     "Pear",
					Quantity: 1,
				},
				{
					Id:       3,
					Name:     "Banana",
					Quantity: 1,
				},
			},
		}

		err := store.SetGroceryList(input)

		assert.NoError(t, err)

		got, err := store.GroceryList("gl_12345672")

		assert.NoError(t, err)
		assert.True(t, proto.Equal(input, got))
	})

	t.Run("set grocery list with invalid data", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, `[
			{"Id": "gl_12345671", "Name": "My first grocery list", "Items": [ {"Id": 1, "Name": "Banana", "Quantity": 3}, {"Id": 2, "Name": "Apple", "Quantity": 2} ]},
			{"Id": "gl_12345672", "Name": "My second grocery list", "Items": [ {"Id": 1, "Name": "Orange", "Quantity": 5}, {"Id": 2, "Name": "Pear", "Quantity": 1} ]}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemGroceryStore(database)

		input := &models.GroceryList{
			Id:   "gl_12345673",
			Name: "",
		}

		err := store.SetGroceryList(input)
		assert.Error(t, err)
		var sve *StoreValidationError
		assert.True(t, errors.As(err, &sve))
	})

	t.Run("set grocery list with new ID", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, `[
			{"Id": "gl_12345671", "Name": "My first grocery list", "Items": [ {"Id": 1, "Name": "Banana", "Quantity": 3}, {"Id": 2, "Name": "Apple", "Quantity": 2} ]},
			{"Id": "gl_12345672", "Name": "My second grocery list", "Items": [ {"Id": 1, "Name": "Orange", "Quantity": 5}, {"Id": 2, "Name": "Pear", "Quantity": 1} ]}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemGroceryStore(database)

		input := &models.GroceryList{
			Id:   "gl_12345673",
			Name: "My better second grocery list",
			Items: []*models.GroceryItem{
				{
					Id:       1,
					Name:     "Orange",
					Quantity: 5,
				},
				{
					Id:       2,
					Name:     "Pear",
					Quantity: 1,
				},
				{
					Id:       3,
					Name:     "Banana",
					Quantity: 1,
				},
			},
		}

		err := store.SetGroceryList(input)

		assert.NoError(t, err)

		got, err := store.GroceryList("gl_12345673")

		assert.NoError(t, err)
		assert.True(t, proto.Equal(input, got))
	})

	t.Run("delete grocery list", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, `[
			{"Id": "gl_12345671", "Name": "My first grocery list", "Items": [ {"Id": 1, "Name": "Banana", "Quantity": 3}, {"Id": 2, "Name": "Apple", "Quantity": 2} ]},
			{"Id": "gl_12345672", "Name": "My second grocery list", "Items": [ {"Id": 1, "Name": "Orange", "Quantity": 5}, {"Id": 2, "Name": "Pear", "Quantity": 1} ]}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemGroceryStore(database)

		err := store.DeleteGroceryList("gl_12345672")

		assert.NoError(t, err)

		got, err := store.GroceryList("gl_12345672")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrNoSuchEntity))
		assert.Nil(t, got)
	})

	t.Run("delete grocery list that does not exist", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, `[
			{"Id": "gl_12345671", "Name": "My first grocery list", "Items": [ {"Id": 1, "Name": "Banana", "Quantity": 3}, {"Id": 2, "Name": "Apple", "Quantity": 2} ]},
			{"Id": "gl_12345672", "Name": "My second grocery list", "Items": [ {"Id": 1, "Name": "Orange", "Quantity": 5}, {"Id": 2, "Name": "Pear", "Quantity": 1} ]}
		]`)
		defer cleanDatabase()

		store, _ := NewFileSystemGroceryStore(database)

		err := store.DeleteGroceryList("gl_12345673")

		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrNoSuchEntity))
	})

	t.Run("grocery lists are storted", func(t *testing.T) {
		database, cleanDatabase := testhelper.CreateTempFile(t, "")
		defer cleanDatabase()

		store, _ := NewFileSystemGroceryStore(database)

		gls := []*models.GroceryList{
			testhelper.FakeGroceryList(),
			testhelper.FakeGroceryList(),
			testhelper.FakeGroceryList(),
		}
		gls[2].Id = "gl_12345671"
		gls[1].Id = "gl_12345672"
		gls[0].Id = "gl_12345673"

		for _, gl := range gls {
			err := store.SetGroceryList(gl)
			assert.NoError(t, err)
		}

		got, err := store.GroceryLists()
		assert.NoError(t, err)

		assert.Equal(t, []*models.GroceryList{gls[2], gls[1], gls[0]}, got)
	})
}
