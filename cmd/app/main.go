package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"order-service/config"
	"order-service/pb"
	"order-service/src/pkg/database"
	"order-service/src/services"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Fetch port from environment variable or use a default
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051" // Default port
	}

	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	userService := services.NewUserService(db)
	pb.RegisterUserServiceServer(grpcServer, userService)

	reflection.Register(grpcServer)

	log.Printf("User Service is running on port %s\n", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
