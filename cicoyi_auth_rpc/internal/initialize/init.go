package initialize

import (
	"github.com/rustuo-m/cicoyi_auth_rpc/internal/common"
	"github.com/rustuo-m/cicoyi_auth_rpc/internal/rpc"
)

type Initialize struct {
	Log    common.LogInitialize
	Config common.ConfigInitialize
	Rpc    rpc.RpcInitialize
}
