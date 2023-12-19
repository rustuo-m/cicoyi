package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string
	Description string
	UserID      uint64
	User        User `gorm:"foreignKey:UserID"`
}
