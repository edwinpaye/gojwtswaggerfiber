package models

import (
	"time"
)

type User struct {
	Id int `json:"id" gorm:"primaryKey"`
	// ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Email     string    `json:"email" validate:"email, required" gorm:"unique;not null"`
	Password  string    `json:"password" validate:"min=3, max=100"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `json:"title"`
	Lastname  string    `json:"author"`
	Age       string    `json:"description"`
}
