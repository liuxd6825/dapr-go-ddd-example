package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/userinterface/rest/controller"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic"
	_ "github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var app *iris.Application

func main() {
	fmt.Println("---------- cmd-service 开始启动 ----------")
	help := flag.Bool("help", false, "参数提示。")
	envType := flag.String("envType", "", "替换配置文件中的envType值。")
	config := flag.String("config", "./config/cmd-config.yaml", "配置文件。")
	flag.Parse()

	if *help {
		return
	}

	app = iris.New()
	if err := restapp.RunWithConfig(*envType, *config, app, subscribes, controllers, events); err != nil {
		panic(err)
	}
}

// 注册消息监听器
func subscribes() *[]restapp.RegisterSubscribe {
	return &[]restapp.RegisterSubscribe{}
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
