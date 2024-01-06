package client

import (
	"context"
	"errors"
)
import "github.com/liuxd6825/dapr-go-ddd-sdk/feign"

type FindAggregateByIdConfig struct {
	feign.RequestConfig
	TenantId string  `param:"tenantId"`
	Id       string  `param:"id"`
	Data     UserDTO `body:"application/json"`
}

type UserDTO struct {
	Email     string `json:"email" validate:"-"`           // 电子邮箱
	Id        string `json:"id" validate:"required"`       // 用户Id
	IsDeleted bool   `json:"isDeleted" validate:"-"`       // 已删除
	Name      string `json:"name" validate:"-"`            // 用户名称
	Remarks   string `json:"remarks" validate:"-"`         // 备注
	TenantId  string `json:"tenantId" validate:"required"` // 租户ID
}

type UserCmdFeign struct {
	FindAggregateById func(ctx context.Context, cfg *FindAggregateByIdConfig) (*UserDTO, error) `method:"GET" url:"api/v1.0/tenants/:tenantId/users/aggregate/:id" `
}

func NewUserCmdFeign() (*UserCmdFeign, error) {
	f := feign.New()
	ci, err := f.Build(&UserCmdFeign{},
		feign.WithAppIdURL("example-cmd-service"),
	)
	if err != nil {
		return nil, err
	}
	c, ok := ci.(*UserCmdFeign)
	if !ok {
		return nil, errors.New("is not UserClient")
	}
	return c, nil
}
