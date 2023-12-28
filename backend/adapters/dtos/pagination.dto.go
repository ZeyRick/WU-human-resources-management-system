package dtos

type PageOpt struct {
	Size *int64 `json:"size,int,required"`
	Page *int64 `json:"page,int,required"`
}
