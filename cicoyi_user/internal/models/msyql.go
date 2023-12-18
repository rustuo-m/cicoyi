package models

import (
	"log"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type MysqlInitialize struct {
}

func (m MysqlInitialize) InitMysql() {
	DB, err = gorm.Open(mysql.Open("root:hzfs123456@tcp(127.0.0.1:3306)/cicoyi_user?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatalln("Unable to connect to the database:", err)
		return
	}
	err = DB.AutoMigrate(User{})
	if err != nil {
		log.Fatalln("Unable to AutoMigrate to the userHandlers Table:", err)
		return
	}
	log.Println("success to connect to the database")
}
