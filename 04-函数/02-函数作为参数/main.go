package main

import (
	"fmt"
	"log"
	"os"
)

// 定义函数类型
type Opertation func(int, int) int

// 加法
func add(a, b int) int {
	return a + b
}

// 减法
func subtract(a, b int) int {
	return a - b
}

// 乘法
func multiply(a, b int) int {
	return a * b
}

// 高阶函数
func calculate(a, b int, op Opertation) int {
	return op(a, b)
}

// 返回函数的函数
func getOperation(name string) Opertation {
	switch name {
	case "add":
		return add
	case "subtract":
		return subtract
	case "multiply":
		return multiply
	default:
		return nil
	}
}

func main() {

	// 日志处理
	file, err := os.OpenFile(
		"app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	// 直接调用
	retval1 := calculate(10, 20, add)
	retval2 := calculate(40, 50, subtract)
	log.Printf("10 + 5 = %d", retval1)
	fmt.Println("calculate(10, 20, add):", retval1, ".")
	fmt.Println("calculate(40, 50, subtract):", retval2, ".")

	//从函数获取操作
	op := getOperation("multiply")
	if op != nil {
		retval3 := calculate(10, 9, op)
		fmt.Println("calculate(10, 9, op):", retval3)
	}

	// 匿名函数
	divide := func(a, b int) int {
		return a / b
	}
	retval4 := calculate(10, 5, divide)
	fmt.Println("retval4:", retval4)

	// fmt.Printf方法
	fmt.Println("=========================")
	fmt.Printf("x=%d, y=%.2f\n", 10, 3.14)
	fmt.Printf("msg=%s\n", "hi")
	fmt.Printf("any=%v\n", 100)
	fmt.Printf("type=%T\n", 100)
	fmt.Printf("|%6d|\n", 10)
	fmt.Printf("|%-6d|\n", 10)
}
