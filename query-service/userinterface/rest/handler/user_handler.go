package handler

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/common/common_user_event"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryhandler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserHandler struct {
	app          *iris.Application
	subscribes   []ddd.Subscribe
	eventHandler ddd.QueryEventHandler
}

func NewUserHandler(app *iris.Application) *UserHandler {
	ctrl := &UserHandler{
		app:          app,
		eventHandler: queryhandler.NewUserQueryHandler(),
	}
	return ctrl
}

func (h *UserHandler) GetSubscribes() []ddd.Subscribe {
	return []ddd.Subscribe{
		{PubsubName: "", Topic: common_user_event.UserCreateEventType.String(), Route: "/users/user-create-event", Handle: h.onEvent},
		{PubsubName: "", Topic: common_user_event.UserUpdateEventType.String(), Route: "/users/user-update-event", Handle: h.onEvent},
		{PubsubName: "", Topic: common_user_event.UserDeleteEventType.String(), Route: "/users/user-delete-event", Handle: h.onEvent},
	}
}

func (h *UserHandler) RegisterSubscribe(subscribe ddd.Subscribe) {
	h.app.Handle("POST", subscribe.Route, subscribe.Handle.(iris.Handler))
}

func (h *UserHandler) onEvent(ctx iris.Context) {
	data, _ := ctx.GetBody()
	eventRecord := ddd.EventRecord{}
	err := json.Unmarshal(data, &eventRecord)
	if err != nil {
		ctx.SetErr(err)
	}
	err = h.eventHandler.OnEvent(&eventRecord)
}
