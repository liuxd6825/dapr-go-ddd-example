package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/rest/controller"
	controller2 "github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface/rest/handler"
)

func Init(port int) {
	app := iris.New()
	app.Use(before)
	controller2.newUserHandler(app)
	mvc.Configure(app.Party("/api/v1.0"), register)
	_ = app.Run(iris.Addr(fmt.Sprintf(":%d", port)))
}

func register(app *mvc.Application) {
	app.Handle(controller.NewUserController())
}

func before(irisCtx iris.Context) {
	irisCtx.Next() // 执行下一个处理器。
}
