package dtos

type UserLogin struct {
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type UserRegister struct {
	Username   string `json:"username,required"`
	Name       string `json:"name,required"`
	Password   string `json:"password,required"`
	ProfilePic string `json:"profilePic,omitempty"`
}

type ListUser struct {
	Name    string  `json:"name,string,omitempty"`
}
