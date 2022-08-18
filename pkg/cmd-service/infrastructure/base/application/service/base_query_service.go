package service

import (
	"context"
	"fmt"
	"github.com/liuxd6825/dapr-go-ddd-sdk/daprclient"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

type BaseQueryAppService struct {
	QueryAppId   string
	ResourceName string
	ApiVersion   string
}

func (s *BaseQueryAppService) Init(queryAppId, resourceName, apiVersion string) {
	s.QueryAppId = queryAppId
	s.ResourceName = resourceName
	s.ApiVersion = apiVersion
}

func (s *BaseQueryAppService) QueryById(ctx context.Context, tenantId, id string, resData interface{}) (isFound bool, err error) {
	return s.QueryData(ctx, tenantId, "/"+id, nil, resData)
}

func (s *BaseQueryAppService) QueryByIds(ctx context.Context, tenantId string, ids []string, resData interface{}) (isFound bool, err error) {
	idParams := ""
	count := len(ids)
	for i, id := range ids {
		idParams = idParams + fmt.Sprintf("id=%v", id)
		if i < count-2 {
			idParams += "&"
		}
	}
	methodName := fmt.Sprintf(":getById?%v", idParams)
	return s.QueryData(ctx, tenantId, methodName, nil, resData)
}

func (s *BaseQueryAppService) QueryData(ctx context.Context, tenantId, methodName string, req interface{}, resData interface{}) (isFound bool, err error) {
	defer func() {
		if e := errors.GetRecoverError(recover()); e != nil {
			err = e
		}
	}()
	methodNameUrl := fmt.Sprintf("/api/%s/tenants/%s/%s%s", s.ApiVersion, tenantId, s.ResourceName, methodName)
	_, err = daprclient.GetDaprDDDClient().InvokeService(ctx, s.QueryAppId, methodNameUrl, "get", nil, resData)
	if err == nil {
		isFound = true
	}
	return isFound, err
}
