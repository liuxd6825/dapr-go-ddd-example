package query

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
)

type UserFindByIdQuery struct {
	TenantId string `json:"tenantId"`
	Id       string `json:"id"`
}

func NewUserFindByIdQuery(tenantId, id string) *UserFindByIdQuery {
	return &UserFindByIdQuery{
		TenantId: tenantId,
		Id:       id,
	}
}

type UserFindByIdsQuery struct {
	TenantId string   `json:"tenantId"`
	Ids      []string `json:"ids"`
}

func NewUserFindByIdsQuery(tenantId string, ids []string) *UserFindByIdsQuery {
	return &UserFindByIdsQuery{
		TenantId: tenantId,
		Ids:      ids,
	}
}

type UserFindAllQuery struct {
	TenantId string `json:"tenantId"`
}

func NewUserFindAllQuery(tenantId string) *UserFindAllQuery {
	return &UserFindAllQuery{
		TenantId: tenantId,
	}
}

type UserFindPagingQuery struct {
	TenantId string `json:"tenantId"`
	Fields   string `json:"fields"`
	Filter   string `json:"filter"`
	Sort     string `json:"sort"`
	PageNum  int64  `json:"pageNum"`
	PageSize int64  `json:"pageSize"`
}

func NewUserFindPagingQuery(tenantId string, fields string, filter string, sort string, pageNum int64, pageSize int64) *UserFindPagingQuery {
	return &UserFindPagingQuery{
		TenantId: tenantId,
		Fields:   fields,
		Filter:   filter,
		Sort:     sort,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}

func (q *UserFindPagingQuery) GetTenantId() string {
	return q.TenantId
}

func (q *UserFindPagingQuery) GetFields() string {
	return q.Fields
}

func (q *UserFindPagingQuery) GetFilter() string {
	return q.Filter
}

func (q *UserFindPagingQuery) GetSort() string {
	return q.Sort
}

func (q *UserFindPagingQuery) GetPageNum() int64 {
	return q.PageNum
}

func (q *UserFindPagingQuery) GetPageSize() int64 {
	return q.PageSize
}

type UserFindPagingResult struct {
	Data       []*view.UserView `json:"data"`
	TotalRows  *int64           `json:"totalRows"`
	TotalPages *int64           `json:"totalPages"`
	PageNum    int64            `json:"pageNum"`
	PageSize   int64            `json:"pageSize"`
	Filter     string           `json:"filter"`
	Sort       string           `json:"sort"`
	Error      error            `json:"-"`
	IsFound    bool             `json:"-"`
}

func NewUserFindPagingResult() *UserFindPagingResult {
	return &UserFindPagingResult{}
}
