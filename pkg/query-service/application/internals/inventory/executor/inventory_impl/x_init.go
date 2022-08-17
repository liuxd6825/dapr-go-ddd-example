package inventory_impl

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/executor"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

func GetInventoryFindByIdExecutor() executor.InventoryFindByIdExecutor {
	return singleutils.GetObject[*inventoryFindByIdExecutor]()
}

func GetInventoryFindByIdsExecutor() executor.InventoryFindByIdsExecutor {
	return singleutils.GetObject[*inventoryFindByIdsExecutor]()
}

func GetInventoryFindPagingExecutor() executor.InventoryFindPagingExecutor {
	return singleutils.GetObject[*inventoryFindPagingExecutor]()
}

func GetInventoryFindAllExecutor() executor.InventoryFindAllExecutor {
	return singleutils.GetObject[*inventoryFindAllExecutor]()
}

func GetInventoryCreateExecutor() executor.InventoryCreateExecutor {
	return singleutils.GetObject[*inventoryCreateExecutor]()
}

func GetInventoryCreateManyExecutor() executor.InventoryCreateManyExecutor {
	return singleutils.GetObject[*inventoryCreateManyExecutor]()
}

func GetInventoryUpdateExecutor() executor.InventoryUpdateExecutor {
	return singleutils.GetObject[*inventoryUpdateExecutor]()
}

func GetInventoryUpdateManyExecutor() executor.InventoryUpdateManyExecutor {
	return singleutils.GetObject[*inventoryUpdateManyExecutor]()
}

func GetInventoryDeleteByIdExecutor() executor.InventoryDeleteByIdExecutor {
	return singleutils.GetObject[*inventoryDeleteByIdExecutor]()
}

func GetInventoryDeleteManyExecutor() executor.InventoryDeleteManyExecutor {
	return singleutils.GetObject[*inventoryDeleteManyExecutor]()
}

func GetInventoryDeleteAllExecutor() executor.InventoryDeleteAllExecutor {
	return singleutils.GetObject[*inventoryDeleteAllExecutor]()
}

func init() {
	if err := singleutils.Set[*inventoryFindByIdExecutor](func() *inventoryFindByIdExecutor { return newInventoryFindByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryFindByIdsExecutor](func() *inventoryFindByIdsExecutor { return newInventoryFindByIdsExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryFindPagingExecutor](func() *inventoryFindPagingExecutor { return newInventoryFindPagingExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryFindAllExecutor](func() *inventoryFindAllExecutor { return newInventoryFindAllExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryCreateExecutor](func() *inventoryCreateExecutor { return newInventoryCreateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryCreateManyExecutor](func() *inventoryCreateManyExecutor { return newInventoryCreateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryUpdateExecutor](func() *inventoryUpdateExecutor { return newInventoryUpdateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryUpdateManyExecutor](func() *inventoryUpdateManyExecutor { return newInventoryUpdateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryDeleteByIdExecutor](func() *inventoryDeleteByIdExecutor { return newInventoryDeleteByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryDeleteManyExecutor](func() *inventoryDeleteManyExecutor { return newInventoryDeleteManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*inventoryDeleteAllExecutor](func() *inventoryDeleteAllExecutor { return newInventoryDeleteAllExecutor() }); err != nil {
		panic(err)
	}
}
