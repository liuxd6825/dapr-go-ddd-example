package appquery

import "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"

type SaleBillFindByIdAppQuery struct {
	TenantId string
	Id       string
}

func NewSaleBillFindByIdAppQuery() *SaleBillFindByIdAppQuery {
	return &SaleBillFindByIdAppQuery{}
}

type SaleBillFindByIdsAppQuery struct {
	TenantId string
	Ids      []string
}

func NewSaleBillFindByIdsAppQuery() *SaleBillFindByIdsAppQuery {
	return &SaleBillFindByIdsAppQuery{}
}

type SaleBillFindByAggregateIdAppQuery struct {
	TenantId    string
	AggregateId string
}

func NewSaleBillFindByAggregateIdAppQuery() *SaleBillFindByAggregateIdAppQuery {
	return &SaleBillFindByAggregateIdAppQuery{}
}

type SaleBillFindPagingAppQuery struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func NewSaleBillFindPagingAppQuery() *SaleBillFindPagingAppQuery {
	return &SaleBillFindPagingAppQuery{}
}

type SaleBillFindPagingResult struct {
	Data       []*view.SaleBillView `json:"data"`
	TotalRows  int64                `json:"totalRows"`
	TotalPages int64                `json:"totalPages"`
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

type SaleBillFindAllAppQuery struct {
	TenantId string
}

func NewSaleBillFindAllAppQuery() *SaleBillFindAllAppQuery {
	return &SaleBillFindAllAppQuery{}
}

type SaleBillFindBySaleBillIdAppQuery struct {
	TenantId   string
	SaleBillId string
}

func NewSaleBillFindBySaleBillIdAppQuery() *SaleBillFindBySaleBillIdAppQuery {
	return &SaleBillFindBySaleBillIdAppQuery{}
}
