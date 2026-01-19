package main

import "fmt"

// Map可以存储复杂的数据类型，包括结构体、切片等。
type Student struct {
	Name  string
	Age   int
	Grade map[string]float64
}

func main() {
	// map存储结构体
	students := make(map[string]Student)

	students["Alice"] = Student{
		Name: "Alice",
		Age:  14,
		Grade: map[string]float64{
			"Math":    95.5,
			"English": 93.5,
			"Science": 89.0,
		},
	}
	students["Bob"] = Student{
		Name: "Bob",
		Age:  15,
		Grade: map[string]float64{
			"Math":    99.5,
			"English": 98,
			"Science": 88.5,
		},
	}

	// 访问嵌套map
	alice := students["Alice"]
	fmt.Printf("Alice的英语成绩：%.2f.\n", alice.Grade["English"])

	// Map存储切片
	phoneBook := make(map[string][]string)
	phoneBook["Alice"] = []string{"123-456", "789-456"}
	phoneBook["Bob"] = []string{"456-789"}

	fmt.Printf("Alice的所有号码是：%v.\n", phoneBook["Alice"])

	// Map作为集合使用
	visited := make(map[string]bool)
	cities := []string{"Beijing", "Shanghai", "Guangzhou", "Shenzhen"}
	for _, city := range cities {
		if !visited[city] {
			visited[city] = true
			fmt.Printf("首次访问：%s.\n", city)
		}

	}
	fmt.Printf("visited: %v.\n", visited)

	// 统计单词频率
	text := "hello world hello go world"
	words := make(map[string]int)
	for _, word := range []string{"hello", "world", "hello", "go", "world"} {
		words[word]++
	}
	fmt.Println("Text:", text)
	fmt.Println("单词频率：")
	for word, count := range words {
		fmt.Printf("  %s: %d\n", word, count)
	}
}
