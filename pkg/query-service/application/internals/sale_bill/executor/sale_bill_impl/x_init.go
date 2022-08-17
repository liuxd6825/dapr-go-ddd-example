package sale_bill_impl

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/executor"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

func GetSaleBillFindByIdExecutor() executor.SaleBillFindByIdExecutor {
	return singleutils.GetObject[*saleBillFindByIdExecutor]()
}

func GetSaleBillFindByIdsExecutor() executor.SaleBillFindByIdsExecutor {
	return singleutils.GetObject[*saleBillFindByIdsExecutor]()
}

func GetSaleBillFindPagingExecutor() executor.SaleBillFindPagingExecutor {
	return singleutils.GetObject[*saleBillFindPagingExecutor]()
}

func GetSaleBillFindAllExecutor() executor.SaleBillFindAllExecutor {
	return singleutils.GetObject[*saleBillFindAllExecutor]()
}

func GetSaleBillCreateExecutor() executor.SaleBillCreateExecutor {
	return singleutils.GetObject[*saleBillCreateExecutor]()
}

func GetSaleBillCreateManyExecutor() executor.SaleBillCreateManyExecutor {
	return singleutils.GetObject[*saleBillCreateManyExecutor]()
}

func GetSaleBillUpdateExecutor() executor.SaleBillUpdateExecutor {
	return singleutils.GetObject[*saleBillUpdateExecutor]()
}

func GetSaleBillUpdateManyExecutor() executor.SaleBillUpdateManyExecutor {
	return singleutils.GetObject[*saleBillUpdateManyExecutor]()
}

func GetSaleBillDeleteByIdExecutor() executor.SaleBillDeleteByIdExecutor {
	return singleutils.GetObject[*saleBillDeleteByIdExecutor]()
}

func GetSaleBillDeleteManyExecutor() executor.SaleBillDeleteManyExecutor {
	return singleutils.GetObject[*saleBillDeleteManyExecutor]()
}

func GetSaleBillDeleteAllExecutor() executor.SaleBillDeleteAllExecutor {
	return singleutils.GetObject[*saleBillDeleteAllExecutor]()
}

func init() {
	if err := singleutils.Set[*saleBillFindByIdExecutor](func() *saleBillFindByIdExecutor { return newSaleBillFindByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillFindByIdsExecutor](func() *saleBillFindByIdsExecutor { return newSaleBillFindByIdsExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillFindPagingExecutor](func() *saleBillFindPagingExecutor { return newSaleBillFindPagingExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillFindAllExecutor](func() *saleBillFindAllExecutor { return newSaleBillFindAllExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillCreateExecutor](func() *saleBillCreateExecutor { return newSaleBillCreateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillCreateManyExecutor](func() *saleBillCreateManyExecutor { return newSaleBillCreateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillUpdateExecutor](func() *saleBillUpdateExecutor { return newSaleBillUpdateExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillUpdateManyExecutor](func() *saleBillUpdateManyExecutor { return newSaleBillUpdateManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillDeleteByIdExecutor](func() *saleBillDeleteByIdExecutor { return newSaleBillDeleteByIdExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillDeleteManyExecutor](func() *saleBillDeleteManyExecutor { return newSaleBillDeleteManyExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillDeleteAllExecutor](func() *saleBillDeleteAllExecutor { return newSaleBillDeleteAllExecutor() }); err != nil {
		panic(err)
	}
}
