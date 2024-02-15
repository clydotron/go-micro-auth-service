package main

import (
	"context"
	"log"
	"net"

	"github.com/clydotron/go-micro-auth-service/data"
	auth "github.com/clydotron/go-micro-auth-service/protos"

	"google.golang.org/grpc"
)

var gRPCPort = "50001"

// AuthServer - implemenation of gRPC interface
type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	UserRepo *data.PostgresUserRepo
}

// Authenticate - accessable via gRPC
func (a *AuthServer) Authenticate(ctx context.Context, req *auth.AuthRequest) (*auth.AuthResponse, error) {

	user, err := a.UserRepo.GetByEmail(req.GetEmail())
	if err != nil {
		return &auth.AuthResponse{Result: "Invalid credentials"}, nil
	}

	valid, err := user.PasswordMatches(req.GetPassword())
	if err != nil || !valid {
		return &auth.AuthResponse{Result: "Invalid credentials"}, nil
	}

	return &auth.AuthResponse{Result: "Authenticated"}, nil
}

func (app *App) gRPCListen() {
	listener, err := net.Listen("tcp", ":"+gRPCPort)
	if err != nil {
		log.Fatalln("Failed to listen for gRPC:", err)
	}

	srv := grpc.NewServer()
	auth.RegisterAuthServiceServer(srv, &AuthServer{UserRepo: app.UserRepo})

	log.Println("gRPC server started on port", gRPCPort)

	if err := srv.Serve(listener); err != nil {
		log.Fatalln("Failed to listen for gRPC:", err)
	}
}
