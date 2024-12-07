package service

import (
	"context"

	"github.com/adwait-godbole/learning-grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userStore  UserStore
	jwtManager *JWTManager
}

// Login is a unary RPC to login user
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}
	return res, nil
}

// NewAuthServer returns a new AuthServer
func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}
