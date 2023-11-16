package main

import (
	"net"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	corev1 "github.com/kiliandbigblue/octoback/octoback/core/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type InMemoryStore struct {
	groceryLists map[string]*models.GroceryList
}

func (s *InMemoryStore) GetGroceryList(id string) (*models.GroceryList, error) {
	gli, ok := s.groceryLists[id]
	if ok {
		return gli, nil
	}
	return nil, corev1.ErrNoSuchEntity
}

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen", zap.Error(err))
	}
	s := grpc.NewServer()
	models.RegisterServiceServer(s, corev1.NewService(log, &InMemoryStore{
		groceryLists: map[string]*models.GroceryList{
			"gl_1": {
				Id:   "gl_1",
				Name: "My first grocery list",
				Items: []*models.GroceryItem{
					{
						Id:       "gi_1",
						Name:     "Banana",
						Quantity: 3,
					},
					{
						Id:       "gi_2",
						Name:     "Apple",
						Quantity: 2,
					},
				},
			},
		},
	}))

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve", zap.Error(err))
	}
}
