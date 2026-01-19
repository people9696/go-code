package main

import "fmt"

// 结构体是Go语言中的自定义数据类型，用于组合多个字段。
// 定义结构体
type Person struct {
	Name string
	Age  int
}

// 带有更多字段的结构体
type Student struct {
	Name     string
	Age      int
	Grade    float64
	Subjects []string
	Address  *Address // 指针
}
type Address struct {
	City    string
	Country string
}

func main() {
	// 创建结构体实例
	var p1 Person
	p1.Name = "Alice"
	p1.Age = 25

	// 结构体字面量
	p2 := Person{
		Name: "Bob",
		Age:  16,
	}
	// 简写形式（按顺序）
	p3 := Person{"Charles", 18}

	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p2: %+v\n", p2)
	fmt.Printf("p3: %+v\n", p3)

	// 访问结构体字段
	fmt.Printf("%s 的年龄是 %d\n", p1.Name, p1.Age)
	// 创建复杂结构体
	student := Student{
		Name:     "David",
		Age:      19,
		Grade:    89.7,
		Subjects: []string{"Math", "English", "Science"},
		Address: &Address{
			City:    "Beijing",
			Country: "China",
		},
	}

	fmt.Printf("学生信息: %+v\n", student)
	fmt.Printf("城市: %s\n", student.Address.City)

	// 结构体比较
	p4 := Person{Name: "Alice", Age: 25}
	if p1 == p4 {
		fmt.Println("p1 和 p4 相等")
	}

	// 结构体数组
	people := []Person{
		{Name: "Eve", Age: 28},
		{Name: "Frank", Age: 32},
		{Name: "Grace", Age: 27},
	}
	fmt.Println("所有人:")
	for _, person := range people {
		fmt.Printf("  %s (%d岁)\n", person.Name, person.Age)
	}
}
