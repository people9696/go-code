package main

import (
	"fmt"
	"math"
)

// 方法是与特定类型关联的函数，可以为任何类型定义方法。

// 定义结构体
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// 值接受者方法
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// 指针接收者方法
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Cirecle的方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 字符串表示方法
func (r Rectangle) String() string {
	return fmt.Sprintf("矩形[宽:%.2f, 高:%.2f]", r.Width, r.Height)
}
func (c Circle) String() string {
	return fmt.Sprintf("圆形[半径:%.2f]", c.Radius)
}

// 自定义类型方法
type MyInt int

func (mi MyInt) IsEven() bool {
	return mi%2 == 0
}
func (mi MyInt) IsPositive() bool {
	return mi > 0
}
func main() {
	rect := Rectangle{Width: 10, Height: 20}
	circle := Circle{9}

	// 调用方法
	fmt.Printf("%s 面积: %.2f\n", rect, rect.Area())
	fmt.Printf("%s 周长: %.2f\n", rect, rect.Perimeter())

	fmt.Printf("%s 面积: %.2f\n", circle, circle.Area())
	fmt.Printf("%s 周长: %.2f\n", circle, circle.Perimeter())

	// 修改结构体
	fmt.Printf("缩放前: %s\n", rect)
	rect.Scale(2)
	fmt.Printf("缩放后: %s\n", rect)

	// 自定义类型方法
	num := MyInt(42)
	fmt.Printf("%d 是偶数: %v\n", num, num.IsEven())
	fmt.Printf("%d 是正数: %v\n", num, num.IsPositive())

	// 方法值和方法表达式
	areaMethod := rect.Area
	fmt.Printf("使用方法值: %.2f\n", areaMethod())

	scaleMethod := (*Rectangle).Scale
	scaleMethod(&rect, 0.5)
	fmt.Printf("使用方法表达式: %s\n", rect)
}
