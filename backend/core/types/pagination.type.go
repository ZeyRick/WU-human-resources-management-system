package types

type ListData[T any] struct {
	Data    *[]T        `json:"data"`
	PageOpt *Pagination `json:"pageOpt"`
}

type Pagination struct {
	TotalCount *int64
	CurPage    *int64
	TotalPage  *int64
	PageSize   *int64
}
