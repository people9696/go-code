package main

import (
	"fmt"
)

// Go的switch语句非常灵活，可以没有表达式，自动break。
func main() {
	day := "Monday"

	// 基本switch语句
	switch day {
	case "Monday":
		fmt.Println("今天是星期一")
	case "Tuesday":
		fmt.Println("今天是星期二")
	case "Wednesday", "Thursday", "Friday":
		fmt.Println("今天是工作日")
	case "Saturday", "Sunday":
		fmt.Println("今天是周末")
	default:
		fmt.Println("无效的日期")
	}

	// 没有表达式的switch（相当于if-else）
	score := 88
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 带条件的switch
	switch x := 10; {
	case x > 0 && x < 10:
		fmt.Println("个位数正数")
	case x >= 10 && x < 100:
		fmt.Println("两位数正数")
	default:
		fmt.Println("其他数字")
	}

	// fallthrough示例(继续执行下一个case)
	num := 2
	switch num {
	case 1:
		fmt.Println("一")
		fallthrough
	case 2:
		fmt.Println("二")
		fallthrough
	case 3:
		fmt.Println("三")
	default:
		fmt.Println("其他数字")
	}

}
