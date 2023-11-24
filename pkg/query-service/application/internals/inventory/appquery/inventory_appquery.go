package appquery

import "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"

type InventoryFindByIdAppQuery struct {
	TenantId string
	Id       string
}

func NewInventoryFindByIdAppQuery() *InventoryFindByIdAppQuery {
	return &InventoryFindByIdAppQuery{}
}

type InventoryFindByIdsAppQuery struct {
	TenantId string
	Ids      []string
}

func NewInventoryFindByIdsAppQuery() *InventoryFindByIdsAppQuery {
	return &InventoryFindByIdsAppQuery{}
}

type InventoryFindByAggregateIdAppQuery struct {
	TenantId    string
	AggregateId string
}

func NewInventoryFindByAggregateIdAppQuery() *InventoryFindByAggregateIdAppQuery {
	return &InventoryFindByAggregateIdAppQuery{}
}

type InventoryFindPagingAppQuery struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func NewInventoryFindPagingAppQuery() *InventoryFindPagingAppQuery {
	return &InventoryFindPagingAppQuery{}
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

type InventoryFindAllAppQuery struct {
	TenantId string
}

func NewInventoryFindAllAppQuery() *InventoryFindAllAppQuery {
	return &InventoryFindAllAppQuery{}
}

type InventoryFindByInventoryIdAppQuery struct {
	TenantId    string
	InventoryId string
}

func NewInventoryFindByInventoryIdAppQuery() *InventoryFindByInventoryIdAppQuery {
	return &InventoryFindByInventoryIdAppQuery{}
}
