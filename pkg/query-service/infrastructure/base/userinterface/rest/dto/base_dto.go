package dto

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/assert"
)

//
// FindRequest
// @Description: 查询请求内容基类
//
type FindRequest struct {
	tenantId string
}

func (r FindRequest) TenantId() string {
	return r.tenantId
}

//
// Init
// @Description: 从上下文中初始化
// @receiver r
// @param ctx
// @return error
//
func (r *FindRequest) Init(ctx iris.Context) error {
	r.tenantId = ctx.Params().Get("tenantId")
	if err := assert.NotEmpty(r.tenantId, assert.NewOptions("url \"{tenantId}\" cannot be empty")); err != nil {
		return err
	}
	return nil
}

//
// FindResponse
// @Description: 查询响应内容基类
//
type FindResponse struct {
}

//
// FindByIdRequest
// @Description: 按ID查找响应内容
//
type FindByIdRequest struct {
	FindRequest
	id string
}

//
// Id
// @Description: 获取ID
// @receiver r
// @return string
//
func (r FindByIdRequest) Id() string {
	return r.id
}

//
// Init
// @Description: 初始化按Id查询请求内容
// @receiver r
// @param ctx
// @return error
//
func (r *FindByIdRequest) Init(ctx iris.Context) error {
	if err := r.FindRequest.Init(ctx); err != nil {
		return err
	}
	r.id = ctx.Params().Get("id")
	if err := assert.NotEmpty(r.id, assert.NewOptions("url \"{id}\" cannot be empty")); err != nil {
		return err
	}
	return nil
}

//
// FindByIdResponse
// @Description: 按ID查询的响应内容
//
type FindByIdResponse struct {
	FindResponse
}

//
// FindAllRequest
// @Description: 查询所有数据的请求内容
//
type FindAllRequest struct {
	FindRequest
}

//
// FindAllResponse
// @Description: 查询所有数据的响应内容
//
type FindAllResponse[T any] List[T]

//
// FindPagingRequest
// @Description: 分页查询的请求内容
//
type FindPagingRequest[T any] struct {
	FindRequest
	pageNum  int64
	pageSize int64
	filter   string
	sort     string
	fields   string
}

//
// Init
// @Description: 初始化分页
// @receiver r
// @param ctx
// @return error
//
func (r *FindPagingRequest[T]) Init(ctx iris.Context) error {
	r.pageNum = ctx.URLParamInt64Default("pageNum", 0)
	r.pageSize = ctx.URLParamInt64Default("pageSize", 20)
	r.filter = ctx.URLParamDefault("filter", "")
	r.sort = ctx.URLParamDefault("sort", "")
	if err := r.FindRequest.Init(ctx); err != nil {
		return err
	}
	return nil
}

//
// TenantId
// @Description: 取租户信息
// @receiver r
// @return string
//
func (r *FindPagingRequest[T]) TenantId() string {
	return r.tenantId
}

//
// Fields
// @Description: 取查询字段
// @receiver r
// @return string
//
func (r *FindPagingRequest[T]) Fields() string {
	return r.fields
}

//
// Filter
// @Description: 取过滤条件
// @receiver r
// @return string
//
func (r *FindPagingRequest[T]) Filter() string {
	return r.filter
}

//
// Sort
// @Description: 取排序条件
// @receiver r
// @return string
//
func (r *FindPagingRequest[T]) Sort() string {
	return r.sort
}

//
// PageNum
// @Description: 取分页页号
// @receiver r
// @return int64
//
func (r *FindPagingRequest[T]) PageNum() int64 {
	return r.pageNum
}

//
// PageSize
// @Description: 取分页页大小
// @receiver r
// @return int64
//
func (r *FindPagingRequest[T]) PageSize() int64 {
	return r.pageSize
}

//
// FindPagingResponse
// @Description: FindPagingResponse
//
type FindPagingResponse struct {
	FindResponse
	Data       any    `json:"data"`
	TotalRows  int64  `json:"totalRows"`
	TotalPages int64  `json:"totalPages"`
	PageNum    int64  `json:"pageNum"`
	PageSize   int64  `json:"pageSize"`
	Filter     string `json:"filter"`
	Sort       string `json:"sort"`
}

//
// Init
// @Description: 初始化分页相应内容
// @receiver r
// @param fr  分页查询结果
// @param dtoList Dto对象列表
// @return error
//
func (r *FindPagingResponse) Init(fr FindPagingResult, dtoList interface{}) error {
	r.Filter = fr.GetFilter()
	r.Sort = fr.GetSort()

	r.PageSize = fr.GetPageSize()
	r.PageNum = fr.GetPageNum()

	r.TotalRows = fr.GetTotalRows()
	r.TotalPages = fr.GetTotalPages()

	r.Data = dtoList
	return nil
}

//
// List
// @Description: List
//
type List[T any] []T

//
// FindPagingResult
// @Description:  分页查询结果接口
//
type FindPagingResult interface {
	GetTotalRows() int64
	GetTotalPages() int64
	GetPageNum() int64
	GetPageSize() int64
	GetFilter() string
	GetSort() string
}
