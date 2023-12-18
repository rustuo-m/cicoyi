package rpc_handlers

import (
	"context"
	
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/rpc_service/pb"
)

func (receiver UserService) SetPassword(ctx context.Context, in *pb.SetPasswordReq) (*pb.SetPasswordResp, error) {
	return nil, nil
}
