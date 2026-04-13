package grammar

import "fmt"

// BoolDemo 演示 Go 的布尔类型（bool）和常见逻辑运算。
func BoolDemo() {
	// bool 只有两个值：true / false。
	var t bool = true
	var f bool = false
	fmt.Printf("t=%t, f=%t\n", t, f)

	// 比较运算会产生 bool 结果。
	a := 10
	fmt.Printf("a==5 -> %t\n", a == 5)
	fmt.Printf("a==10 -> %t\n", a == 10)
	fmt.Printf("a!=5 -> %t\n", a != 5)

	// ! 是逻辑非，&& 是逻辑与，|| 是逻辑或。
	fmt.Printf("!true -> %t\n", !true)
	fmt.Printf("true && false -> %t\n", true && false)
	fmt.Printf("true || false -> %t\n", true || false)

	// && / || 都有短路行为：左侧已能决定结果时，右侧不会再计算。
	s := ""
	// 这里安全：当 s == "" 时，右侧 s[0] 不会执行，避免越界。
	isXPrefix := s != "" && s[0] == 'x'
	fmt.Printf("short-circuit safe check -> %t\n", isXPrefix)

	// && 的优先级高于 ||，下面表达式可用于判断字符是否是字母或数字。
	c := 'G'
	isAlphaNum := 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9'
	fmt.Printf("'%c' is alnum -> %t\n", c, isAlphaNum)

	// bool 不会和数字自动互转；通常通过 if 显式转换。
	fmt.Printf("btoi(true)=%d, btoi(false)=%d\n", btoi(true), btoi(false))
	fmt.Printf("itob(0)=%t, itob(8)=%t\n", itob(0), itob(8))
}

// btoi: true -> 1, false -> 0。
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob: 非 0 -> true，0 -> false。
func itob(i int) bool {
	return i != 0
}
