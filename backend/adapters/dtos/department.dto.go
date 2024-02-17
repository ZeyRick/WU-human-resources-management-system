package dtos

type DepartmentFilter struct {
	ID   *int   `json:"id,string,omitempty"`
	Alias string `json:"name,string,omitempty"`
}

type AddDepartment struct {
	Alias string `json:"alias" validate:"required"`
}
