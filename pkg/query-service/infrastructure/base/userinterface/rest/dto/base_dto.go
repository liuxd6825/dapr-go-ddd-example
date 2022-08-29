package dto

import "time"

type FindByIdRequest struct {
	TenantId string
	Id       string
}

type FindByIdsRequest struct {
	TenantId string
	Ids      []string
}

type FindAllRequest struct {
	TenantId string
}

type FindPagingRequest struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func (q *FindPagingRequest) GetTenantId() string { return q.TenantId }
func (q *FindPagingRequest) GetPageNum() int64   { return q.PageNum }
func (q *FindPagingRequest) GetPageSize() int64  { return q.PageSize }
func (q *FindPagingRequest) GetFilter() string   { return q.Filter }
func (q *FindPagingRequest) GetSort() string     { return q.Sort }
func (q *FindPagingRequest) GetFields() string   { return q.Fields }

func (q *FindPagingRequest) SetTenantId(value string) { q.TenantId = value }
func (q *FindPagingRequest) SetPageNum(value int64)   { q.PageNum = value }
func (q *FindPagingRequest) SetPageSize(value int64)  { q.PageSize = value }
func (q *FindPagingRequest) SetFilter(value string)   { q.Filter = value }
func (q *FindPagingRequest) SetSort(value string)     { q.Sort = value }
func (q *FindPagingRequest) SetFields(value string)   { q.Fields = value }

//
// FindPagingResponse
// @Description: FindPagingResponse
//
type FindPagingResponse[T any] struct {
	Data       []T    `json:"data"`
	TotalRows  int64  `json:"totalRows"`
	TotalPages int64  `json:"totalPages"`
	PageNum    int64  `json:"pageNum"`
	PageSize   int64  `json:"pageSize"`
	Filter     string `json:"filter"`
	Sort       string `json:"sort"`
}

func NewFindPagingResponse[T any]() *FindPagingResponse[T] {
	res := FindPagingResponse[T]{}
	res.InitData()
	return &res
}

func (r *FindPagingResponse[T]) Init() {
	r.InitData()
}

func (r *FindPagingResponse[T]) InitData() {
	data := make([]T, 0)
	r.Data = data
}

type BaseDto struct {
	CreatedTime *time.Time `json:"createdTime,omitempty"` // 创建时间
	CreatorId   string     `json:"creatorId,omitempty"`   // 创建人ID
	CreatorName string     `json:"creatorName,omitempty"` // 创建人名称
	DeletedTime *time.Time `json:"deletedTime,omitempty"` // 删除时间
	DeleterId   string     `json:"deleterId,omitempty"`   // 删除人ID
	DeleterName string     `json:"deleterName,omitempty"` // 删除人名称
	Id          string     `json:"id,omitempty"`          // 主键
	IsDeleted   bool       `json:"isDeleted,omitempty"`   // 是否删除
	Remarks     string     `json:"remarks,omitempty"`     // 备注
	TenantId    string     `json:"tenantId,omitempty"`    // 租户ID
	UpdatedTime *time.Time `json:"updatedTime,omitempty"` // 修改时间
	UpdaterId   string     `json:"updaterId,omitempty"`   // 修改人ID
	UpdaterName string     `json:"updaterName,omitempty"` // 修改人名称
}
