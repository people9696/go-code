package main

import "fmt"

func main() {
	age := 25

	// 基本if语句
	if age >= 18 {
		fmt.Println("成年人")

	} else {
		fmt.Println("未成年")
	}

	// if-else
	if age < 13 {
		fmt.Println("儿童")
	} else if age < 18 {
		fmt.Println("青少年")

	} else if age < 60 {
		fmt.Println("成年人")
	} else {
		fmt.Println("老年人")
	}

	// 在if前执行语句
	if x := 10; x > 5 {
		fmt.Printf("%d 大于5.\n", x)
	}

	// 多个条件
	score := 85
	if score >= 90 && score <= 100 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
