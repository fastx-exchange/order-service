package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"order-service/config"
	"order-service/pb"
	"order-service/src/pkg/database"
	"order-service/src/services"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := services.NewUserService(db)
	pb.RegisterUserServiceServer(grpcServer, userService)

	reflection.Register(grpcServer)

	log.Println("User Service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
