package main

import (
	"fmt"
	config "funnel/app/config"
	grpccanteen "funnel/app/controller/grpc/canteen"
	"funnel/router"
	rpc "funnel/rpc"
	"net"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func startOpenApi() {
	r := gin.Default()
	config.SetupConfigs(r)
	router.SetupRouter(r)

	r.Run()
}

func startRpc() {
	// 监听本地端口
	lis, err := net.Listen("tcp", "0.0.0.0:8890")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}
	fmt.Println("Now Listen: 8890")
	// 创建gRPC服务器
	s := grpc.NewServer()
	fmt.Println("Now start: grpc")
	// 注册服务
	rpc.RegisterCanteenServiceServer(s, &grpccanteen.CanteenRpc{})
	fmt.Println("Now start: grpc")

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
	fmt.Printf("RPC!")
}

func main() {
	go startOpenApi()
	startRpc()
}
