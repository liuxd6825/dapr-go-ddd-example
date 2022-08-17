package executor

import "github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"

func GetInventoryCreateCommandExecutor() InventoryCreateCommandExecutor {
	return singleutils.GetObject[*inventoryCreateCommandExecutor]()
}

func GetInventoryUpdateCommandExecutor() InventoryUpdateCommandExecutor {
	return singleutils.GetObject[*inventoryUpdateCommandExecutor]()
}

func GetFindAggregateByIdExecutor() FindAggregateByIdExecutor {
	return singleutils.GetObject[*findAggregateByIdExecutor]()
}

func init() {
	if err := singleutils.Set[*inventoryCreateCommandExecutor](func() *inventoryCreateCommandExecutor { return newInventoryCreateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryUpdateCommandExecutor](func() *inventoryUpdateCommandExecutor { return newInventoryUpdateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*findAggregateByIdExecutor](func() *findAggregateByIdExecutor { return newFindAggregateByIdExecutor() }); err != nil {
		panic(err)
	}
}
