package register

import (
	inventory_model "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/model"
	sale_bill_model "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/model"
	user_model "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func init() {
	ddd.RegisterAggregateType(inventory_model.AggregateType, inventory_model.NewAggregate)
	ddd.RegisterAggregateType(sale_bill_model.AggregateType, sale_bill_model.NewAggregate)
	ddd.RegisterAggregateType(user_model.AggregateType, user_model.NewAggregate)
}
