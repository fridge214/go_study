package collector

import "fmt"

// ArrayDemo 演示 Go 数组的声明、初始化、遍历和比较。
func ArrayDemo() {
	// 1) 数组长度是类型的一部分，默认元素值为零值。
	var a [3]int
	fmt.Printf("a=%v, len=%d\n", a, len(a))
	fmt.Printf("a[0]=%d, a[last]=%d\n", a[0], a[len(a)-1])

	// 2) 指定长度初始化；未提供的元素会补零值。
	b := [3]int{1, 2}
	fmt.Printf("b=%v\n", b)

	// 3) 使用 ... 让编译器根据初始值数量推导数组长度。
	c := [...]int{1, 2, 3}
	fmt.Printf("c=%v, type=%T\n", c, c)

	// 4) 遍历数组：range 同时拿到索引和值。
	team := [3]string{"hammer", "soldier", "mum"}
	for i, v := range team {
		fmt.Printf("team[%d]=%s\n", i, v)
	}

	// 5) 数组比较：类型相同（长度和元素类型都相同）才能比较。
	x := [2]int{1, 2}
	y := [...]int{1, 2}
	z := [2]int{1, 3}
	fmt.Printf("x==y -> %t, x==z -> %t\n", x == y, x == z)

	// 说明：
	// [3]int 和 [4]int 是不同类型，不能直接互相赋值或比较。
}
