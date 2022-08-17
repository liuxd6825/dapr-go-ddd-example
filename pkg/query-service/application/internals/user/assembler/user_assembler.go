package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/query"
)

func AssUserFindByIdQuery(tenantId, id string) *appquery.UserFindByIdAppQuery {
	res := appquery.NewUserFindByIdAppQuery()
	res.TenantId = tenantId
	res.Id = id
	return res
}

func AssUserFindByIdsQuery(tenantId string, ids []string) *appquery.UserFindByIdsAppQuery {
	res := appquery.NewUserFindByIdsAppQuery()
	res.TenantId = tenantId
	res.Ids = ids
	return res
}

func AssUserFindAllQuery(tenantId string) *appquery.UserFindAllAppQuery {
	return &appquery.UserFindAllAppQuery{TenantId: tenantId}
}

func AssUserFindPagingResult(fpr *query.UserFindPagingResult) *appquery.UserFindPagingResult {
	res := &appquery.UserFindPagingResult{}
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
