package facade

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

type BaseApi struct {
}

func (a BaseApi) SetError(ctx iris.Context, err error) {
	ctx.SetErr(err)
	ctx.StatusCode(http.StatusInternalServerError)
}
