package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface/rest/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

func main() {
	app := iris.New()
	initDb()
	registerEventStorage("localhost", 9021, "pubsub")
	registerSubscribe(app)
	registerHttpServer(app)
	start(app, 9020)
}

func initDb() {
	err := mongodb.Init(&ddd_mongodb.Config{
		Host:         "192.168.64.4",
		DatabaseName: "example-query-service",
		UserName:     "dapr",
		Password:     "123456",
	})
	if err != nil {
		panic(err)
	}
}

func registerSubscribe(app *iris.Application) {
	app.Get("dapr/subscribe", func(context *context.Context) {
		_, _ = context.JSON(ddd.GetSubscribes())
	})
	handler.Register(app)
}

func registerEventStorage(host string, port int, pubsubName string) {
	eventStorage, err := ddd.NewDaprEventStorage(host, port, ddd.PubsubName(pubsubName))
	if err != nil {
		panic(err)
	}
	ddd.RegisterEventStorage("", eventStorage)
}

func registerHttpServer(app *iris.Application) {
	userinterface.RegisterMvcController(app)
}

func start(app *iris.Application, port int) {
	ddd.Start()
	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", port)))
}
