package register

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/model/user_model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func init() {
	ddd.RegisterAggregateType(user_model.AggregateType, user_model.NewAggregate)
}
