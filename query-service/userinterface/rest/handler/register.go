package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func Register(app *iris.Application) {
	ddd.RegisterSubscribe(NewUserHandler(app))
}
