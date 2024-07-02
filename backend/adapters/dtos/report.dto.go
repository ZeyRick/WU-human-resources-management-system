package dtos

type ReportFilter struct {
	StartDate    string   `json:"startDate,string,omitempty"`
	EndDate      string   `json:"endDate,string,omitempty"`
	EmployeeName string   `json:"employeeName,string,omitempty"`
	CourseId     int      `json:"courseId,string,omitempty"`
	EmployeeType []string `json:"employeeType,omitempty"`
	IsTeaching   bool     `json:"isTeaching,omitempty"`
}
