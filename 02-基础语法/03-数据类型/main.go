package main

import "fmt"

func main() {
	// 整数类型
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	// 无符号整数
	var ui8 uint8 = 255
	var ui16 uint16 = 65535

	// 浮点数
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	// 布尔类型
	var isTrue bool = true
	var isFalse bool = false

	// 字符串
	var str string = "Hello, 世界"

	// 字符类型
	var b byte = 'A' // 等同于 uint8
	var r rune = '中' // 等同于 int32

	fmt.Printf("int8: %d, uint8: %d\n", i8, ui8)
	fmt.Printf("int16: %d, uint16: %d\nint32: %d\nint64: %d\n", i16, ui16, i32, i64)
	fmt.Printf("float32: %.2f, float64: %.15f\n", f32, f64)
	fmt.Printf("bool: %v, %v\n", isTrue, isFalse)
	fmt.Printf("string: %s\n", str)
	fmt.Printf("byte: %c, rune: %c\n", b, r)
}
