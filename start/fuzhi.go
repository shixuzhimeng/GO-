package main

// type Person struct {
// 	name string
// }

// func main() {
// 	x := 10
// 	p := new(bool)
// 	x = 1     // 命令变量的赋值
// 	*p = true // 通过指针间接赋值
// 	person := Person{}
// 	person.name = "bob" // 结构体字段赋值
// 	count := make(map[int]int)
// 	scale := 2
// 	count[x] = count[x] * scale // 数组、slice或map的元素赋值
// 	fmt.Println()
// }

// 原组赋值
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
