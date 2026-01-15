// 切片是动态数组，是Go中最常用的数据结构。
package main

import "fmt"

func main() {
	// 创建切片
	var slice []int
	fmt.Printf("空切片：%v, 长度：%d\n", slice, len(slice))

	// 使用make创建切片
	slice2 := make([]int, 5)     // 长度为5，容量为5
	slice3 := make([]int, 3, 10) // 长度为3，容量为10

	fmt.Printf("slice2: %v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3: %v, len: %d, cap: %d\n", slice3, len(slice3), cap(slice3))

	// 切片字面量
	numbers := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("数字切片: %v\n", numbers)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice4 := arr[2:7] //左闭右开
	fmt.Printf("从数组切片: %v， 切片长度：%d\n", slice4, len(slice4))

	// 追加元素
	slice5 := []int{1, 2, 3}
	fmt.Printf("追加前: %v\n", slice5)
	slice5 = append(slice5, 4, 5, 6)
	fmt.Printf("追加后: %v\n", slice5)

	// 追加切片
	slice6 := []int{7, 8, 9}
	slice5 = append(slice5, slice6...)
	fmt.Printf("追加切片后: %v\n", slice5)

	// 复制切片
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(source))
	copy(dest, source)
	fmt.Printf("复制后: %v\n", dest)

	// 删除元素
	slice7 := []int{1, 2, 3, 4, 5}
	// 删除索引为2的元素
	slice7 = append(slice7[:2], slice7[3:]...)
	fmt.Printf("删除元素后: %v\n", slice7)

	// 遍历切片
	fmt.Println("遍历切片:")
	for i, val := range numbers {
		fmt.Printf("  索引 %d: %d\n", i, val)
	}
}
