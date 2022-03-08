package event_handler

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type SubscribeController struct{}

func NewSubscribeController(app *iris.Application) *SubscribeController {
	ctrl := &SubscribeController{}
	ctrl.register(app)
	return ctrl
}

func (m *SubscribeController) register(app *iris.Application) {
	app.Handle("GET", "/dapr/subscribe", m.getSubscribeData)
	app.Handle("POST", "/subscribe/user-create-event", m.onCreateUserEvent)
	app.Handle("POST", "/subscribe/user-update-event", m.onUpdateUserEvent)
}

func (m *SubscribeController) onCreateUserEvent(ctx iris.Context) {
	bytes, _ := ctx.GetBody()
	fmt.Println("create_user_event_handler")
	fmt.Println(string(bytes))
}

func (m *SubscribeController) onUpdateUserEvent(ctx iris.Context) {
	bytes, _ := ctx.GetBody()
	fmt.Println("update_user_event_handler")
	fmt.Println(string(bytes))
}

func (m *SubscribeController) getSubscribeData(ctx iris.Context) {
	subItems := &[]ddd.Subscription{
		{PubsubName: "pubsub", Topic: "user-create-event", Route: "/subscribe/user-create-event"},
		{PubsubName: "pubsub", Topic: "user-update-event", Route: "/subscribe/user-update-event"},
	}
	bytes, _ := json.Marshal(subItems)
	_, _ = ctx.Write(bytes)
}
