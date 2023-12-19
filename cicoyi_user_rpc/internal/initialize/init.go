package initialize

import (
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/common"
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/rpc"
)

type Initialize struct {
	Log    common.LogInitialize
	Config common.ConfigInitialize
	Rpc    rpc.RpcInitialize
	Mysql  common.Database
}
