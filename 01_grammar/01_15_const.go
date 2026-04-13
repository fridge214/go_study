package grammar

import (
	"fmt"
	"math"
)

// Weekday 使用 iota 定义一组递增常量（类似枚举）。
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 常量可批量声明，未写表达式时会复用上一行表达式。
const (
	constA = 1
	constB
	constC = 2
	constD
)

// ConstDemo 演示 Go 常量的常见用法。
func ConstDemo() {
	// 1) 常量使用 const 定义，值在编译期确定。
	const pi = 3.14159
	const appName string = "go_study"
	fmt.Printf("pi=%v, appName=%s\n", pi, appName)

	// 2) 常量可以是常量表达式。
	const x = 2 / 3 // 编译期求值，结果是 0（整数除法）
	fmt.Printf("const expression 2/3 = %d\n", x)

	// 3) 批量常量与表达式复用。
	fmt.Printf("constA=%d constB=%d constC=%d constD=%d\n", constA, constB, constC, constD)

	// 4) iota 自动递增，常用于状态/枚举值。
	fmt.Printf("Weekday: Sunday=%d Monday=%d Tuesday=%d\n", Sunday, Monday, Tuesday)

	// 5) 无类型常量可在需要时适配目标类型（如 math.Pi）。
	var f32 float32 = math.Pi
	var f64 float64 = math.Pi
	var c128 complex128 = math.Pi
	fmt.Printf("untyped const convert: f32=%.4f f64=%.6f c128=%v\n", f32, f64, c128)
}
