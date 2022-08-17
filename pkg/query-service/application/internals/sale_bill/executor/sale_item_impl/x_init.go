package sale_item_impl

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/executor"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

func GetSaleItemFindByIdExecutor() executor.SaleItemFindByIdExecutor {
	return singleutils.GetObject[*saleItemFindByIdExecutor]()
}
func GetSaleItemFindBySaleBillIdExecutor() executor.SaleItemFindBySaleBillIdExecutor {
	return singleutils.GetObject[*saleItemFindBySaleBillIdExecutor]()
}

func GetSaleItemFindByIdsExecutor() executor.SaleItemFindByIdsExecutor {
	return singleutils.GetObject[*saleItemFindByIdsExecutor]()
}

func GetSaleItemFindPagingExecutor() executor.SaleItemFindPagingExecutor {
	return singleutils.GetObject[*saleItemFindPagingExecutor]()
}

func GetSaleItemFindAllExecutor() executor.SaleItemFindAllExecutor {
	return singleutils.GetObject[*saleItemFindAllExecutor]()
}

func GetSaleItemCreateExecutor() executor.SaleItemCreateExecutor {
	return singleutils.GetObject[*saleItemCreateExecutor]()
}

func GetSaleItemCreateManyExecutor() executor.SaleItemCreateManyExecutor {
	return singleutils.GetObject[*saleItemCreateManyExecutor]()
}

func GetSaleItemUpdateExecutor() executor.SaleItemUpdateExecutor {
	return singleutils.GetObject[*saleItemUpdateExecutor]()
}

func GetSaleItemUpdateManyExecutor() executor.SaleItemUpdateManyExecutor {
	return singleutils.GetObject[*saleItemUpdateManyExecutor]()
}

func GetSaleItemDeleteByIdExecutor() executor.SaleItemDeleteByIdExecutor {
	return singleutils.GetObject[*saleItemDeleteByIdExecutor]()
}
func GetSaleItemDeleteBySaleBillIdExecutor() executor.SaleItemDeleteBySaleBillIdExecutor {
	return singleutils.GetObject[*saleItemDeleteBySaleBillIdExecutor]()
}

func GetSaleItemDeleteManyExecutor() executor.SaleItemDeleteManyExecutor {
	return singleutils.GetObject[*saleItemDeleteManyExecutor]()
}

func GetSaleItemDeleteAllExecutor() executor.SaleItemDeleteAllExecutor {
	return singleutils.GetObject[*saleItemDeleteAllExecutor]()
}

func init() {
	if err := singleutils.Set[*saleItemFindByIdExecutor](func() *saleItemFindByIdExecutor { return newSaleItemFindByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemFindBySaleBillIdExecutor](func() *saleItemFindBySaleBillIdExecutor { return newSaleItemFindBySaleBillIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemFindByIdsExecutor](func() *saleItemFindByIdsExecutor { return newSaleItemFindByIdsExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemFindPagingExecutor](func() *saleItemFindPagingExecutor { return newSaleItemFindPagingExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemFindAllExecutor](func() *saleItemFindAllExecutor { return newSaleItemFindAllExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemCreateExecutor](func() *saleItemCreateExecutor { return newSaleItemCreateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemCreateManyExecutor](func() *saleItemCreateManyExecutor { return newSaleItemCreateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemUpdateExecutor](func() *saleItemUpdateExecutor { return newSaleItemUpdateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemUpdateManyExecutor](func() *saleItemUpdateManyExecutor { return newSaleItemUpdateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemDeleteByIdExecutor](func() *saleItemDeleteByIdExecutor { return newSaleItemDeleteByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemDeleteBySaleBillIdExecutor](func() *saleItemDeleteBySaleBillIdExecutor { return newSaleItemDeleteBySaleBillIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemDeleteManyExecutor](func() *saleItemDeleteManyExecutor { return newSaleItemDeleteManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemDeleteAllExecutor](func() *saleItemDeleteAllExecutor { return newSaleItemDeleteAllExecutor() }); err != nil {
		panic(err)
	}
}
