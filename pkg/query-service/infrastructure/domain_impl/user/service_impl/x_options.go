package service_impl

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
)

type Options struct {
	timeout    *int
	updateMask *[]string
}

func NewOptions() *Options {
	return &Options{}
}

func MergeOptions(opts ...service.Options) *Options {
	opt := &Options{}
	for _, o := range opts {
		if o == nil {
			continue
		}
		if o.UpdateMask() == nil {
			opt.updateMask = o.UpdateMask()
		}
		if o.Timeout() != nil {
			opt.timeout = o.Timeout()
		}
	}
	return opt
}

func (o *Options) SetUpdateMask(updateMask []string) *Options {
	o.updateMask = &updateMask
	return o
}

func (o *Options) UpdateMask() *[]string {
	return o.updateMask
}

func (o *Options) SetTimeout(timeout int) *Options {
	o.timeout = &timeout
	return o
}

func (o *Options) Timeout() *int {
	return o.timeout
}
