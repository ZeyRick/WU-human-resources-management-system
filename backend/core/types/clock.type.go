package types

type ClockType string

var (
	ClockIn  ClockType = "in"
	ClockOut ClockType = "out"
)

type ClockReports struct {
	TotalWorkMinue *int   `json:"totalWorkMinute" gorm:"type:string;not null"`
	EmployeeId     *int   `json:"employeeId" gorm:"type:string;not null"`
	Name           string `json:"employeeName" gorm:"type:string;not null"`
	DepartmentId   *int   `json:"departmentId" gorm:"type:number;not null"`
	Alias          string `json:"departmentAlias" gorm:"type:string;not null"`
}
