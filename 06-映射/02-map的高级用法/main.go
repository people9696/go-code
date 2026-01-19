package main

// Map可以存储复杂的数据类型，包括结构体、切片等。
type Student struct {
	Name  string
	Age   int
	Grade map[string]float64
}

func mian() {
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

}
