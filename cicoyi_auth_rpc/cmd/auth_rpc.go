package main

import "github.com/rustuo-m/cicoyi_auth_rpc/internal/initialize"

func main() {
	start := initialize.Initialize{}
	start.Log.InitLog()
	start.Config.InitConfig()
	start.Rpc.InitRpc()
}
