package main

import "fmt"

func main() {
	// 1. 使用var声明
	var name string = "Golang"
	var age int = 10
	var isOpen bool = true

	// 2. 类型推断
	var language = "Go"
	var version = 1.23

	// 3. 简短变量声明（只能在函数体内使用）
	message := "Hello world!"
	count := 42

	// 4. 多变量声明
	var (
		author  string = "Google"
		year    int    = 2009
		license string = "BSD"
	)

	fmt.Printf("语言：%s, 版本：%.2f, 授权协议: %s.\n", language, version, license)
	fmt.Printf("作者: %s, 年份: %d\n", author, year)
	fmt.Printf("名称：%s, 年龄：%d, 是否开源：%t.\n", name, age, isOpen)
	fmt.Printf("Message: %s, Count: %d.\n", message, count)
}
