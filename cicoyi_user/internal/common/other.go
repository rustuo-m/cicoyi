package common

import (
	"fmt"
	"os"
)

// GetDir 获取路径
func GetDir() string {
	// 获取当前工作目录
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return rootDir
}
