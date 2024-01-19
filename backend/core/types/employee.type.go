package types

type FormatedEmployee struct {
	Name         string `json:"name"`
	DepartmentId *int   `json:"departmentId"`
	ProfilePic   string `json:"profilePic"`
	ClockInTime  string `json:"clockInTime"`
	ClockOutTime string `json:"clockOutTime"`
}
