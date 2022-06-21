package serviceclient

import (
	"gomicro/handler"
	"gomicro/proto"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

var Port string

func RegisterService() {
	// 连接服务注册中心
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "p",
				Usage: "port",
			},
		),
	)
	// 解析命令行参数
	// 我们希望可以使用 -p 参数来手动指定我们HTTP服务对外提供服务时的端口
	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			Port = ctx.String("p")
			if Port == "" {
				Port = "8091"
			}
			return nil
		},
		),
	)
	// 复用服务注册的客户端
	client := service.Client()
	// 获取在服务注册中心上 micro.service.account 的客户端
	handler.AccountServiceClient = proto.NewAccountService("micro.service.account", client)
}
