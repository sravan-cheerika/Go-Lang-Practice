package models

type User struct{
	gorm.Model
	ID uint `json:"id" gorm:"primary_key"`
	UserName string `json:"name"`
	Age uint `json:"age"`
}
