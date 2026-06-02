package main

// type Weekday int

// const (
// 	Sunday Weekday = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// )
// iota 每定义一个常量，iota的值就会加一

// type Flags uint

// const (
// 	FlagUp           Flags = 1 << iota // is up
// 	FlagBroadcast                      // supports broadcast access capability
// 	FlagLoopback                       // is a loopback interface
// 	FlagPointToPoint                   // belongs to a point-to-point link
// 	FlagMulticast                      // supports multicast access capability
// )

// 随着iota的递增，每个常量对应表达式1 << iota，是连续的2的幂，分别对应一个bit位置

// const (
// 	a = 1
// 	b
// 	c = 2
// 	d
// )

// func main() {
// 	fmt.Println(a, b, c, d) // "1 1 2 2"
// }

// 如果一个常量声明没有提供表达式，它会重复使用前一个非空常量声明的表达式

// const (
// 	_   = 1 << (10 * iota)
// 	KiB // 1024
// 	MiB // 1048576
// 	GiB // 1073741824
// 	TiB // 1099511627776
// 	PiB // 1125899906842624
// 	EiB // 1152921504606846976
// 	ZiB // 1180591620717411303424
// 	YiB // 1208925819614629174706176
// )

// 不过iota常量生成规则也有其局限性。
// 例如，它并不能用于产生1000的幂（KB、MB等），因为Go语言并没有计算幂的运算符。

// 练习：编写KB、MB的常量声明，然后扩展到YB
const (
	_  = 1 << (10 * iota) // iota = 0, 忽略第一个值 1<<0 = 1
	KB                    // 1 << (10*1) = 2^10 = 1024
	MB                    // 2^20 = 1048576
	GB                    // 2^30 = 1073741824
	TB                    // 2^40
	PB                    // 2^50
	EB                    // 2^60
	ZB                    // 2^70 从这里开始会数据溢出
	YB                    // 2^80
)
