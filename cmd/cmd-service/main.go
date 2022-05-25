package main

import (
	"flag"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/event"
	_ "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/model"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/userinterface/rest/controller"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func main() {
	help := flag.Bool("help", false, "参数提示。")
	envType := flag.String("envType", "", "替换配置文件中的envType值。")
	config := flag.String("config", "./config/cmd-config.yaml", "配置文件。")
	flag.Parse()

	if *help {
		return
	}

	if _, err := restapp.RunWithConfig(*envType, *config, subscribes, controllers, events, restapp.DddActors); err != nil {
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
	return event.GetRegisterEventTypes()
}
