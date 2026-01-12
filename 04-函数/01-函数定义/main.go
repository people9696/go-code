package main

import (
	"fmt"
)

// 基本函数
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 带返回值的函数
func add(num1, num2 int) int {
	return num1 + num2
}

// 多返回值
func calculate(a, b int) (int, int, int) {
	sum := a + b
	difference := a - b
	product := a * b
	return sum, difference, product
}

// 命名返回值
func divmod(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return // 隐式返回

}

// 可变参数
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
func main() {

	greet("mike")

	num1 := 10
	num2 := 8

	fmt.Printf("10 + 10 = %d.\n", add(num1, num2))

	sum, different, product := calculate(num1, num2)
	fmt.Printf("num1 + num2 = %d.\nnum1 - num2 = %d.\nnum1 * num2 = %d.\n", sum, different, product)

	quotient, remainder := divmod(num1, num2)
	fmt.Printf("num1 / num2 = %d.\nnum1 %% num2 = %d.\n", quotient, remainder)

	num3 := 7
	num4 := 5
	sumNumber := sumNumbers(num1, num2, num3, num4)
	fmt.Printf("total : %d.\n", sumNumber)
}
