package main

import (
	"github.com/kataras/iris/v12"
	_ "github.com/liuxd6825/dapr-go-ddd-example/common"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryhandler"
	_ "github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/db"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface/rest/controller"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/daprclient"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/rest"
)

var app *iris.Application

func main() {
	//创建dapr客户端
	hc, err := daprclient.NewHttpClient("localhost", 9011)
	if err != nil {
		panic(err)
	}

	//创建dapr事件存储器
	eventStorage, err := ddd.NewDaprEventStorage(hc, ddd.PubsubName("pubsub"))
	if err != nil {
		panic(err)
	}
	esMap := map[string]ddd.EventStorage{}
	esMap[eventStorage.GetPubsubName()] = eventStorage

	// 注册消息监听器
	handlers := &[]rest.RegisterSubscribe{
		queryhandler.NewUserSubscribeHandler(),
	}

	// 注册Http控制器
	controllers := &[]rest.Controller{
		controller.NewUserController(),
	}

	// 设置启动参数
	options := rest.StartOptions{
		AppId:      "query-example",
		AppPort:    9020,
		LogLevel:   applog.DEBUG,
		HttpClient: hc,
	}

	// 启动服务
	app = iris.New()
	if err := rest.Start(options, app, "/api/v1.0", handlers, controllers, esMap); err != nil {
		panic(err)
	}
}
