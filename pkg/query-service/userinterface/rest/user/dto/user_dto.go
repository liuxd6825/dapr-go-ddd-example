package dto

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

//
// UserFindByIdRequest
// @Description:  请求内容
//
type UserFindByIdRequest struct {
	TenantId string
	Id       string
}

func (r *UserFindByIdRequest) GetTenantId() string      { return r.TenantId }
func (r *UserFindByIdRequest) SetTenantId(value string) { r.TenantId = value }

func (r *UserFindByIdRequest) GetId() string      { return r.Id }
func (r *UserFindByIdRequest) SetId(value string) { r.Id = value }

//
// UserFindByIdResponse
// @Description:  请求内容
//
type UserFindByIdResponse struct {
	UserResponseItem
}

//
// UserFindAllRequest
// @Description:
//
type UserFindAllRequest struct {
	TenantId string
}

func (r *UserFindAllRequest) GetTenantId() string      { return r.TenantId }
func (r *UserFindAllRequest) SetTenantId(value string) { r.TenantId = value }

type UserFindAllResponse []*UserFindAllResponseItem

type UserFindAllResponseItem struct {
	UserResponseItem
}

//
// UserFindPagingRequest
// @Description:
//
type UserFindPagingRequest struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func (r *UserFindPagingRequest) GetTenantId() string      { return r.TenantId }
func (r *UserFindPagingRequest) SetTenantId(value string) { r.TenantId = value }

func (r *UserFindPagingRequest) GetPageNum() int64      { return r.PageNum }
func (r *UserFindPagingRequest) SetPageNum(value int64) { r.PageNum = value }

func (r *UserFindPagingRequest) GetPageSize() int64      { return r.PageSize }
func (r *UserFindPagingRequest) SetPageSize(value int64) { r.PageSize = value }

func (r *UserFindPagingRequest) GetFilter() string      { return r.Filter }
func (r *UserFindPagingRequest) SetFilter(value string) { r.Filter = value }

func (r *UserFindPagingRequest) GetSort() string      { return r.Sort }
func (r *UserFindPagingRequest) SetSort(value string) { r.Sort = value }

func (r *UserFindPagingRequest) GetFields() string      { return r.Fields }
func (r *UserFindPagingRequest) SetFields(value string) { r.Fields = value }

//
// UserFindPagingResponse
// @Description:
//
type UserFindPagingResponse struct {
	Data       *[]*UserFindPagingResponseItem `json:"data"`
	TotalRows  int64                          `json:"totalRows"`
	TotalPages int64                          `json:"totalPages"`
	PageNum    int64                          `json:"pageNum"`
	PageSize   int64                          `json:"pageSize"`
	Filter     string                         `json:"filter"`
	Sort       string                         `json:"sort"`
}

type UserFindPagingResponseItem struct {
	UpdateTime *types.TimeString `json:"updateTime"`
	UserResponseItem
}

//
// UserResponseItem
// @Description:   请求业务数据
//
type UserResponseItem struct {
	Id         string            `json:"id"`
	TenantId   string            `json:"tenantId"`
	UserCode   string            `json:"userCode"`
	UserName   string            `json:"userName"`
	Email      string            `json:"email"`
	Telephone  string            `json:"telephone"`
	Address    string            `json:"address"`
	CreateTime *types.DateString `json:"createTime"`
}
