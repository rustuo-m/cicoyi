package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LogInitialize struct {
}

func (l LogInitialize) InitLog() {
	// 根目录
	rootDir := GetDir()
	// 日志目录
	logDir := fmt.Sprintf("%s\\log\\", rootDir)

	// 确保日志目录存在
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Println("日志目录不存在")
		return
	}

	// 拼接日志文件名
	currentTime := time.Now().Format("2006_01_02")
	logFilePath := filepath.Join(logDir, currentTime+".log")
	// 打开日志文件(获取文件句柄)
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("日志文件打开错误")
		return
	}
	// 日志重定向输出到该文件
	log.SetOutput(logFile)
	// 打印日志（带时间戳和前缀）
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("日志服务启动")
}
