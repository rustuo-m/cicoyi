package common

import (
	"fmt"
	"os"
)

// GetDir 获取路径
func GetDir() string {
	//_, fileName, _, ok := runtime.Caller(0)
	//if !ok {
	//	log.Panicln("获取根目录失败")
	//	return ""
	//}
	//// 根目录
	//rootDir := filepath.Dir(filepath.Dir(fileName))
	// 获取当前工作目录
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return rootDir
}
