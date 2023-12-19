package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	Password      string
	Email         string
	Phone         string
	Gender        string
	Birth         string
	Organizations []Organization `gorm:"foreignKey:UserID"`
}
