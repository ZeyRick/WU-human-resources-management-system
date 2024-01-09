package dtos

type UserLogin struct {
	Username string `json:"username,string,required"`
	Password string `json:"password,string,required"`
}

type UserRegister struct {
	Username   string `json:"username,string,required"`
	Name       string `json:"name,string,required"`
	Password   string `json:"password,string,required"`
	ProfilePic string `json:"profilePic,string,omitempty"`
}

type GetUserData struct {
	DataPerPage *int64 `json:"dataPerPage,int,required"`
	PageNumber  *int64 `json:"pageNumber,int,required"`
}
