package register

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func init() {
	ddd.RegisterAggregateType(model.AggregateType, model.NewAggregate)
}
