package auth

import (
	"context"
	ssov1 "github.com/Nor1ch/protos/gen/go/sso"
	"google.golang.org/grpc"
)

// serverAPI mock for debugging without methods implementation
type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

// Register Registration handlers for gRPC server
func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

// Login login handler
func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

// Register Register handler
func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

// IsAdmin IsAdmin handler
func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
