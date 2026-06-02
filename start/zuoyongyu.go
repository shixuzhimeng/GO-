package main

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

//	func main() {
//		x := "hello"
//		for _, x := range x {
//			x := x + 'A' - 'a'
//			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
//		}
//	}

// const (
// 	deadbeef = 0xdeadbeef        // untyped int with value 3735928559
// 	a        = uint32(deadbeef)  // uint32 with value 3735928559
// 	b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
// 	c        = float64(deadbeef) // float64 with value 3735928559 (exact)
// 	d        = int32(deadbeef)   // compile error: constant overflows int32
// 	e        = float64(1e309)    // compile error: constant overflows float64
// 	f        = uint(-1)          // compile error: constant underflows uint
// )

// func main() {
// 	var f float64 = 212
// 	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
// 	fmt.Println(5 / 9 * (f - 32))     // "0";
// 	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
// }

//	func main() {
//		i := 0
//		r := '\000'
//		f := 0.0
//		fmt.Println(i)
//		fmt.Println(r)
//		fmt.Println(f)
//	}
//

// 内置的append函数用于向slice追加元素
// func appendInt(x []int, y int) []int {
// 	var z []int
// 	zlen := len(x) + 1
// 	if zlen <= cap(x) {
// 		z = x[:zlen]
// 	} else {

// 		zcap := zlen
// 		if zcap < 2*len(x) {
// 			zcap = 2 * len(x)
// 		}
// 		z = make([]int, zlen, zcap)
// 		copy(z, x)
// 	}
// 	z[len(x)] = y
// 	return z
// }

// func main() {
// 	var x, y []int
// 	for i := 0; i < 10; i++ {
// 		y = appendInt(x, i)
// 		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
// 		x = y
// 	}
// }

func appendInt(x []int, y ...int) []int {
	var z []int
	copy(z[len(x):], y)
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
