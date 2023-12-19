package common

import (
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Database struct{}

var DB *gorm.DB

func (d Database) InitDatabase() {
	DB, err := gorm.Open(mysql.Open(viper.GetString("database")))
	if err != nil {
		log.Fatalln("failed to connect to database", err.Error())
		return
	}
	err = DB.AutoMigrate(model.Organization{}, model.User{})
	if err != nil {
		log.Fatalln("Failed to autoMigrate: ", err.Error())
		return
	}
}

func (d Database) GetDB() *gorm.DB {
	return DB
}
