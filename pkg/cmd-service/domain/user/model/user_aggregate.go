package model

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// UserAggregate
// @Description:  <no value> 聚合类型
//
type UserAggregate struct {
	Email     string `json:"email" validate:"-"`           // 电子邮箱
	Id        string `json:"id" validate:"required"`       // 用户Id
	IsDeleted bool   `json:"isDeleted" validate:"-"`       // 已删除
	Name      string `json:"name" validate:"-"`            // 用户名称
	Remarks   string `json:"remarks" validate:"-"`         // 备注
	TenantId  string `json:"tenantId" validate:"required"` // 租户ID
}

const AggregateType = "dapr-ddd-demo.UserAggregate"

// MaskMapper时不复制的属性
var aggMapperRemove []string

func init() {
}

//
// NewUserAggregate
// @Description: 新建<no value> 聚合对象
// @return *UserAggregate
//
func NewUserAggregate() *UserAggregate {
	return &UserAggregate{}
}

//
// NewAggregate
// @Description: 新建当前包中的聚合对象，当前包中只能有一个聚合类型
// @return ddd.Aggregate
//
func NewAggregate() ddd.Aggregate {
	return NewUserAggregate()
}

//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *UserAggregate) GetAggregateVersion() string {
	return "v1.0"
}

//
// GetAggregateType
// @Description: 获取 聚合的类型
// @receiver a
// @return string 聚合的类型
//
func (a *UserAggregate) GetAggregateType() string {
	return AggregateType
}

//
// GetAggregateId
// @Description: 获取 聚合id
// @receiver a
// @return string 聚合id
//
func (a *UserAggregate) GetAggregateId() string {
	return a.Id
}

//
// GetTenantId
// @Description: 租户id
// @receiver a
// @return string 租户id
//
func (a *UserAggregate) GetTenantId() string {
	return a.TenantId
}
