package userinterface

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/userinterface/rest/controller"
)

func RegisterMvcController(app *iris.Application) {
	mvc.Configure(app.Party("/api/v1.0"), registerMvcController)
}

func registerMvcController(app *mvc.Application) {
	app.Handle(controller.NewUserController())
}
