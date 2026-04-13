package grammar

import (
	"fmt"
	"math/cmplx"
)

// ComplexDemo 演示 Go 复数类型与基本运算
func ComplexDemo() {
	// Go 提供两种复数类型：complex64 和 complex128，默认推荐使用 complex128。
	// 复数由实部和虚部组成：RE + IMi。
	var x complex128 = complex(1, 2) // 1 + 2i
	y := complex(3, 4)               // 简写形式，等价于 var y complex128 = complex(3, 4)

	// real(z) 获取实部，imag(z) 获取虚部。
	fmt.Printf("x = %v, real(x) = %.1f, imag(x) = %.1f\n", x, real(x), imag(x))
	fmt.Printf("y = %v, real(y) = %.1f, imag(y) = %.1f\n", y, real(y), imag(y))

	// 复数基本运算法则示例。
	// (a+bi) + (c+di) = (a+c) + (b+d)i
	// (a+bi) - (c+di) = (a-c) + (b-d)i
	// (a+bi) * (c+di) = (ac-bd) + (ad+bc)i
	// (a+bi) / (c+di) = [(ac+bd) + (bc-ad)i] / (c^2+d^2)
	sum := x + y
	diff := x - y
	prod := x * y
	quot := x / y
	fmt.Printf("x + y = %v\n", sum)
	fmt.Printf("x - y = %v\n", diff)
	fmt.Printf("x * y = %v\n", prod)
	fmt.Printf("x / y = %v\n", quot)

	// 复数支持 == / != 比较：实部和虚部都相等时才相等。
	fmt.Printf("x == y ? %t\n", x == y)
	fmt.Printf("x == complex(1, 2) ? %t\n", x == complex(1, 2))

	// math/cmplx 包提供复数常用函数，例如复数模长 Abs。
	fmt.Printf("|x| = %.4f\n", cmplx.Abs(x))
}
