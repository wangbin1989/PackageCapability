package common

import (
	"log"
	"os"
)

func EnsureNoError(err error, format string) {
	if err != nil {
		log.Fatalf(format, err)
	}
}

func EnsureKeyExists(exists bool, key any) {
	if !exists {
		log.Fatalf("Key %v not exists.", key)
	}
}

func EnsureFileExists(file string) {
	_, err := os.Stat(file)

	if err == nil {
		return
	}

	if os.IsNotExist(err) {
		log.Fatalf("文件 %s 不存在。", file)
	}

	log.Fatalf("检查文件时出错: %v", err)
}
