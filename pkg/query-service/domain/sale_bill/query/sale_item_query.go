package query

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
)

type SaleItemFindByIdQuery struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

func NewSaleItemFindByIdQuery(tenantId, id string) *SaleItemFindByIdQuery {
	return &SaleItemFindByIdQuery{
		TenantId: tenantId,
		Id:       id,
	}
}

type SaleItemFindByIdsQuery struct {
	TenantId string   `json:"tenantId"`
	Ids      []string `json:"ids"`
}

func NewSaleItemFindByIdsQuery(tenantId string, ids []string) *SaleItemFindByIdsQuery {
	return &SaleItemFindByIdsQuery{
		TenantId: tenantId,
		Ids:      ids,
	}
}

type SaleItemFindAllQuery struct {
	TenantId string `json:"tenantId"`
}

func NewSaleItemFindAllQuery(tenantId string) *SaleItemFindAllQuery {
	return &SaleItemFindAllQuery{
		TenantId: tenantId,
	}
}

type SaleItemFindPagingQuery struct {
	TenantId string `json:"tenantId"`
	Fields   string `json:"fields"`
	Filter   string `json:"filter"`
	Sort     string `json:"sort"`
	PageNum  int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

func NewSaleItemFindPagingQuery(tenantId string, fields string, filter string, sort string, pageNum int64, pageSize int64) *SaleItemFindPagingQuery {
	return &SaleItemFindPagingQuery{
		TenantId: tenantId,
		Fields:   fields,
		Filter:   filter,
		Sort:     sort,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func (q *SaleItemFindPagingQuery) GetTenantId() string {
	return q.TenantId
}

func (q *SaleItemFindPagingQuery) GetFields() string {
	return q.Fields
}

func (q *SaleItemFindPagingQuery) GetFilter() string {
	return q.Filter
}

func (q *SaleItemFindPagingQuery) GetSort() string {
	return q.Sort
}

func (q *SaleItemFindPagingQuery) GetPageNum() int64 {
	return q.PageNum
}

func (q *SaleItemFindPagingQuery) GetPageSize() int64 {
	return q.PageSize
}

type SaleItemFindBySaleBillIdQuery struct {
	TenantId   string `json:"tenantId"`
	SaleBillId string `json:"saleBillId"`
}

func NewSaleItemFindBySaleBillIdQuery(tenantId string, saleBillId string) *SaleItemFindBySaleBillIdQuery {
	return &SaleItemFindBySaleBillIdQuery{
		TenantId:   tenantId,
		SaleBillId: saleBillId,
	}
}

type SaleItemFindPagingResult struct {
	Data       []*view.SaleItemView `json:"data"`
	TotalRows  *int64               `json:"totalRows"`
	TotalPages *int64               `json:"totalPages"`
	PageNum    int64                `json:"pageNum"`
	PageSize   int64                `json:"pageSize"`
	Filter     string               `json:"filter"`
	Sort       string               `json:"sort"`
	Error      error                `json:"-"`
	IsFound    bool                 `json:"-"`
}

func NewSaleItemFindPagingResult() *SaleItemFindPagingResult {
	return &SaleItemFindPagingResult{}
}
