package dtos

type AddEmployee struct {
	Name         string `json:"name,string" validate:"required"`
	ProfilePic   string `json:"profilePic,string,omitempty"`
	DepartmentId int    `json:"departmentId,number" validate:"required"`
}

type EmployeeFilter struct {
	ID   *int   `json:"id,string,omitempty"`
	Name string `json:"name,string,omitempty"`
}
