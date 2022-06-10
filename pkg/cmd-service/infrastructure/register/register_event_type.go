package register

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventTypes
// @Description: 获取所有注册事件类型
// @return *[]restapp.RegisterEventType
//
func GetRegisterEventTypes() *[]restapp.RegisterEventType {
	var list []restapp.RegisterEventType
	list = append(list, event.GetRegisterEventTypes()...)
	return &list
}
