package appquery

import "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"

type UserFindByIdAppQuery struct {
	TenantId string
	Id       string
}

func NewUserFindByIdAppQuery() *UserFindByIdAppQuery {
	return &UserFindByIdAppQuery{}
}

type UserFindByIdsAppQuery struct {
	TenantId string
	Ids      []string
}

func NewUserFindByIdsAppQuery() *UserFindByIdsAppQuery {
	return &UserFindByIdsAppQuery{}
}

type UserFindByAggregateIdAppQuery struct {
	TenantId    string
	AggregateId string
}

func NewUserFindByAggregateIdAppQuery() *UserFindByAggregateIdAppQuery {
	return &UserFindByAggregateIdAppQuery{}
}

type UserFindPagingAppQuery struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func NewUserFindPagingAppQuery() *UserFindPagingAppQuery {
	return &UserFindPagingAppQuery{}
}

type UserFindPagingResult struct {
	Data       []*view.UserView `json:"data"`
	TotalRows  int64            `json:"totalRows"`
	TotalPages int64            `json:"totalPages"`
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

type UserFindAllAppQuery struct {
	TenantId string
}

func NewUserFindAllAppQuery() *UserFindAllAppQuery {
	return &UserFindAllAppQuery{}
}

type UserFindByUserIdAppQuery struct {
	TenantId string
	UserId   string
}

func NewUserFindByUserIdAppQuery() *UserFindByUserIdAppQuery {
	return &UserFindByUserIdAppQuery{}
}
