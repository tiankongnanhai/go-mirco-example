package account

import (
	"context"
	"fmt"
	"gomicro/proto"
	"log"

	"github.com/micro/go-micro/v2"
)

type AccountService struct{}

func (a *AccountService) AccountRegister(c context.Context, req *proto.ReqAccountRegister, res *proto.ResAccountRegister) error {
	fmt.Println("hit here")
	if req.Username == "guaosi" && req.Password == "guaosi" {
		res.Code = 0
		res.Message = "service success"
		return nil
	}
	res.Code = -1
	res.Message = "service failed"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("micro.service.account"),
	)
	// 初始化相关操作
	service.Init()
	// 注册实现了服务的handler
	if err := proto.RegisterAccountServiceHandler(service.Server(), new(AccountService)); err != nil {
		log.Print(err.Error())
		return
	}
	// 运行server
	if err := service.Run(); err != nil {
		log.Print(err.Error())
		return
	}
}
