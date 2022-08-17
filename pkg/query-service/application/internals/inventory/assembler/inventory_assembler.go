package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/query"
)

func AssInventoryFindByIdQuery(tenantId, id string) *appquery.InventoryFindByIdAppQuery {
	res := appquery.NewInventoryFindByIdAppQuery()
	res.TenantId = tenantId
	res.Id = id
	return res
}

func AssInventoryFindByIdsQuery(tenantId string, ids []string) *appquery.InventoryFindByIdsAppQuery {
	res := appquery.NewInventoryFindByIdsAppQuery()
	res.TenantId = tenantId
	res.Ids = ids
	return res
}

func AssInventoryFindAllQuery(tenantId string) *appquery.InventoryFindAllAppQuery {
	return &appquery.InventoryFindAllAppQuery{TenantId: tenantId}
}

func AssInventoryFindPagingResult(fpr *query.InventoryFindPagingResult) *appquery.InventoryFindPagingResult {
	res := &appquery.InventoryFindPagingResult{}
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
