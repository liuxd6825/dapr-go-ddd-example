package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	_ "github.com/liuxd6825/dapr-go-ddd-example/common"
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
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
		Host:         "192.168.64.8:27019",
		DatabaseName: "query-example",
		UserName:     "query-example",
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
	// 注册User消息处理器
	_ = ddd.RegisterSubscribeHandler(newUserSubscribeHandler())
}

func registerEventStorage(host string, port int, pubsubName string) {
	hc, err := httpclient.NewHttpClient("localhost", 9011)
	if err != nil {
		panic(err)
	}
	applog.Init(hc, "query-example", applog.DEBUG)
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
		{PubsubName: "pubsub", Topic: user_events.UserCreateEventType.String(), Route: "/users/user-create-event"},
		{PubsubName: "pubsub", Topic: user_events.UserUpdateEventType.String(), Route: "/users/user-update-event"},
		{PubsubName: "pubsub", Topic: user_events.UserDeleteEventType.String(), Route: "/users/user-delete-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressCreateEventType.String(), Route: "/users/user-address-create-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressUpdateEventType.String(), Route: "/users/user-address-update-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressDeleteEventType.String(), Route: "/users/user-address-delete-event"},
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
