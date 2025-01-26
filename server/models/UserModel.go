package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role" gorm:"default:'user'"`
	Active   bool   `json:"active" gorm:"default:true"`
}

type MCQ struct {
	*gorm.Model
}
