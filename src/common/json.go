package common

import (
	"encoding/json"
	"log"
	"os"
)

func ReadFromJson[T any](path string) T {
	// 读取文件内容
	data, err := os.ReadFile(path)
	EnsureNoError(err, "读取文件错误: %v")

	var t T
	// 将JSON数据解码到结构体实例
	err = json.Unmarshal(data, &t)
	EnsureNoError(err, "JSON 解码错误: %v")

	return t
}

func WriteToJson[T any](path string, t T) {
	// 将结构体编码为JSON格式的字节切片，且进行缩进格式化
	data, err := json.MarshalIndent(t, "", "  ")
	EnsureNoError(err, "JSON 编码错误: %v")

	// 将JSON数据写入文件
	err = os.WriteFile(path, data, 0644)
	EnsureNoError(err, "写入文件错误: %v")

	log.Printf("JSON数据已成功写入文件. %s", path)
}
