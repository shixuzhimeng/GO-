package main

import "fmt"

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
// func gcd(x, y int) int {
// 	for y != 0 {
// 		x, y = y, x%y
// 	}
// 	return x
// }

// 隐式的赋值行为
// 可赋值性
// func main() {
// 	//medals := []string{"gold", "silver", "bronze"}
// 	var medals [3]string
// 	medals[0] = "gold"
// 	medals[1] = "silver"
// 	medals[2] = "bronze"

// 	fmt.Println(medals)
// }

// 类型
// type 类型名字 底层类型
type Celsius float64

// 摄氏温度
type Fahrenheit float64 // 华氏温度
const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func fzmain() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	// 虽然它们底层都是 float64，但 Go 将它们视为不同的类型，以防止意外的单位混用
	// fmt.Printf("%g\n", boilingF-FreezingC)
	// compile error: type mismatch
}
