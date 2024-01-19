package dtos

type AddSchedule struct {
	EmployeeId *int   `json:"employeeId,string" validate:"required"`
	Scope      string `json:"scope"`
	Dates      []int  `json:"dates"`
}

type ScheduleFilter struct {
	EmployeeId   *int   `json:"employeeId,omitempty"`
	Scope        string `json:"scope" validate:"required"`
	DepartmentId *int   `json:"departmentId" validate:"required"`
}
