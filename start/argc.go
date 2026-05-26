package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//	func main() {
//		var s, sep string
//		for i := 1; i < len(os.Args); i++ {
//			s += sep + os.Args[i]
//			sep = " "
//		}
//		fmt.Println(s)
//	}

//  输出Argv[0]
// func main() {
// 	var s, sep string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// 输出value和index，每个value和index显示一行
// func main() {
// 	for i, arg := range os.Args {
// 		fmt.Printf("index: %d, value: %s\n", i, arg)
// 	}
// }

// 上手实践前面提到的strings.Join和直接Println，并观察输出结果的区别
// strings.Join 生成一个干净、用空格分隔的字符串，一次输出；适合需要特定分隔符的格式化输出
func mainArgc() {
	args := os.Args[1:] // 忽略程序名

	// 方法1：strings.Join
	start := time.Now()
	result := strings.Join(args, " ")
	fmt.Println("Join result:", result)
	fmt.Printf("Join took: %v\n", time.Since(start))

	// 方法2：多次Println（模拟直接拼接或逐次打印）
	start = time.Now()
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
	fmt.Println()
	fmt.Printf("Loop print took: %v\n", time.Since(start))
}
