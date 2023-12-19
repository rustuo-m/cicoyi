package rpc

import (
	"log"

	"github.com/rustuo-m/cicoyi/cicoyi_user/rpc_service/pb"
	"google.golang.org/grpc"
)

type GrpcInitialize struct {
}

var AuthGrpcConn *grpc.ClientConn
var UserGrpcConn *grpc.ClientConn
var AuthGrpcClient pb.AuthServiceClient
var UserGrpcClient pb.AuthServiceClient

func (g GrpcInitialize) InitGrpc() {
	// 创建 gRPC 连接
	AuthGrpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v\n", err)
	}

	//UserGrpcConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	//if err != nil {
	//	log.Printf("Failed to connect to gRPC server: %v\n", err)
	//}

	// 创建 gRPC 客户端
	AuthGrpcClient = pb.NewAuthServiceClient(AuthGrpcConn)
}
