package types

type ListData[T any] struct {
	Data    *[]T        `json:"data"`
	PageOpt *Pagination `json:"pageOpt"`
}

type Pagination struct {
	TotalCount *int64 `json:"totalCount"`
	CurPage    *int64 `json:"curPage"`
	TotalPage  *int64 `json:"totalPage"`
	PageSize   *int64 `json:"pageSize"`
}
