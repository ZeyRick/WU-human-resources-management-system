package dtos

type AddEmployee struct {
	Name         string `json:"name" validate:"required"`
	ProfilePic   string `json:"profilePic" validate:"omitempty"`
	DepartmentId int    `json:"departmentId" validate:"required"`
}

type EmployeeFilter struct {
	ID           *int   `json:"id" validate:"omitempty"`
	Name         string `json:"name" validate:"omitempty"`
	DepartmentId *int   `json:"departmentId" validate:"omitempty"`
	EmployeeId   *int   `json:"employeeId" validate:"omitempty"`
}
