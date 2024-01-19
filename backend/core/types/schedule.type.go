package types

type ScheduleInfo struct {
	Scope     string     `json:"scope"`
	Employees []FormatedEmployee `json:"employees"`
}
