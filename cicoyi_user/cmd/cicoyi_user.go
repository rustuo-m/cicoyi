package main

import (
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/initialize"
)

func main() {
	start := initialize.Initialize{}
	start.Log.InitLog()
	start.Config.InitConfig()
	start.Mysql.InitMysql()
	start.Rpc.InitGrpc()
	start.Router.InitRouter()
}
