package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/rest/controller"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/rest/event-handler"
)

func Init(port int) {
	app := iris.New()
	app.Use(before)
	event_handler.NewSubscribeController(app)
	mvc.Configure(app.Party("/api/v1.0"), registerMvcController)
	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", port)))
}

func registerMvcController(app *mvc.Application) {
	app.Handle(controller.NewUserController())
}

func before(irisCtx iris.Context) {
	irisCtx.Next() // 执行下一个处理器。
}
