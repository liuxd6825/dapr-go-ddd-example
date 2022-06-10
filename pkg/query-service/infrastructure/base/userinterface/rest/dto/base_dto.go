package dto

type FindByIdQuery struct {
	TenantId string
	Id       string
}

type FindAllQuery struct {
	TenantId string
}

type FindPagingQuery struct {
	TenantId string
	PageNum  int64
	PageSize int64
	Filter   string
	Sort     string
	Fields   string
}

func (q *FindPagingQuery) GetTenantId() string { return q.TenantId }
func (q *FindPagingQuery) GetPageNum() int64   { return q.PageNum }
func (q *FindPagingQuery) GetPageSize() int64  { return q.PageSize }
func (q *FindPagingQuery) GetFilter() string   { return q.Filter }
func (q *FindPagingQuery) GetSort() string     { return q.Sort }
func (q *FindPagingQuery) GetFields() string   { return q.Fields }

func (q *FindPagingQuery) SetTenantId(value string) { q.TenantId = value }
func (q *FindPagingQuery) SetPageNum(value int64)   { q.PageNum = value }
func (q *FindPagingQuery) SetPageSize(value int64)  { q.PageSize = value }
func (q *FindPagingQuery) SetFilter(value string)   { q.Filter = value }
func (q *FindPagingQuery) SetSort(value string)     { q.Sort = value }
func (q *FindPagingQuery) SetFields(value string)   { q.Fields = value }

//
// FindPagingResponse
// @Description: FindPagingResponse
//
type FindPagingResponse[T any] struct {
	Data       *[]T   `json:"data"`
	TotalRows  int64  `json:"totalRows"`
	TotalPages int64  `json:"totalPages"`
	PageNum    int64  `json:"pageNum"`
	PageSize   int64  `json:"pageSize"`
	Filter     string `json:"filter"`
	Sort       string `json:"sort"`
}
