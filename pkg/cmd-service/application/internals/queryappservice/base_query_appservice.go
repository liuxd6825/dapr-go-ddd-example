package queryappservice

import (
	"context"
	"fmt"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_errors"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type GetOptions interface {
	SetResourceName(resourceName string)
	SetApiVersion(apiVersion string)
	GetResourceName() *string
	GetApiVersion() *string
	Merge(opts ...GetOptions)
}

type getOptions struct {
	resourceName *string
	apiVersion   *string
}

func (o *getOptions) SetResourceName(resourceName string) {
	o.resourceName = &resourceName
}

func (o *getOptions) SetApiVersion(apiVersion string) {
	o.apiVersion = &apiVersion
}

func (o *getOptions) GetResourceName() *string {
	return o.resourceName
}

func (o *getOptions) GetApiVersion() *string {
	return o.apiVersion
}

func (o *getOptions) Merge(opts ...GetOptions) {
	if opts == nil {
		return
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.GetResourceName() != nil {
			o.SetResourceName(*opt.GetResourceName())
		}
		if opt.GetApiVersion() != nil {
			o.SetApiVersion(*opt.GetResourceName())
		}
	}
}

type QueryAppService[T any] interface {
	AppId() string
	ResourceName() string
	ApiVersion() string
}

type BaseQueryAppService[T any] struct {
	appId        string
	resourceName string
	apiVersion   string
}

func (s *BaseQueryAppService[T]) GetById(ctx context.Context, tenantId, id string, resData interface{}, opts ...GetOptions) (isFound bool, err error) {
	return s.GetData(ctx, tenantId, id, resData)
}

func (s *BaseQueryAppService[T]) GetData(ctx context.Context, tenantId, methodName string, resData interface{}, opts ...GetOptions) (isFound bool, err error) {
	defer func() {
		if e := ddd_errors.GetRecoverError(recover()); e != nil {
			err = e
		}
	}()
	options := s.newOptions(opts...)
	apiVersion := *options.GetApiVersion()
	resourceName := *options.GetResourceName()
	methodNameUrl := fmt.Sprintf("/api/%s/tenants/%s/%s/%s", apiVersion, tenantId, resourceName, methodName)
	_, err = restapp.GetDaprClient().InvokeService(ctx, s.appId, methodNameUrl, "get", nil, resData)
	if err == nil {
		isFound = true
	}
	return isFound, err
}

func (s *BaseQueryAppService[T]) newOptions(opts ...GetOptions) GetOptions {
	options := &getOptions{
		resourceName: &s.resourceName,
		apiVersion:   &s.apiVersion,
	}
	options.Merge(opts...)
	return options
}

func (s *BaseQueryAppService[T]) AppId() string {
	return s.appId
}

func (s *BaseQueryAppService[T]) ResourceName() string {
	return s.resourceName
}

func (s *BaseQueryAppService[T]) ApiVersion() string {
	return s.apiVersion
}
