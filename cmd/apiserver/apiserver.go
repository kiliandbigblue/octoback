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

	log.Info("grpc-ping: starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Info("Defaulting to port %s", zap.String("port", port))
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("net.Listen: %v", zap.Error(err))
	}

	s := grpc.NewServer()

	db, _ := os.OpenFile("database.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o666) //nolint:gosec //Permissive permissions.

	fs, err := corev1.NewFileSystemGroceryStore(db)
	if err != nil {
		log.Fatal("failed to create file system grocery store", zap.Error(err))
	}

	models.RegisterServiceServer(s, corev1.NewService(log, fs))

	reflection.Register(s)
	log.Info("Starting server", zap.String("port", port))
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve", zap.Error(err))
	}
}
