package dtos

type AddSchedule struct {
	EmployeeId *int   `json:"employeeId,string,required"`
	Scope      string `json:"scope"`
	Dates      []int  `json:"dates,numbers"`
}

type ScheduleFilter struct {
	EmployeeId *int    `json:"employeeId,omitempty"`
	Scope      string  `json:"scope,omitempty"`
}
