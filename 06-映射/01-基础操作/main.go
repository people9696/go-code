package main

import "fmt"

func main() {
	// 创建map
	var m1 map[string]int
	fmt.Printf("空map: %v\n", m1)

	// 使用make创建map
	m2 := make(map[string]int)
	m3 := make(map[string]int, 10) // 指定初始容量

	// Map字面量
	m4 := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 8,
	}
	fmt.Printf("m2: %v\n", m2)
	fmt.Printf("m3: %v\n", m3)
	fmt.Printf("m4: %v\n", m4)

	// 添加和修改元素
	m4["grape"] = 12
	m4["apple"] = 19
	fmt.Printf("m4修改后: %v\n", m4)

	// 访问元素
	value := m4["apple"]
	fmt.Printf("苹果的数量是：%d\n", value)

	// 检查键是否存在
	value, exists := m4["pear"]
	if exists {
		fmt.Printf("梨数量：%d\n", value)
	} else {
		fmt.Printf("梨不存在")
	}

	// 删除元素
	delete(m4, "banana")
	fmt.Printf("m4: %v\n", m4)

	//获取map长度
	fmt.Printf("m4的长度：%d\n", len(m4))

	// 遍历map
	fmt.Println("遍历map：")
	for k, v := range m4 {
		fmt.Printf("	%s： %d\n", k, v)
	}
}
