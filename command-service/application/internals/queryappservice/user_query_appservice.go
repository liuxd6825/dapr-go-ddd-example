package queryappservice

import (
	"context"
	"fmt"
	infr_dapr "github.com/liuxd6825/dapr-go-ddd-example/command-service/infrastructure/idapr"
)

var appId = "query-example"
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

func GetUserByUserId(ctx context.Context, tenantId, userId string) (*UserView, bool, error) {
	resp := &UserView{}
	methodName := fmt.Sprintf("/api/%s/tenants/%s/users/%s", apiVersion, tenantId, userId)
	_, err := infr_dapr.GetClient().InvokeMethod(ctx, appId, methodName, "get", nil, resp)
	if err != nil {
		return nil, false, err
	}
	return resp, true, nil
}
