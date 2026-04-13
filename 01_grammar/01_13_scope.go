package grammar

import "fmt"

// 包级变量（全局变量）：在整个 package grammar 内可见。
var scopeGlobalA int = 13

// ScopeDemo 演示 Go 中变量作用域：全局、局部、形参。
func ScopeDemo() {
	// 局部变量：只在当前函数内可见。
	a := 3
	b := 4

	fmt.Printf("global a = %d\n", scopeGlobalA)
	fmt.Printf("ScopeDemo local a = %d, b = %d\n", a, b)

	// 同名遮蔽示例：innerA 局部变量不会影响全局变量。
	innerA := 100
	fmt.Printf("ScopeDemo innerA(local shadow demo) = %d\n", innerA)
	fmt.Printf("global a still = %d\n", scopeGlobalA)

	c := scopeSum(a, b) // a、b 作为实参传给形参
	fmt.Printf("ScopeDemo c = %d\n", c)
}

// scopeSum 的 a、b 是形参（函数局部变量的一种）。
func scopeSum(a, b int) int {
	fmt.Printf("scopeSum param a = %d, b = %d\n", a, b)
	num := a + b // num 是 scopeSum 的局部变量
	return num
}
