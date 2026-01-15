package main

import (
	"fmt"
)

// 数组是固定长度的同类型元素序列。
func main() {
	// 声明数组,从下标0开始
	var numbers [5]int
	numbers[0] = 0
	numbers[1] = 1
	numbers[2] = 4
	numbers[3] = 2
	numbers[4] = 5

	// 数组字面量
	var friutes = [5]string{"apple", "banana", "orange", "grape", "pear"}

	// 让编译器推断长度
	colors := [...]string{"red", "blue", "orange", "green"}

	// 访问数组元素
	fmt.Printf("第一个水果： %s.\n", friutes[0])
	fmt.Printf("第三个颜色：%s.\n", colors[2])

	// 遍历数组
	fmt.Println("所有的水果：")
	for i, fruit := range friutes {
		fmt.Printf(" %d: %s.\n", i, fruit)
	}

	// 多维数组
	var matrix [3][3]int
	matrix[0][0] = 0
	matrix[1][1] = 1
	matrix[2][2] = 3

	fmt.Println(len(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

}
