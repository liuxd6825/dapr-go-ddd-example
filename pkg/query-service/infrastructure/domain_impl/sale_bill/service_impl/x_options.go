package service_impl

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"time"
)

type options struct {
	timeout    *time.Duration
	updateMask *[]string
}

func NewOptions() service.Options {
	return &options{}
}

func MergeOptions(opts ...service.Options) service.Options {
	o := NewOptions().Merge(opts...)
	return o
}

func (o *options) SetUpdateMask(updateMask *[]string) service.Options {
	o.updateMask = updateMask
	return o
}

func (o *options) GetUpdateMask() *[]string {
	return o.updateMask
}

func (o *options) SetTimeout(timeout *time.Duration) service.Options {
	o.timeout = timeout
	return o
}

func (o *options) GetTimeout() *time.Duration {
	return o.timeout
}

func (o *options) Merge(opts ...service.Options) service.Options {
	for _, item := range opts {
		if item == nil {
			continue
		}
		if item.GetTimeout() != nil {
			o.SetTimeout(item.GetTimeout())
		}
		if item.GetUpdateMask() != nil {
			o.SetUpdateMask(item.GetUpdateMask())
		}
	}
	return o
}
