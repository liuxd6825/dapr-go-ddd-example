package model

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// InventoryAggregate
// @Description:  <no value> 聚合类型
//
type InventoryAggregate struct {
	Brand     string `json:"brand" validate:"-"`           // 品牌
	Id        string `json:"id" validate:"required"`       // Id
	IsDeleted bool   `json:"isDeleted" validate:"-"`       // 已删除
	Keywords  string `json:"keywords" validate:"-"`        // 搜索关键字
	Name      string `json:"name" validate:"-"`            // 名称
	Remarks   string `json:"remarks" validate:"-"`         // 备注
	Spec      string `json:"spec" validate:"-"`            // 规格
	TenantId  string `json:"tenantId" validate:"required"` // 租户ID
}

const AggregateType = "dapr-ddd-demo.InventoryAggregate"

// MaskMapper时不复制的属性
var aggMapperRemove []string

func init() {
}

//
// NewInventoryAggregate
// @Description: 新建<no value> 聚合对象
// @return *InventoryAggregate
//
func NewInventoryAggregate() *InventoryAggregate {
	return &InventoryAggregate{}
}

//
// NewAggregate
// @Description: 新建当前包中的聚合对象，当前包中只能有一个聚合类型
// @return ddd.Aggregate
//
func NewAggregate() ddd.Aggregate {
	return NewInventoryAggregate()
}

//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *InventoryAggregate) GetAggregateVersion() string {
	return "v1.0"
}

//
// GetAggregateType
// @Description: 获取 聚合的类型
// @receiver a
// @return string 聚合的类型
//
func (a *InventoryAggregate) GetAggregateType() string {
	return AggregateType
}

//
// GetAggregateId
// @Description: 获取 聚合id
// @receiver a
// @return string 聚合id
//
func (a *InventoryAggregate) GetAggregateId() string {
	return a.Id
}

//
// GetTenantId
// @Description: 租户id
// @receiver a
// @return string 租户id
//
func (a *InventoryAggregate) GetTenantId() string {
	return a.TenantId
}
