package service

import (
	"context"
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/model"
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/rpc_service/pb"
)

type UserModel struct {
	model.User
}

func (u UserModel) GetUserListByCondition(ctx context.Context, in *pb.GetUserListByConditionReq) (out *pb.GetUserListByConditionResp, err error) {
	var users []model.User
	DB().
		Where("name LIKE ?", "%"+in.Name+"%").
		Or("email LIKE ?", "%"+in.Email+"%").
		Limit(int(in.PageSize)).
		Offset(int((in.PageNumber - 1) * in.PageSize)).
		Find(&users)

	for i := 0; i < len(users); i++ {
		out.UserList[i].Name = users[i].Name
		out.UserList[i].Email = users[i].Email
		out.UserList[i].Birth = users[i].Birth
		out.UserList[i].Gender = users[i].Gender
		out.UserList[i].Phone = users[i].Phone
		out.UserList[i].Id = int64(users[i].ID)
	}
	return out, nil
}
