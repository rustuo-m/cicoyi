package initialize

import (
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/common"
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/models"
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/route"
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/rpc"
)

type Initialize struct {
	Log    common.LogInitialize
	Config common.ConfigInitialize
	Rpc    rpc.GrpcInitialize
	Router route.RouterInitialize
	Mysql  models.MysqlInitialize
}
