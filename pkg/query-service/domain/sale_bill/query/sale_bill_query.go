package query

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
)

type SaleBillFindByIdQuery struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

func NewSaleBillFindByIdQuery(tenantId, id string) *SaleBillFindByIdQuery {
	return &SaleBillFindByIdQuery{
		TenantId: tenantId,
		Id:       id,
	}
}

type SaleBillFindByIdsQuery struct {
	TenantId string   `json:"tenantId"`
	Ids      []string `json:"ids"`
}

func NewSaleBillFindByIdsQuery(tenantId string, ids []string) *SaleBillFindByIdsQuery {
	return &SaleBillFindByIdsQuery{
		TenantId: tenantId,
		Ids:      ids,
	}
}

type SaleBillFindAllQuery struct {
	TenantId string `json:"tenantId"`
}

func NewSaleBillFindAllQuery(tenantId string) *SaleBillFindAllQuery {
	return &SaleBillFindAllQuery{
		TenantId: tenantId,
	}
}

type SaleBillFindPagingQuery struct {
	TenantId string `json:"tenantId"`
	Fields   string `json:"fields"`
	Filter   string `json:"filter"`
	Sort     string `json:"sort"`
	PageNum  int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

func NewSaleBillFindPagingQuery(tenantId string, fields string, filter string, sort string, pageNum int64, pageSize int64) *SaleBillFindPagingQuery {
	return &SaleBillFindPagingQuery{
		TenantId: tenantId,
		Fields:   fields,
		Filter:   filter,
		Sort:     sort,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func (q *SaleBillFindPagingQuery) GetTenantId() string {
	return q.TenantId
}

func (q *SaleBillFindPagingQuery) GetFields() string {
	return q.Fields
}

func (q *SaleBillFindPagingQuery) GetFilter() string {
	return q.Filter
}

func (q *SaleBillFindPagingQuery) GetSort() string {
	return q.Sort
}

func (q *SaleBillFindPagingQuery) GetPageNum() int64 {
	return q.PageNum
}

func (q *SaleBillFindPagingQuery) GetPageSize() int64 {
	return q.PageSize
}

type SaleBillFindPagingResult struct {
	Data       []*view.SaleBillView `json:"data"`
	TotalRows  *int64               `json:"totalRows"`
	TotalPages *int64               `json:"totalPages"`
	PageNum    int64                `json:"pageNum"`
	PageSize   int64                `json:"pageSize"`
	Filter     string               `json:"filter"`
	Sort       string               `json:"sort"`
	Error      error                `json:"-"`
	IsFound    bool                 `json:"-"`
}

func NewSaleBillFindPagingResult() *SaleBillFindPagingResult {
	return &SaleBillFindPagingResult{}
}
