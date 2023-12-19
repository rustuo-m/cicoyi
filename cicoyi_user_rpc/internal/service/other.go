package service

import (
	"github.com/rustuo-m/cicoyi/cicoyi_user_rpc/internal/common"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return common.Database{}.GetDB()
}
