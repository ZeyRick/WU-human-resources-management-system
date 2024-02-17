package types

import "time"

type FormatedEmployee struct {
	Name         string `json:"name"`
	DepartmentId *int   `json:"departmentId"`
	ProfilePic   string `json:"profilePic"`
	ClockInTime  string `json:"clockInTime"`
	ClockOutTime string `json:"clockOutTime"`
}

type EmployeeWithSchedule struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name" gorm:"type:string;not null"`
	DepartmentId *int   `json:"departmentId" gorm:"type:number;not null"`
	Department   struct {
		DepartmentId uint      `json:"departmentId" gorm:"primaryKey;autoIncrement"`
		Alias        string    `json:"alias" gorm:"type:string;not null"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
	} `json:"department" gorm:"embedded"`
	ProfilePic string `json:"profilePic" gorm:"type:string;not null"`
	Schedules  struct {
		ScheduleId   uint      `json:"scheduleId"`
		EmployeeId   int       `json:"employeeId" gorm:"type:int;not null"`
		Scope        string    `json:"scope" gorm:"type:string;not null"`
		Dates        string    `json:"dates" gorm:"type:string"`
		ClockInTime  time.Time `json:"clockInTime"`
		ClockOutTime time.Time `json:"clockOutTime"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
	} `json:"schedule" gorm:"embedded"`
}
