package dtos

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegister struct {
	Username   string `json:"username" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Password   string `json:"password" validate:"required"`
	ProfilePic string `json:"profilePic,omitempty"`
}

type ListUser struct {
	Name string `json:"name,string,omitempty"`
}

type UserResetPwParams struct {
	Password string `json:"password" validate:"required"`
}
