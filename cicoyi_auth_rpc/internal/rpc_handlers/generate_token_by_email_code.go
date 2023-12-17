package rpc_handlers

import (
	"context"

	"github.com/rustuo-m/cicoyi_auth_rpc/rpc_service/pb"
)

func (a AuthService) GenerateTokenByEmailCode(context.Context, *pb.AuthRequestByEmailCode) (*pb.AuthResponse, error) {
	return nil, nil
}
