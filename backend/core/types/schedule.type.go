package types

import "time"

type ScheduleInfo struct {
	Scope     string     `json:"scope"`
	Employees []FormatedEmployee `json:"employees"`
}

type AddSchedule struct {
	EmployeeId   *int   `json:"employeeId" validate:"required"`
	Scope        string `json:"scope" validate:"required"`
	Dates        string  `json:"dates" validate:"required"`
	ClockInTime  *time.Time `json:"clockInTime" validate:"required"`
	ClockOutTime *time.Time `json:"clockOutTime" validate:"required"`
	DepartmentId *int `json:"departmentId" validate:"required"`
}
