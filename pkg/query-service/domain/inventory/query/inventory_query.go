package query

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
)

type InventoryFindByIdQuery struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

func NewInventoryFindByIdQuery(tenantId, id string) *InventoryFindByIdQuery {
	return &InventoryFindByIdQuery{
		TenantId: tenantId,
		Id:       id,
	}
}

type InventoryFindByIdsQuery struct {
	TenantId string   `json:"tenantId"`
	Ids      []string `json:"ids"`
}

func NewInventoryFindByIdsQuery(tenantId string, ids []string) *InventoryFindByIdsQuery {
	return &InventoryFindByIdsQuery{
		TenantId: tenantId,
		Ids:      ids,
	}
}

type InventoryFindAllQuery struct {
	TenantId string `json:"tenantId"`
}

func NewInventoryFindAllQuery(tenantId string) *InventoryFindAllQuery {
	return &InventoryFindAllQuery{
		TenantId: tenantId,
	}
}

type InventoryFindPagingQuery struct {
	TenantId string `json:"tenantId"`
	Fields   string `json:"fields"`
	Filter   string `json:"filter"`
	Sort     string `json:"sort"`
	PageNum  int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

func NewInventoryFindPagingQuery(tenantId string, fields string, filter string, sort string, pageNum int64, pageSize int64) *InventoryFindPagingQuery {
	return &InventoryFindPagingQuery{
		TenantId: tenantId,
		Fields:   fields,
		Filter:   filter,
		Sort:     sort,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func (q *InventoryFindPagingQuery) GetTenantId() string {
	return q.TenantId
}

func (q *InventoryFindPagingQuery) GetFields() string {
	return q.Fields
}

func (q *InventoryFindPagingQuery) GetFilter() string {
	return q.Filter
}

func (q *InventoryFindPagingQuery) GetSort() string {
	return q.Sort
}

func (q *InventoryFindPagingQuery) GetPageNum() int64 {
	return q.PageNum
}

func (q *InventoryFindPagingQuery) GetPageSize() int64 {
	return q.PageSize
}

type InventoryFindPagingResult struct {
	Data       []*view.InventoryView `json:"data"`
	TotalRows  *int64                `json:"totalRows"`
	TotalPages *int64                `json:"totalPages"`
	PageNum    int64                 `json:"pageNum"`
	PageSize   int64                 `json:"pageSize"`
	Filter     string                `json:"filter"`
	Sort       string                `json:"sort"`
	Error      error                 `json:"-"`
	IsFound    bool                  `json:"-"`
}

func NewInventoryFindPagingResult() *InventoryFindPagingResult {
	return &InventoryFindPagingResult{}
}
