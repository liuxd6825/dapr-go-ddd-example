package register

import (
	inventory_event "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/event"
	sale_bill_event "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	user_event "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterEventType
// @Description: 获取所有注册事件类型
// @return *[]restapp.RegisterEventType
//
func GetRegisterEventType() []restapp.RegisterEventType {
	var list []restapp.RegisterEventType
	list = append(list, inventory_event.GetRegisterEventTypes()...)
	list = append(list, sale_bill_event.GetRegisterEventTypes()...)
	list = append(list, user_event.GetRegisterEventTypes()...)
	return list
}
