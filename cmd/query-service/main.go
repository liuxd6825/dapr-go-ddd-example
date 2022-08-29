package main

import (
	"flag"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/register"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func main() {
	help := flag.Bool("help", false, "参数提示。")
	env := flag.String("env", "", "替换配置文件中的env值。")
	config := flag.String("config", "./config/query-config.yaml", "配置文件。")
	flag.Parse()

	if *help {
		return
	}

	if _, err := restapp.RunWithConfig(*env, *config, subscribes, controllers, events, restapp.Actors); err != nil {
		panic(err)
	}
}

// 注册消息监听器
func subscribes() []restapp.RegisterSubscribe {
	return register.GetRegisterSubscribe()
}

// 注册Http控制器
func controllers() []restapp.Controller {
	return register.GetRegisterController()
}

// 注册领域事件
func events() []restapp.RegisterEventType {
	return register.GetRegisterEventType()
}
