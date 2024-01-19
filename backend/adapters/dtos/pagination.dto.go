package dtos

// type-parametrized struct
type FilterWithPageOpt[T any] struct {
	Filter T
	PageOpt 
}

type PageOpt struct {
	Size *int64 `json:"size,int" validate:"required"`
	Page *int64 `json:"page,int" validate:"required"`
}
