package dtos

type PageOpt struct {
	Size *int `json:"size,int,required"`
	Page *int `json:"page,int,required"`
}
