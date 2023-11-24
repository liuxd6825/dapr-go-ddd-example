package main

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/register"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/liuxd6825/dapr-go-sdk/actor"
)

func main() {
	flag := restapp.NewRunFlag("./config/query-config.yaml")
	if flag.Help {
		return
	}

	opts := restapp.NewRunOptions().SetFlag(flag)
	if _, err := restapp.RunWithConfig(flag.Env, flag.Config, subscribes, controllers, events, actors, opts); err != nil {
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

func actors() []actor.Factory {
	return register.GetActors()
}
