package models

import "gorm.io/gorm"

type Question struct {
	*gorm.Model
	question string
	options  []string
}
