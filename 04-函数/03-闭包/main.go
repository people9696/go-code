package main

import "fmt"

// 闭包最简单的例子：函数里用到了外面的变量
// func main() {
// 	x := 10
// 	f := func() {
// 		fmt.Println(x)
// 	}
// 	f()
// }

// 闭包经典用法：外层函数返回内层函数
// func makeCounter() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func main() {
// 	c1 := makeCounter()
// 	fmt.Println(c1()) // 1
// 	fmt.Println(c1()) // 2
// }

// 闭包捕获的是"变量"，不是"值"（非常重要）
// func main() {
// 	x := 0
// 	f1 := func() {
// 		fmt.Println("f1: ", x)
// 	}
// 	x = 100
// 	f2 := func() { fmt.Println("f2: ", x) }

// 	f1() // 100 f1函数用到了函数外部的变量x，捕获的一直是变量，后来x变成100，f1输出也是100
// 	f2() // 100

// }

// 返回一个闭包
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 记忆函数
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}
		result := f(n)
		cache[n] = result
		return result
	}
}

// 斐波那契数列
func fibnacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibnacci(n-1) + fibnacci(n-2)
}

func main() {
	// 使用闭包
	counter1 := counter()
	fmt.Println("counter1: ", counter1())
	fmt.Println("counter1: ", counter1())
	fmt.Println("counter1: ", counter1())

	counter2 := counter()
	fmt.Println("counter2: ", counter2())
	fmt.Println("counter2: ", counter2())
	fmt.Println("counter2: ", counter2())

	// 使用记忆化.自动记住算过的结果，下次同样输入直接返回缓存
	memoFib := memoize(fibnacci)

	fmt.Println(memoFib(10)) // 第一次计算
	fmt.Println(memoFib(10)) // 第二次从缓存读取
	fmt.Println(memoFib(20)) // 新的计算

	fmt.Println(fibnacci(30)) // 直接使用斐波那契函数
}
