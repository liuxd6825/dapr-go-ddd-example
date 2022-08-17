package utils

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
	"time"
)

type ViewObject interface {
	GetCreatedTime() *time.Time // 获取 创建时间
	SetCreatedTime(*time.Time)  // 设置 创建时间
	GetCreatorId() string       // 获取 创建人ID
	SetCreatorId(string)        // 设置 创建人ID
	GetCreatorName() string     // 获取 创建人名称
	SetCreatorName(string)      // 设置 创建人名称
	GetDeletedTime() *time.Time // 获取 删除时间
	SetDeletedTime(*time.Time)  // 设置 删除时间
	GetDeleterId() string       // 获取 删除人ID
	SetDeleterId(string)        // 设置 删除人ID
	GetDeleterName() string     // 获取 删除人名称
	SetDeleterName(string)      // 设置 删除人名称
	GetId() string              // 获取 主键
	SetId(string)               // 设置 主键
	GetIsDeleted() bool         // 获取 是否删除
	SetIsDeleted(bool)          // 设置 是否删除
	GetRemarks() string         // 获取 备注
	SetRemarks(string)          // 设置 备注
	GetTenantId() string        // 获取 租户ID
	SetTenantId(string)         // 设置 租户ID
	GetUpdatedTime() *time.Time // 获取 修改时间
	SetUpdatedTime(*time.Time)  // 设置 修改时间
	GetUpdaterId() string       // 获取 修改人ID
	SetUpdaterId(string)        // 设置 修改人ID
	GetUpdaterName() string     // 获取 修改人名称
	SetUpdaterName(string)      // 设置 修改人名称
}

type SetViewType int

const (
	SetViewCreated SetViewType = iota // 开始生成枚举值, 默认为0
	SetViewUpdated
	SetViewDeleted
	SetViewOther
)

const StringEmpty = ""

//
// ViewMapper
// @Description: 视图属性自动复制
// @param ctx 上下文
// @param toView 视图对象
// @param fromData Event.Data 事件数据对象
// @return error 错误
//
func ViewMapper(ctx context.Context, toView ViewObject, event ddd.DomainEvent, setType SetViewType) error {
	err := Mapper(event.GetData(), toView)
	if err != nil {
		return err
	}
	err = SetViewDefaultFields(ctx, toView, event.GetCreatedTime(), setType)
	if err != nil {
		return err
	}
	return nil
}

//
// SetViewDefaultFields
// @Description:      通过ctx上下文，设置view视图对象属性， 如从ctx中的Token信息服务
// @param ctx         上下文
// @param viewFields  view视图对象
// @return error      错误
//
func SetViewDefaultFields(ctx context.Context, viewObj ViewObject, setTime time.Time, setType SetViewType) error {
	if viewObj == nil {
		return nil
	}
	userName := "userName"
	userId := "userId"
	eventTime := &setTime
	if eventTime.IsZero() {
		t := time.Now()
		eventTime = &t
	}

	switch setType {
	case SetViewCreated:
		viewObj.SetCreatorName(userName)
		viewObj.SetCreatorId(userId)
		viewObj.SetCreatedTime(eventTime)

		viewObj.SetUpdaterName(userName)
		viewObj.SetUpdaterId(userId)
		viewObj.SetUpdatedTime(eventTime)

		viewObj.SetDeleterName(StringEmpty)
		viewObj.SetDeleterId(StringEmpty)
		viewObj.SetDeletedTime(nil)
		viewObj.SetIsDeleted(false)
		break
	case SetViewUpdated:
		viewObj.SetUpdaterName(userName)
		viewObj.SetUpdaterId(userId)
		viewObj.SetUpdatedTime(eventTime)
		break
	case SetViewDeleted:
		viewObj.SetDeleterName(userName)
		viewObj.SetDeleterId(userId)
		viewObj.SetDeletedTime(eventTime)
		viewObj.SetIsDeleted(true)
		break
	default:
		break
	}
	return nil
}

//
// Mapper
// @Description: 进行struct属性复制，支持深度复制
// @param fromObj 来源
// @param toObj 目标
// @return error
//
func Mapper(fromObj, toObj interface{}) error {
	return types.Mapper(fromObj, toObj)
}

//
// MaskMapper
// @Description: 根据指定进行属性复制，不支持深度复制
// @param fromObj 来源
// @param toObj 目标
// @param mask 要复制属性列表
// @return error
//
func MaskMapper(fromObj, toObj interface{}, mask []string) error {
	return types.MaskMapper(fromObj, toObj, mask)
}
