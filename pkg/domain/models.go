package domain

type User struct{
	UserId   int 	`gorm:"primaryKey"`
	UserName string `json:"u_name"`
	Password string `json:"password"`
	Verification bool
}