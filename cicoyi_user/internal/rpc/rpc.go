package rpc

import (
	"log"
	
	"github.com/rustuo-m/cicoyi/cicoyi_user/rpc_service/pb"
	"google.golang.org/grpc"
)

type GrpcInitialize struct {
}

var GrpcConn *grpc.ClientConn
var GrpcClient pb.AuthServiceClient

func (g GrpcInitialize) InitGrpc() {
	// 创建 gRPC 连接
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v\n", err)
	}
	
	GrpcConn = grpcConn
	
	// 创建 gRPC 客户端
	GrpcClient = pb.NewAuthServiceClient(GrpcConn)
}
