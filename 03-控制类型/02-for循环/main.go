package main

import "fmt"

//  go只有for循环，没有while和do-while循环

func main() {
	// 1. 基本for循环
	fmt.Println("基本for循环:")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// 2. 类似while循环
	fmt.Println("类似while循环:")
	j := 0
	for j < 5 {
		fmt.Println("j =", j)
		j++
	}

	// 3. 无限循环
	fmt.Println("无限循环:")
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println("k =", k)
		k++
	}

	// 4. range循环遍历数组和切片
	fmt.Println("range循环:")
	arr := []string{"apple", "banana", "cherry"}
	for index, value := range arr {
		fmt.Println("index =", index, "value =", value)
	}

	// range循环遍历字符串
	fmt.Println("range循环字符串:")
	str := "hello"
	for index, char := range str {
		fmt.Println("index =", index, "char =", string(char))
	}

	// 5. 只获取索引或值
	// 只获取索引
	fmt.Println("只获取索引:")
	for index := range arr {
		fmt.Println("index =", index)
	}
	// 只获取值
	fmt.Println("只获取值:")
	for _, value := range arr {
		fmt.Println("value =", value)
	}

	// 6. 映射遍历
	fmt.Println("映射遍历:")
	mapping := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range mapping {
		fmt.Println("key =", key, "value =", value)
	}

}
