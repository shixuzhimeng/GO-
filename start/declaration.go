package main

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
// func main() {
// 	var x, y int
// 	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
// }

// 另一个创建变量的方法是调用用内建的new函数
// func main() {
// 	p := new(int)
// 	// p, *int 类型, 指向匿名的 int 变量
// 	fmt.Println(*p) // "0"
// 	*p = 2
// 	// 设置 int 匿名变量的值为 2
// 	fmt.Println(*p) // "2"
// }

// 这里的两个方式是相同的
// func newInt1() *int {
// 	return new(int)
// }
// func newInt2() *int {
// 	var dummy int
// 	return &dummy
// }

// 对于在包一级声明的变量来说，它们
// 的生命周期和整个程序的运行周期是一致的。而相比之下，在局部变量的声明周期则是动态的：从每次
// 创建一个新变量的声明语句开始，直到该变量不再被引用为止
