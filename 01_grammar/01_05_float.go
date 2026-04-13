package grammar

import (
	"fmt"
	"math"
)

// FloatDemo 演示 Go 浮点类型（float32 / float64）的基础用法
func FloatDemo() {
	// Go 提供两种浮点类型：float32 和 float64（遵循 IEEE 754 标准）。
	// 通常优先使用 float64，因为精度更高。
	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793
	fmt.Printf("f32 = %f, f64 = %f\n", f32, f64)

	// math 包中提供了浮点最大值常量。
	fmt.Printf("math.MaxFloat32 = %e\n", math.MaxFloat32)
	fmt.Printf("math.MaxFloat64 = %e\n", math.MaxFloat64)

	// float32 只有大约 6~7 位十进制有效数字，可能出现精度问题。
	// 下面这个例子会输出 true：因为 16777216 是 float32 可精确表示的边界附近，+1 后无法区分。
	var boundary float32 = 16777216 // 1 << 24
	fmt.Printf("boundary == boundary+1 ? %t\n", boundary == boundary+1)

	// 浮点字面量可以写成科学计数法，适合很大或很小的数。
	const avogadro = 6.02214129e23
	const planck = 6.62606957e-34
	fmt.Printf("Avogadro = %e, Planck = %e\n", avogadro, planck)

	// Printf 中使用 %f 控制浮点输出格式，%.2f 表示保留 2 位小数。
	fmt.Printf("Pi default = %f\n", math.Pi)
	fmt.Printf("Pi 2 digits = %.2f\n", math.Pi)
}
