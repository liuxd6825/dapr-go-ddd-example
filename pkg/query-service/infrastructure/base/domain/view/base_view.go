package view

import (
	"time"
)

//
// BaseView
// @Description: 视图基类
//
type BaseView struct {
	CreatedTime *time.Time `json:"createdTime,omitempty"  bson:"created_time" gorm:"" `                 // 创建时间
	CreatorId   string     `json:"creatorId,omitempty"  bson:"creator_id" gorm:"index:idx_creator_id" ` // 创建人ID
	CreatorName string     `json:"creatorName,omitempty"  bson:"creator_name" gorm:"" `                 // 创建人名称
	DeletedTime *time.Time `json:"deletedTime,omitempty"  bson:"deleted_time" gorm:"" `                 // 删除时间
	DeleterId   string     `json:"deleterId,omitempty"  bson:"deleter_id" gorm:"index:idx_deleter_id" ` // 删除人ID
	DeleterName string     `json:"deleterName,omitempty"  bson:"deleter_name" gorm:"" `                 // 删除人名称
	Id          string     `json:"id,omitempty"  bson:"_id" gorm:"primaryKey" `                         // 主键
	IsDeleted   bool       `json:"isDeleted,omitempty"  bson:"is_deleted" gorm:"" `                     // 是否删除
	Remarks     string     `json:"remarks,omitempty"  bson:"remarks" gorm:"" `                          // 备注
	TenantId    string     `json:"tenantId,omitempty"  bson:"tenant_id" gorm:"index:idx_tenant_id" `    // 租户ID
	UpdatedTime *time.Time `json:"updatedTime,omitempty"  bson:"updated_time" gorm:"" `                 // 修改时间
	UpdaterId   string     `json:"updaterId,omitempty"  bson:"updater_id" gorm:"index:idx_updater_id" ` // 修改人ID
	UpdaterName string     `json:"updaterName,omitempty"  bson:"updater_name" gorm:"" `                 // 修改人名称
}

//
// GetCreatedTime
// @Description: 获取 创建时间
//
func (v *BaseView) GetCreatedTime() *time.Time {
	return v.CreatedTime
}

//
// SetCreatedTime
// @Description: 设置 创建时间
//
func (v *BaseView) SetCreatedTime(value *time.Time) {
	v.CreatedTime = value
}

//
// GetCreatorId
// @Description: 获取 创建人ID
//
func (v *BaseView) GetCreatorId() string {
	return v.CreatorId
}

//
// SetCreatorId
// @Description: 设置 创建人ID
//
func (v *BaseView) SetCreatorId(value string) {
	v.CreatorId = value
}

//
// GetCreatorName
// @Description: 获取 创建人名称
//
func (v *BaseView) GetCreatorName() string {
	return v.CreatorName
}

//
// SetCreatorName
// @Description: 设置 创建人名称
//
func (v *BaseView) SetCreatorName(value string) {
	v.CreatorName = value
}

//
// GetDeletedTime
// @Description: 获取 删除时间
//
func (v *BaseView) GetDeletedTime() *time.Time {
	return v.DeletedTime
}

//
// SetDeletedTime
// @Description: 设置 删除时间
//
func (v *BaseView) SetDeletedTime(value *time.Time) {
	v.DeletedTime = value
}

//
// GetDeleterId
// @Description: 获取 删除人ID
//
func (v *BaseView) GetDeleterId() string {
	return v.DeleterId
}

//
// SetDeleterId
// @Description: 设置 删除人ID
//
func (v *BaseView) SetDeleterId(value string) {
	v.DeleterId = value
}

//
// GetDeleterName
// @Description: 获取 删除人名称
//
func (v *BaseView) GetDeleterName() string {
	return v.DeleterName
}

//
// SetDeleterName
// @Description: 设置 删除人名称
//
func (v *BaseView) SetDeleterName(value string) {
	v.DeleterName = value
}

//
// GetId
// @Description: 获取 主键
//
func (v *BaseView) GetId() string {
	return v.Id
}

//
// SetId
// @Description: 设置 主键
//
func (v *BaseView) SetId(value string) {
	v.Id = value
}

//
// GetIsDeleted
// @Description: 获取 是否删除
//
func (v *BaseView) GetIsDeleted() bool {
	return v.IsDeleted
}

//
// SetIsDeleted
// @Description: 设置 是否删除
//
func (v *BaseView) SetIsDeleted(value bool) {
	v.IsDeleted = value
}

//
// GetRemarks
// @Description: 获取 备注
//
func (v *BaseView) GetRemarks() string {
	return v.Remarks
}

//
// SetRemarks
// @Description: 设置 备注
//
func (v *BaseView) SetRemarks(value string) {
	v.Remarks = value
}

//
// GetTenantId
// @Description: 获取 租户ID
//
func (v *BaseView) GetTenantId() string {
	return v.TenantId
}

//
// SetTenantId
// @Description: 设置 租户ID
//
func (v *BaseView) SetTenantId(value string) {
	v.TenantId = value
}

//
// GetUpdatedTime
// @Description: 获取 修改时间
//
func (v *BaseView) GetUpdatedTime() *time.Time {
	return v.UpdatedTime
}

//
// SetUpdatedTime
// @Description: 设置 修改时间
//
func (v *BaseView) SetUpdatedTime(value *time.Time) {
	v.UpdatedTime = value
}

//
// GetUpdaterId
// @Description: 获取 修改人ID
//
func (v *BaseView) GetUpdaterId() string {
	return v.UpdaterId
}

//
// SetUpdaterId
// @Description: 设置 修改人ID
//
func (v *BaseView) SetUpdaterId(value string) {
	v.UpdaterId = value
}

//
// GetUpdaterName
// @Description: 获取 修改人名称
//
func (v *BaseView) GetUpdaterName() string {
	return v.UpdaterName
}

//
// SetUpdaterName
// @Description: 设置 修改人名称
//
func (v *BaseView) SetUpdaterName(value string) {
	v.UpdaterName = value
}
