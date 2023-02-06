package main

import (
	"flag"
	"fmt"

	"go-zero-nacos-demo/internal/config"
	"go-zero-nacos-demo/internal/server"
	"go-zero-nacos-demo/internal/svc"
	"go-zero-nacos-demo/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var bootstrapFile = flag.String("f", "etc/bootstrap.yaml", "the config file")

func main() {
	flag.Parse()
	var bootstrapConfig config.BootstrapConfig
	conf.MustLoad(*bootstrapFile, &bootstrapConfig)

	//解析业务配置
	var c config.Config
	nacosConfig := bootstrapConfig.NacosConfig
	serviceConfigContent := nacosConfig.InitConfig(func(data string) {
		err := conf.LoadFromYamlBytes([]byte(data), &c)
		if err != nil {
			panic(err)
		}
	})
	err := conf.LoadFromYamlBytes([]byte(serviceConfigContent), &c)
	if err != nil {
		panic(err)
	}
	// 注册到nacos
	nacosConfig.Discovery(c)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
