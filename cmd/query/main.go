package main

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/queryhandler"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/controller"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var app *iris.Application

func main() {
	configFile := "./config/query-config.yaml"
	app = iris.New()
	if err := restapp.RunWithConfig(configFile, app, subscribes, controllers, events); err != nil {
		panic(err)
	}
}

// 注册消息监听器
func subscribes() *[]restapp.RegisterSubscribe {
	return &[]restapp.RegisterSubscribe{
		queryhandler.NewUserSubscribes(),
	}
}

// 注册Http控制器
func controllers() *[]restapp.Controller {
	return &[]restapp.Controller{
		controller.NewUserController(),
	}
}

// 注册Http控制器
func events() *[]restapp.RegisterEventType {
	return xpublic.GetRegisterEventTypes()
}
