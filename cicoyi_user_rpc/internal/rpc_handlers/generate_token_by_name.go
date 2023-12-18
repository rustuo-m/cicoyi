package rpc_handlers

import (
	"context"
	
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/rpc_service/pb"
)

func (receiver UserService) GenerateTokenByName(ctx context.Context, in *pb.RegisterByNameReq) (*pb.RegisterResponse, error) {
	return nil, nil
}
