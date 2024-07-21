package types

type DashboardSummary struct {
	EmployeeCounts []EmployeeCountType `json:"employeeCounts"`
	DegreeCount    int64               `json:"degreeCount"`
	CourseCount    int64               `json:"courseCount"`
	UserCount      int64               `json:"userCount"`
}

type EmployeeCountType struct {
	TotalCount   int64        `json:"totalCount"`
	EmployeeType EmployeeType `json:"employeeType"`
}
