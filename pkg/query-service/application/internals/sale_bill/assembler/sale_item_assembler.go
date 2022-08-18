package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/query"
)

type saleItemAssembler struct {
}

var SaleItem = saleItemAssembler{}

func (a saleItemAssembler) AssFindByIdAppQuery(tenantId, id string) *appquery.SaleItemFindByIdAppQuery {
	res := appquery.NewSaleItemFindByIdAppQuery()
	res.TenantId = tenantId
	res.Id = id
	return res
}

func (a saleItemAssembler) AssFindByIdsAppQuery(tenantId string, ids []string) *appquery.SaleItemFindByIdsAppQuery {
	res := appquery.NewSaleItemFindByIdsAppQuery()
	res.TenantId = tenantId
	res.Ids = ids
	return res
}
func (a saleItemAssembler) AssFindBySaleBillIdAppQuery(tenantId string, saleBillId string) *appquery.SaleItemFindBySaleBillIdAppQuery {
	res := appquery.NewSaleItemFindBySaleBillIdAppQuery()
	res.TenantId = tenantId
	res.SaleBillId = saleBillId
	return res
}

func (a saleItemAssembler) AssFindAllAppQuery(tenantId string) *appquery.SaleItemFindAllAppQuery {
	return &appquery.SaleItemFindAllAppQuery{TenantId: tenantId}
}

func (a saleItemAssembler) AssFindPagingResult(fpr *query.SaleItemFindPagingResult) *appquery.SaleItemFindPagingResult {
	res := &appquery.SaleItemFindPagingResult{}
	res.Sort = fpr.Sort
	res.PageNum = fpr.PageNum
	res.PageSize = fpr.PageSize
	res.Filter = fpr.Filter
	res.Sort = fpr.Sort
	res.Data = fpr.Data
	res.Error = fpr.Error
	res.IsFound = fpr.IsFound
	res.TotalPages = fpr.TotalPages
	res.TotalRows = fpr.TotalRows
	return res
}
