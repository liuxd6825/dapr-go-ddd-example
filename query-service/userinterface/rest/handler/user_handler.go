package handler

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/liuxd6825/dapr-go-ddd-example/common/event_type"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryhandler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserHandler struct {
	app               *iris.Application
	subscribes        []ddd.Subscribe
	queryEventHandler ddd.QueryEventHandler
}

func NewUserHandler(app *iris.Application) *UserHandler {
	ctrl := &UserHandler{
		app:               app,
		queryEventHandler: queryhandler.NewUserQueryHandler(),
	}
	return ctrl
}

func (h *UserHandler) GetSubscribes() (*[]ddd.Subscribe, error) {
	es, err := ddd.GetEventStorage("")
	if err != nil {
		return nil, err
	}
	pubsubName := es.GetPubsubName()
	return &[]ddd.Subscribe{
		{PubsubName: pubsubName, Topic: event_type.UserCreateEventType.String(), Route: "/users/user-create-event"},
		{PubsubName: pubsubName, Topic: event_type.UserUpdateEventType.String(), Route: "/users/user-update-event"},
		{PubsubName: pubsubName, Topic: event_type.UserDeleteEventType.String(), Route: "/users/user-delete-event"},
	}, nil
}

func (h *UserHandler) RegisterSubscribe(subItem ddd.Subscribe) error {
	h.app.Handle("POST", subItem.Route, h.onEventHandler)
	return nil
}

func (h *UserHandler) onEventHandler(ctx *context.Context) {
	data, _ := ctx.GetBody()
	ddd.EventRecordJsonMarshal(data).OnSuccess(func(eventRecord *ddd.EventRecord) error {
		return ddd.CallEventHandler(ctx, h.queryEventHandler, eventRecord)
	}).OnError(func(err error) {
		ctx.SetErr(err)
	})
}
