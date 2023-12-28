package dtos

type AddEmployee struct {
	Name       string `json:"name,string,required"`
	ProfilePic string `json:"profilePic,string,omitempty"`
}

type ListEmployee struct {
	PageOpt PageOpt `json:"pageOpt,required"`
	Name    string  `json:"name,string,omitempty"`
}
