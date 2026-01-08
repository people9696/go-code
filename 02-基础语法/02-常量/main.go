package main

import "fmt"

func main() {
	// 单行常量
	const pi = 3.1415926
	const AppName = "Go Master"

	// 多行常量
	const (
		StatusActive   = "active"
		StatusInactive = "inactive"
		StatusPending  = "pending"
	)

	// iota枚举
	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Printf("应用名称：%s.\n", AppName)
	fmt.Printf("圆周率是：%.5f.\n", pi)
	fmt.Printf("状态是：%s.\n", StatusInactive)
	fmt.Printf("今天是周%d.\n", Wednesday)
}
