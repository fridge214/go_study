package grammar

import "fmt"

// NewInt 是“新定义类型”，底层类型是 int，但它本身是新类型。
type NewInt int

// IntAlias 是“类型别名”，本质上就是 int。
type IntAlias = int

// AliasDemo 演示类型定义与类型别名的区别。
func AliasDemo() {
	var a NewInt = 10
	var b IntAlias = 10
	var c int = 10

	// %T 可看到：NewInt 会保留自己的类型名；IntAlias 显示为 int。
	fmt.Printf("a type=%T, value=%v\n", a, a)
	fmt.Printf("b type=%T, value=%v\n", b, b)
	fmt.Printf("c type=%T, value=%v\n", c, c)

	// 新定义类型与原类型不同，通常需要显式转换。
	c = int(a)
	fmt.Printf("after convert NewInt -> int, c=%d\n", c)

	// 别名与原类型是同一类型，可直接赋值。
	c = b
	fmt.Printf("assign alias to int directly, c=%d\n", c)

	// 说明：
	// 1) type T U      -> 定义新类型 T（与 U 不同）
	// 2) type T = U    -> 定义别名 T（与 U 相同）
	// 3) 非本地类型的别名不能在当前包上新增方法（会编译错误）
}
