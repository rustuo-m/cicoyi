package models

import (
	"log"
	
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Birth    string `json:"birth"`
}

func (User) TableName() string {
	return "user"
}

// ExistLogin 存在返回true 不存在返回false
func (u *User) ExistLogin() (*User, int) {
	var userDataBase *User
	tx := DB.Where("name = ?", u.Name).First(&userDataBase)
	if tx.Error != nil {
		log.Printf("get userHandlers %s error %v", u.Name, tx.Error)
		return nil, global.ERROR
	}
	if tx.RowsAffected == 0 {
		log.Printf("userHandlers %s not found", u.Name)
		return nil, global.NotFind
	} else {
		return userDataBase, global.SUCCESS
	}
}

func (u *User) ExistRegister() (*User, int) {
	var userDataBase *User
	tx := DB.Where("name = ?", u.Name).First(&userDataBase)
	if tx.RowsAffected == 0 {
		return nil, global.NotFind
	} else if tx.Error != nil {
		log.Printf("err:%v", tx.Error)
		return nil, global.ERROR
	} else {
		return userDataBase, global.SUCCESS
	}
}

func (u *User) Create() int {
	err := DB.Create(&u).Error
	if err != nil {
		log.Printf("create userHandlers %s error: %v", u.Name, err)
		return global.ERROR
	}
	return global.SUCCESS
}
