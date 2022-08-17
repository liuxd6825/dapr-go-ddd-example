package register

import (
	inventory_api "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/inventory/facade"
	sale_bill_api "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/sale_bill/facade"
	user_api "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/user/facade"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// GetRegisterController
// @Description: 注册的控制器
// @return []restapp.Controller
//
func GetRegisterController() []restapp.Controller {
	var list []restapp.Controller
	list = append(list, inventory_api.NewInventoryQueryApi())
	list = append(list, sale_bill_api.NewSaleBillQueryApi())
	list = append(list, sale_bill_api.NewSaleItemQueryApi())
	list = append(list, user_api.NewUserQueryApi())

	return list
}
