package model

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"time"
)

//
// SaleBillAggregate
// @Description:  <no value> 聚合类型
//
type SaleBillAggregate struct {
	Id         string               `json:"id" validate:"required"`            // 主键
	IsDeleted  bool                 `json:"isDeleted" validate:"-"`            // 已删除
	Remarks    string               `json:"remarks" validate:"-"`              // 备注
	SaleItems  *SaleItemItems       `json:"saleItems" copier:"-" validate:"-"` //
	SaleMoney  float64              `json:"saleMoney" validate:"-"`            // 销售金额
	SaleTime   time.Time            `json:"saleTime" validate:"-"`             // 文件大小
	Statue     field.SaleBillStatue `json:"statue" validate:"-"`               // 单据状态
	TenantId   string               `json:"tenantId" validate:"required"`      // 租户ID
	UserId     string               `json:"userId" validate:"required"`        // 用户Id
	UserName   string               `json:"userName" validate:"-"`             // 用户名称
	TotalMoney float64              `json:"totalMoney" validate:"-"`
}

const AggregateType = "dapr-ddd-demo.SaleBillAggregate"

// MaskMapper时不复制的属性
var aggMapperRemove []string

func init() {
	aggMapperRemove = append(aggMapperRemove, "SaleItems")
}

//
// NewSaleBillAggregate
// @Description: 新建<no value> 聚合对象
// @return *SaleBillAggregate
//
func NewSaleBillAggregate() *SaleBillAggregate {
	return &SaleBillAggregate{
		SaleItems: NewSaleItemItems(),
	}
}

//
// NewAggregate
// @Description: 新建当前包中的聚合对象，当前包中只能有一个聚合类型
// @return ddd.Aggregate
//
func NewAggregate() ddd.Aggregate {
	return NewSaleBillAggregate()
}

//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *SaleBillAggregate) GetAggregateVersion() string {
	return "v1.0"
}

//
// GetAggregateType
// @Description: 获取 聚合的类型
// @receiver a
// @return string 聚合的类型
//
func (a *SaleBillAggregate) GetAggregateType() string {
	return AggregateType
}

//
// GetAggregateId
// @Description: 获取 聚合id
// @receiver a
// @return string 聚合id
//
func (a *SaleBillAggregate) GetAggregateId() string {
	return a.Id
}

//
// GetTenantId
// @Description: 租户id
// @receiver a
// @return string 租户id
//
func (a *SaleBillAggregate) GetTenantId() string {
	return a.TenantId
}
