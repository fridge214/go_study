package grammar

import "fmt"

// VariableEscapeDemo 演示 Go 变量逃逸的常见场景。
func VariableEscapeDemo() {
	fmt.Println("=== Variable Escape Demo ===")

	// 1) 返回局部变量地址：x 需要在函数外继续存活，通常会逃逸到堆。
	p := newIntPtr()
	fmt.Printf("newIntPtr -> *p=%d\n", *p)

	// 2) 只在函数内计算并返回值：通常可在栈上完成。
	v := localCompute()
	fmt.Printf("localCompute -> %d\n", v)

	// 3) 值放入 interface（装箱）：在部分场景可能引起逃逸。
	boxed := toInterface()
	fmt.Printf("toInterface -> type=%T, value=%v\n", boxed, boxed)

	// 4) 闭包捕获外部变量：被捕获变量常需要延长生命周期。
	adder := makeAdder(100)
	fmt.Printf("makeAdder -> %d\n", adder(23))
}

// newIntPtr 返回局部变量地址，典型逃逸示例。
func newIntPtr() *int {
	x := 10
	return &x
}

// localCompute 仅返回值，局部变量通常不逃逸。
func localCompute() int {
	x := 10
	return x + 1
}

// toInterface 演示装箱到接口。
func toInterface() any {
	x := 42
	return x
}

// makeAdder 返回闭包，闭包会捕获 base。
func makeAdder(base int) func(int) int {
	return func(v int) int {
		return base + v
	}
}
