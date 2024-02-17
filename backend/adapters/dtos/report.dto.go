package dtos

type ReportFilter struct {
	StartDate    string `json:"startDate,string,omitempty"`
	EndDate      string `json:"endDate,string,omitempty"`
	EmployeeName string `json:"employeeName,string,omitempty"`
	DepartmentId int    `json:"departmentId,string,omitempty"`
}
