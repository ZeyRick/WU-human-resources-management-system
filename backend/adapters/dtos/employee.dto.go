package dtos

type AddEmployee struct {
	Name         string `json:"name" validate:"required"`
	ProfilePic   string `json:"profilePic" validate:"omitempty"`
	DepartmentId int    `json:"departmentId" validate:"required"`
}

type EmployeeFilter struct {
	EmployeeName string `json:"employeeName" validate:"omitempty"`
	DepartmentId *int    `json:"departmentId" validate:"omitempty"`
	EmployeeId   *int    `json:"employeeId" validate:"omitempty"`
}
