package dtos

type AddEmployee struct {
	Name       string `json:"name,string,required"`
	ProfilePic string `json:"profilePic,string,omitempty"`
}

type EmployeeFilter struct {
	ID   *int   `json:"id,string,omitempty"`
	Name string `json:"name,string,omitempty"`
}
