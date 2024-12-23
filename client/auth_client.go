package client

import (
	"context"
	"time"

	"github.com/adwait-godbole/learning-grpc/pb"
	"google.golang.org/grpc"
)

// AuthClient is a client to call authentication RPC
type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

// Login login user and returns the access token
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}

// NewAuthClient returns a new auth client
func NewAuthClient(cc *grpc.ClientConn, username, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}
