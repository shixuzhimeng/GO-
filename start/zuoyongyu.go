package main

import "fmt"

// func f() int {
// 	return 5
// }

// func g(x int) int {
// 	return x + 1
// }

// func main() {
// 	f := "f"
// 	fmt.Println(f) // 输出 "f"；局部变量 f 遮蔽了包级别的函数 f
// 	fmt.Println(g) // 输出函数 g 的内存地址；因为局部变量 f 的遮蔽，这里的 g 依然指向包级别的函数 g<websource>source_group_web_1</websource>

// 	var x, y int

// 	if x = f(); x == 0 { // 使用 = 赋值给外部声明的 x，而不是用 := 重新声明
// 		fmt.Println(x)
// 	} else if y = g(x); x == y { // 使用 = 赋值给外部声明的 y
// 		fmt.Println(x, y)
// 	} else {
// 		fmt.Println(x, y)
// 	}
// 	// 现在 x 和 y 在这里是可见的
// 	fmt.Println(x, y)
// }

//	func printHello() {
//		x := "hello"
//		for _, r := range x { // 使用不同的变量名 r 接收 range 的 rune，避免遮蔽外部的 x
//			upper := r + 'A' - 'a' // 使用新变量 upper 存储转换后的大写字母
//			fmt.Printf("%c", upper)
//		}
//		fmt.Println() // 换行
//	}

// func main() {
// 	x := "hello"
// 	for _, x := range x {
// 		x := x + 'A' - 'a'
// 		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
// 	}
// }

func main() {
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0";
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
}

const (
	deadbeef = 0xdeadbeef        // untyped int with value 3735928559
	a        = uint32(deadbeef)  // uint32 with value 3735928559
	b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
	c        = float64(deadbeef) // float64 with value 3735928559 (exact)
	d        = int32(deadbeef)   // compile error: constant overflows int32
	e        = float64(1e309)    // compile error: constant overflows float64
	f        = uint(-1)          // compile error: constant underflows uint
)
