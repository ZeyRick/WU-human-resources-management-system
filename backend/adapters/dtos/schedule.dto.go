package dtos

type AddSchedule struct {
	EmployeeIds     *[]int `json:"employeeId" validate:"required"`
	Scope           string `json:"scope" validate:"required"`
	Dates           string `json:"dates" validate:"required"`
	ClockInTime     string `json:"clockInTime" validate:"required"`
	ClockOutTime    string `json:"clockOutTime" validate:"required"`
	MinuteBreakTime *int   `json:"minuteBreakTime" validate:"required"`
}

type ScheduleFilter struct {
	EmployeeId   *int   `json:"employeeId,omitempty"`
	Scope        string `json:"scope" validate:"required"`
	DepartmentId *int   `json:"departmentId" validate:"required"`
}
