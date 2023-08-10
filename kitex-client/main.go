package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	usermodel "github.com/jun-chiang/go-framework-practice/kitex-client/kitex_gen/usermodel/userservice"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	cli, err := clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig: &constant.ClientConfig{
			NamespaceId:         "go-framework-practice",
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogDir:              "/tmp/nacos/log",
			CacheDir:            "/tmp/nacos/cache",
			LogLevel:            "info",
		},
		ServerConfigs: []constant.ServerConfig{
			*constant.NewServerConfig("192.168.209.159", 8848),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// 创建客户端
	clt := usermodel.MustNewClient("userService",
		client.WithResolver(resolver.NewNacosResolver(cli)),
	)
	for i := 0; i < 10; i++ {
		id := rand.Int63n(3)
		user, err := clt.GetUserById(context.Background(), id, callopt.WithConnectTimeout(10*time.Second))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("请求返回的结果是：%+v\n", user)
		time.Sleep(2 * time.Second)
	}
}
