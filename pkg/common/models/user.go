package models

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"title"`
	Lastname string `json:"author"`
	Age      string `json:"description"`
}
