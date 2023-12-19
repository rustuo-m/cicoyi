package main

import (
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/initialize"
)

func main() {
	start := initialize.Initialize{}
	start.Log.InitLog()
	start.Config.InitConfig()
	start.Mysql.InitDatabase()
	start.Rpc.InitRpc()
}
