package rpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/rpc_handlers"
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/rpc_service/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type RpcInitialize struct {
}

func (r RpcInitialize) InitRpc() {
	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, rpc_handlers.UserService{})
	// 开启监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("port")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()
	
	// 开启 gRPC 服务器
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	
	log.Println("The user rpc service has been started, and the port is :", viper.GetString("port"))
	
	// 等待程序关闭信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	// 关闭服务器
	grpcServer.GracefulStop()
	log.Println("---------------------The program has been closed---------------------------")
}
