package models

type User struct {
	BaseModel
	Username   string `json:"username" gorm:"type:string;not null"`
	Name       string `json:"name" gorm:"type:string;not null"`
	Password   string `json:"password" gorm:"type:string;not null"`
	UserLevel  string `json:"userLevel" gorm:"type:string;not null"`
	ProfilePic string `json:"profilePic" gorm:"type:string;not null"`
}