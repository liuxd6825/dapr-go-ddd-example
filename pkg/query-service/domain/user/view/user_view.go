package view

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/domain/view"
)

//
// UserView
// @Description: 用户
//
type UserView struct {
	base.BaseView `bson:",inline"`
	Email         string `json:"email,omitempty" bson:"email"  gorm:""` // 电子邮箱
	Name          string `json:"name,omitempty" bson:"name"  gorm:""`   // 用户名称
}

//
// NewUserView
// @Description: 用户
//
func NewUserView() *UserView {
	return &UserView{}
}
