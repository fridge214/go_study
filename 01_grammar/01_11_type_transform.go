package grammar

import "fmt"

// TypeTransformDemo 演示 Go 的类型转换（必须显式转换）。
func TypeTransformDemo() {
	// Go 没有隐式类型转换，必须写成：目标类型(原值)。
	var f float64 = 5.8
	i := int(f) // 浮点转整型会直接去掉小数部分（截断，不是四舍五入）
	fmt.Printf("float64 -> int: %v -> %d\n", f, i)

	// 从大范围整型转小范围整型，可能发生截断。
	var big int32 = 1047483647
	small := int16(big)
	fmt.Printf("int32: 0x%x %d\n", big, big)
	fmt.Printf("int16: 0x%x %d (truncated)\n", small, small)

	// 从小范围转大范围通常安全（值不变）。
	var a int16 = 32000
	b := int32(a)
	fmt.Printf("int16 -> int32: %d -> %d\n", a, b)

	// 不同底层类型不能随意转换，例如 bool 与 int 不能互转（以下仅注释示例）。
	// var ok bool = true
	// fmt.Println(int(ok)) // 编译错误：cannot convert ok (type bool) to type int
	
}
