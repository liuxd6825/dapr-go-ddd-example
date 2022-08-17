package register

import (
	inventory_handler "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/handler"
	sale_bill_handler "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/handler"
	user_handler "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterSubscribe
// @Description: 注册领域事件监听器
// @return []restapp.RegisterSubscribe
//
func GetRegisterSubscribe() []restapp.RegisterSubscribe {
	var list []restapp.RegisterSubscribe
	list = append(list, inventory_handler.NewInventorySubscribe())
	list = append(list, sale_bill_handler.NewSaleBillSubscribe())
	list = append(list, sale_bill_handler.NewSaleItemSubscribe())
	list = append(list, user_handler.NewUserSubscribe())

	return list
}
