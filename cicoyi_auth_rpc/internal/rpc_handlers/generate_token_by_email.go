package rpc_handlers

import (
	"context"

	"github.com/rustuo-m/cicoyi_auth_rpc/rpc_service/pb"
)

func (a AuthService) GenerateTokenByEmail(context.Context, *pb.AuthRequestByEmail) (*pb.AuthResponse, error) {
	return nil, nil
}
