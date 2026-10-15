package main

import (
	"fmt"
)

// 简单的实现了一个乘法表
// func main() {
// 	num := 9
// 	for i := 1; i <= num; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Printf("%v * %v = %v	", j, i, j*i)
// 		}
// 		fmt.Println()
// 	}
// }

// 实现随机数生成和break跳出循环
// func main() {
// 	// 但是现在的Seed已经不用了，直接换为New了
// 	rand.Seed(time.Now().Unix())
// 	fmt.Println("n", rand.Intn(100)+1)

// 	//
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	// 等价于
// 	// seed := time.Now().UnixNano() // 1. 拿一个种子
// 	// src := rand.NewSource(seed)   // 2. 用种子造一个"随机源"
// 	// r := rand.New(src)            // 3. 用随机源造一个"生成器"

// 	for {
// 		// 生成 [1, 100] 的随机数
// 		n := r.Intn(100) + 1

// 		fmt.Println(n)

// 		// 如果是 99，退出循环（程序结束）
// 		if n == 99 {
// 			fmt.Println("生成到 99，程序退出")
// 			break
// 		}
// 	}
// }

func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}

}

func main() {
	res := fbn(3)
	fmt.Println("res = ", res)
	fmt.Println("res = ", fbn(4))
	fmt.Println("res = ", fbn(5))
	fmt.Println("res = ", fbn(6))
}
