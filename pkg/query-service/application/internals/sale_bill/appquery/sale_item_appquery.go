package appquery

import "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"

type SaleItemFindByIdAppQuery struct {
	TenantId string
	Id       string
}

func NewSaleItemFindByIdAppQuery() *SaleItemFindByIdAppQuery {
	return &SaleItemFindByIdAppQuery{}
}

type SaleItemFindByIdsAppQuery struct {
	TenantId string
	Ids      []string
}

func NewSaleItemFindByIdsAppQuery() *SaleItemFindByIdsAppQuery {
	return &SaleItemFindByIdsAppQuery{}
}

type SaleItemFindByAggregateIdAppQuery struct {
	TenantId    string
	AggregateId string
}

func NewSaleItemFindByAggregateIdAppQuery() *SaleItemFindByAggregateIdAppQuery {
	return &SaleItemFindByAggregateIdAppQuery{}
}

type SaleItemFindPagingAppQuery struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func NewSaleItemFindPagingAppQuery() *SaleItemFindPagingAppQuery {
	return &SaleItemFindPagingAppQuery{}
}

type SaleItemFindPagingResult struct {
	Data       []*view.SaleItemView `json:"data"`
	TotalRows  int64                `json:"totalRows"`
	TotalPages int64                `json:"totalPages"`
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

type SaleItemFindAllAppQuery struct {
	TenantId string
}

func NewSaleItemFindAllAppQuery() *SaleItemFindAllAppQuery {
	return &SaleItemFindAllAppQuery{}
}

type SaleItemFindBySaleBillIdAppQuery struct {
	TenantId   string
	SaleBillId string
}

func NewSaleItemFindBySaleBillIdAppQuery() *SaleItemFindBySaleBillIdAppQuery {
	return &SaleItemFindBySaleBillIdAppQuery{}
}
