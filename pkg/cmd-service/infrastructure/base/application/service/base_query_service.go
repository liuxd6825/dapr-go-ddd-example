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
	return s.QueryData(ctx, tenantId, "/"+id, resData)
}

func (s *BaseQueryAppService) QueryData(ctx context.Context, tenantId, methodName string, resData interface{}) (isFound bool, err error) {
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
