package types

import "time"

type ScheduleInfo struct {
	Scope     string             `json:"scope"`
	Employees []FormatedEmployee `json:"employees"`
}

type AddSchedule struct {
	EmployeeIds     *[]int     `json:"employeeId" validate:"required"`
	Scope           string     `json:"scope" validate:"required"`
	Dates           string     `json:"dates" validate:"required"`
	ClockInTime     *time.Time `json:"clockInTime" validate:"required"`
	ClockOutTime    *time.Time `json:"clockOutTime" validate:"required"`
	MinuteBreakTime *int       `json:"minuteBreakTime" validate:"required"`
}

type UpdateSchedule struct {
	EmployeeIds  *[]int     `json:"employeeId" validate:"required"`
	Scope        string     `json:"scope" validate:"required"`
	Dates        string     `json:"dates" validate:"required"`
	ClockInTime  *time.Time `json:"clockInTime" validate:"required"`
	ClockOutTime *time.Time `json:"clockOutTime" validate:"required"`
}
