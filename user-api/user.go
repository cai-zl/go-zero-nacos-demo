package main

import (
	"flag"
	"fmt"

	"go-zero-nacos-demo/user-api/internal/config"
	"go-zero-nacos-demo/user-api/internal/handler"
	"go-zero-nacos-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var bootstrapConfig config.BootstrapConfig
	conf.MustLoad(*configFile, &bootstrapConfig)

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

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
