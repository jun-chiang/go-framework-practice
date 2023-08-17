package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	usermodel "github.com/jun-chiang/go-framework-practice/kitex-service/kitex_gen/usermodel/userservice"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	cli, err := clients.NewNamingClient(vo.NacosClientParam{
		// 客户端配置
		ClientConfig: &constant.ClientConfig{
			NamespaceId:         "go-framework-practice",
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogDir:              "/tmp/nacos/log",
			CacheDir:            "/tmp/nacos/cache",
			LogLevel:            "info",
		},
		// 配置Nacos的地址和端口
		ServerConfigs: []constant.ServerConfig{
			*constant.NewServerConfig("192.168.209.159", 8848),
		},
	})
	if err != nil {
		panic(err)
	}
	// 初始化server服务
	svr := usermodel.NewServer(
		new(UserServiceImpl),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "userService",
		}),
		// 配置当前RPC服务对外暴露的IP和端口
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8081,
		}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
