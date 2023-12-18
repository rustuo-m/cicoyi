package common

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigInitialize struct {
}

func (c ConfigInitialize) InitConfig() {
	// 获取根目录
	dir := GetDir()

	// 设置Viper配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

}
