package entity

import "gorm.io/gorm"

type Object struct {
	gorm.Model
	Component1 string
}
