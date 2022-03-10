package controller

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_errors"
	"net/http"
)

const (
	ContentTypeApplicationJson = "application/json"
	ContentTypeTextPlain       = "text/plain"
)

type GetAction = func(ctx iris.Context) (interface{}, bool, error)

func ErrorNotFond(ctx iris.Context) {
	ctx.SetErr(iris.ErrNotFound)
	ctx.StatusCode(http.StatusNotFound)
	ctx.ContentType(ContentTypeTextPlain)
}

func ErrorInternalServerError(ctx iris.Context, err error) {
	ctx.SetErr(err)
	ctx.StatusCode(http.StatusInternalServerError)
	ctx.ContentType(ContentTypeTextPlain)
}

func ErrorVerifyError(ctx iris.Context, err error) {
	bytes, _ := json.Marshal(err)
	_, _ = ctx.Write(bytes)
	ctx.StatusCode(http.StatusInternalServerError)
	ctx.ContentType(ContentTypeTextPlain)
}

func SetError(ctx iris.Context, err error) {
	switch err.(type) {
	case *ddd_errors.NullError:
		ErrorNotFond(ctx)
		break
	case *ddd_errors.VerifyError:
		ErrorVerifyError(ctx, err)
		break
	default:
		ErrorInternalServerError(ctx, err)
		break
	}
}

func DoGet(ctx iris.Context, getAction GetAction) {
	data, ok, err := getAction(ctx)
	if err != nil {
		SetError(ctx, err)
		return
	}
	if !ok {
		ErrorNotFond(ctx)
		return
	}
	if data == nil {
		ErrorNotFond(ctx)
		return
	}
	setResponseData(ctx, data, nil)
}

func setResponseData(ctx iris.Context, data interface{}, err error) {
	if err != nil {
		SetError(ctx, err)
		return
	}
	ctx.ContentType(ContentTypeApplicationJson)
	_, _ = ctx.Write(getJsonBytes(data))
}

func getJsonBytes(data interface{}) []byte {
	bytes, _ := json.Marshal(data)
	return bytes
}
