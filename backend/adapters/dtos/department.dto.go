package dtos

type DepartmentFilter struct {
	ID   *int   `json:"id,string,omitempty"`
	Alias string `json:"name,string,omitempty"`
}
