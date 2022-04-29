package queryappservice

import (
	"context"
	"fmt"
	infr_dapr "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/infrastructure/idapr"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_errors"
)

var appId = "query-service"
var apiVersion = "v1.0"

type UserView struct {
	Id        string `json:"id" `
	TenantId  string `json:"tenantId" `
	UserCode  string `json:"userCode"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
}

func GetUserByUserId(ctx context.Context, tenantId, userId string) (res *UserView, isFound bool, err error) {
	defer func() {
		if e := ddd_errors.GetRecoverError(recover()); e != nil {
			err = e
		}
	}()
	resp := &UserView{}
	methodName := fmt.Sprintf("/api/%s/tenants/%s/users/%s", apiVersion, tenantId, userId)
	_, err = infr_dapr.GetClient().InvokeService(ctx, appId, methodName, "get", nil, resp)
	if err == nil {
		res = resp
		isFound = true
	}
	return res, isFound, err
}
