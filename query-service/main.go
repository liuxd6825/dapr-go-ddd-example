package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/liuxd6825/dapr-go-ddd-example/common/event_type"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryhandler"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/httpclient"
)

var app *iris.Application

func main() {
	app = iris.New()
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
	// dapr 服务通过访问http://locahost:<port>/dapr/subscribe获取订阅的消息
	app.Get("dapr/subscribe", func(context *context.Context) {
		_, _ = context.JSON(ddd.GetSubscribes())
	})

	_ = ddd.RegisterEventType(event_type.UserCreateEventType.String(), "1.0", func() interface{} {
		return user_events.NewUserCreateEventV1()
	})
	_ = ddd.RegisterEventType(event_type.UserUpdateEventType.String(), "1.0", func() interface{} {
		return user_events.NewUserUpdateEvent()
	})
	_ = ddd.RegisterEventType(event_type.UserDeleteEventType.String(), "1.0", func() interface{} {
		return user_events.NewUserDeleteEvent()
	})

	// 注册User消息处理器
	_ = ddd.RegisterSubscribeHandler(newUserSubscribeHandler())
}

func registerEventStorage(host string, port int, pubsubName string) {
	hc, err := httpclient.NewHttpClient("localhost", 9011)
	if err != nil {
		panic(err)
	}
	applog.Init(hc, "query-example")
	eventStorage, err := ddd.NewDaprEventStorage(hc, ddd.PubsubName(pubsubName))
	if err != nil {
		panic(err)
	}
	ddd.RegisterEventStorage("", eventStorage)
}

func registerHttpServer(app *iris.Application) {
	userinterface.RegisterMvcController(app)
}

func start(app *iris.Application, port int) {
	_ = ddd.Start()
	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", port)))
}

func newUserSubscribeHandler() ddd.SubscribeHandler {
	subscribes := []ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event_type.UserCreateEventType.String(), Route: "/users/user-create-event"},
		{PubsubName: "pubsub", Topic: event_type.UserUpdateEventType.String(), Route: "/users/user-update-event"},
		{PubsubName: "pubsub", Topic: event_type.UserDeleteEventType.String(), Route: "/users/user-delete-event"},
	}
	return ddd.NewSubscribeHandler(subscribes, queryhandler.NewUserQueryHandler(), subscribeHandler)
}

func subscribeHandler(sh ddd.SubscribeHandler, subscribe ddd.Subscribe) error {
	app.Handle("POST", subscribe.Route, func(c *context.Context) {
		if err := sh.CallQueryEventHandler(c, c); err != nil {
			c.SetErr(err)
		}
	})
	return nil
}
