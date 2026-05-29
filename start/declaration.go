package main

import "fmt"

const boilingF = 212.0

// 声明
// func main() {
// 	var f = boilingF
// 	var c = (f - 32) * 5 / 9
// 	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

// }

// func main() {
// 	const freezingF, boilingF = 32.0, 212.0
// 	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
// 	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
// }
// func fToC(f float64) float64 {
// 	return (f - 32) * 5 / 9
// }

// 定义
//var 变量名字 类型 = 表达式

// 简短变量声明  --- 只能定义在函数内部
// func main() {
// 	anim := gif.GIF{LoopCount: nframes}
// 	freq := rand.Float64() * 3.0
// 	t := 0.0
// }

// 指针
//
//	func main() {
//		x := 1
//		p := &x
//		// p, of type *int, points to x
//		fmt.Println(*p) // "1"
//		*p = 2          // equivalent to x = 2
//		fmt.Println(x)  // "2"
//	}

// 任何类型的指针的零值都是nil。如果p != nil测试为真，
// 那么p是指向某个有效变量。指针之间也是可以进行相等测试的，只有当它们指向同一个变量或全部是nil时才相等
func main() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
}
