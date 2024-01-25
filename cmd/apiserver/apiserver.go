package main

import (
	"net"
	"os"

	models "github.com/kiliandbigblue/octoback/gen/proto/go/octoback/core/v1"
	corev1 "github.com/kiliandbigblue/octoback/internal/core/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("failed to listen", zap.Error(err))
	}
	s := grpc.NewServer()

	db, _ := os.OpenFile("database.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666) //nolint:gosec //Permissive permissions.

	fs := corev1.NewFileSystemGroceryStore(db)
	models.RegisterServiceServer(s, corev1.NewService(log, fs))

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve", zap.Error(err))
	}
}
