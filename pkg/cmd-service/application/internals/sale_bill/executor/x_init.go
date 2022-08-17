package executor

import "github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"

func GetSaleBillConfirmCommandExecutor() SaleBillConfirmCommandExecutor {
	return singleutils.GetObject[*saleBillConfirmCommandExecutor]()
}

func GetSaleBillCreateCommandExecutor() SaleBillCreateCommandExecutor {
	return singleutils.GetObject[*saleBillCreateCommandExecutor]()
}

func GetSaleBillDeleteCommandExecutor() SaleBillDeleteCommandExecutor {
	return singleutils.GetObject[*saleBillDeleteCommandExecutor]()
}

func GetSaleBillUpdateCommandExecutor() SaleBillUpdateCommandExecutor {
	return singleutils.GetObject[*saleBillUpdateCommandExecutor]()
}

func GetSaleItemCreateCommandExecutor() SaleItemCreateCommandExecutor {
	return singleutils.GetObject[*saleItemCreateCommandExecutor]()
}

func GetSaleItemDeleteCommandExecutor() SaleItemDeleteCommandExecutor {
	return singleutils.GetObject[*saleItemDeleteCommandExecutor]()
}

func GetSaleItemUpdateCommandExecutor() SaleItemUpdateCommandExecutor {
	return singleutils.GetObject[*saleItemUpdateCommandExecutor]()
}

func GetFindAggregateByIdExecutor() FindAggregateByIdExecutor {
	return singleutils.GetObject[*findAggregateByIdExecutor]()
}

func init() {
	if err := singleutils.Set[*saleBillConfirmCommandExecutor](func() *saleBillConfirmCommandExecutor { return newSaleBillConfirmCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillCreateCommandExecutor](func() *saleBillCreateCommandExecutor { return newSaleBillCreateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillDeleteCommandExecutor](func() *saleBillDeleteCommandExecutor { return newSaleBillDeleteCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleBillUpdateCommandExecutor](func() *saleBillUpdateCommandExecutor { return newSaleBillUpdateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemCreateCommandExecutor](func() *saleItemCreateCommandExecutor { return newSaleItemCreateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemDeleteCommandExecutor](func() *saleItemDeleteCommandExecutor { return newSaleItemDeleteCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*saleItemUpdateCommandExecutor](func() *saleItemUpdateCommandExecutor { return newSaleItemUpdateCommandExecutor() }); err != nil {
		panic(err)
	}
	if err := singleutils.Set[*findAggregateByIdExecutor](func() *findAggregateByIdExecutor { return newFindAggregateByIdExecutor() }); err != nil {
		panic(err)
	}
}
