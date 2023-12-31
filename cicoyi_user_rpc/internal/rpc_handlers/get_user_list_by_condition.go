package rpc_handlers

import (
	"context"
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/service"

	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/rpc_service/pb"
)

func (receiver UserService) GetUserListByCondition(ctx context.Context, in *pb.GetUserListByConditionReq) (*pb.GetUserListByConditionResp, error) {
	return service.UserModel{}.GetUserListByCondition(ctx, in)
}
