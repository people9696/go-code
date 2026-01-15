package main

import (
	"fmt"
	"sort"
)

// 过滤切片
func filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, val := range slice {
		if predicate(val) {
			result = append(result, val)
		}
	}
	return result
}

// 映射切片
func mapSlice(slice []int, transform func(int) int) []int {
	result := make([]int, len(slice))
	for i, val := range slice {
		result[i] = transform(val)
	}
	return result
}

// 反转切片
func reverse(slice []int) []int {
	result := make([]int, len(slice))
	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		result[i] = slice[j]
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 过滤偶数
	evens := filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("偶数: %v\n", evens)

	// 过滤奇数
	odds := filter(numbers, func(n int) bool {
		return n%2 != 0
	})
	fmt.Printf("奇数: %v\n", odds)

	// 映射：所有数字乘以2
	doubled := mapSlice(numbers, func(n int) int {
		return n * 2
	})
	fmt.Printf("乘以2: %v\n", doubled)

	// 反转切片
	reversed := reverse(numbers)
	fmt.Printf("反转: %v\n", reversed)

	// 切片排序
	unsorted := []int{5, 2, 8, 1, 9, 3}
	// 使用sort包排序
	fmt.Printf("排序前: %v\n", unsorted)

	// 注意：需要导入 sort 包
	// import "sort"
	sort.Ints(unsorted)
	fmt.Printf("排序后: %v\n", unsorted)
}
