package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface/rest/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func main() {
	app := iris.New()
	registerEventStorage("localhost", 9021, "pubsub")
	registerSubscribe(app)
	registerHttpServer(app)
	start(app, 9020)
}

func registerSubscribe(app *iris.Application) {
	app.Handle("GET", "/dapr/subscribe", func(context *context.Context) {
		data, _ := json.Marshal(ddd.GetSubscribes())
		_, _ = context.Write(data)
		context.ContentType("application/json")
	})
	handler.Register(app)
}

func registerEventStorage(host string, port int, pubsubName string) {
	eventStorage, err := ddd.NewEventStorage(host, port, ddd.PubsubName(pubsubName))
	if err != nil {
		panic(err)
	}
	ddd.RegisterEventStorage(eventStorage)
}

func registerHttpServer(app *iris.Application) {

	userinterface.RegisterMvcController(app)
}

func start(app *iris.Application, port int) {
	ddd.Start()
	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", port)))
}
