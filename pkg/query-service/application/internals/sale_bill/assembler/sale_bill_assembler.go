package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/query"
)

type saleBillAssembler struct {
}

var SaleBill = saleBillAssembler{}

func (a saleBillAssembler) AssFindByIdAppQuery(tenantId, id string) *appquery.SaleBillFindByIdAppQuery {
	res := appquery.NewSaleBillFindByIdAppQuery()
	res.TenantId = tenantId
	res.Id = id
	return res
}

func (a saleBillAssembler) AssFindByIdsAppQuery(tenantId string, ids []string) *appquery.SaleBillFindByIdsAppQuery {
	res := appquery.NewSaleBillFindByIdsAppQuery()
	res.TenantId = tenantId
	res.Ids = ids
	return res
}

func (a saleBillAssembler) AssFindAllAppQuery(tenantId string) *appquery.SaleBillFindAllAppQuery {
	return &appquery.SaleBillFindAllAppQuery{TenantId: tenantId}
}

func (a saleBillAssembler) AssFindPagingResult(fpr *query.SaleBillFindPagingResult) *appquery.SaleBillFindPagingResult {
	res := &appquery.SaleBillFindPagingResult{}
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
