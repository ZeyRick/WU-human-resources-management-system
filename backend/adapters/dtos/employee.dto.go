package dtos

type AddEmployee struct {
	Name       string `json:"name,string,required"`
	ProfilePic string `json:"profilePic,string,omitempty"`
}

type EmployeeFilter struct {
	Name    string  `json:"name,string,omitempty"`
}
