package grammar

import "fmt"

// PtrDemo 演示 Go 指针的基础用法。
func PtrDemo() {
	// 1) 使用 & 取变量地址，得到指针。
	n := 10
	p := &n
	fmt.Printf("n=%d, &n=%p, p=%p\n", n, &n, p)

	// 2) 使用 * 读取指针指向的值（解引用）。
	fmt.Printf("*p=%d\n", *p)

	// 3) 使用 * 修改指针指向的值，会影响原变量。
	*p = 20
	fmt.Printf("after *p=20, n=%d\n", n)

	// 4) 指针作为函数参数，可在函数内部修改外部变量。
	a, b := 1, 2
	swapByValue(&a, &b)
	fmt.Printf("swapByValue -> a=%d, b=%d\n", a, b)

	// 5) 仅交换函数内部的指针变量本身，不会影响外部变量值。
	a, b = 1, 2
	swapPointerOnly(&a, &b)
	fmt.Printf("pointerAddress -> &a=%p, &b=%p\n", &a, &b)
	fmt.Printf("swapPointerOnly -> a=%d, b=%d\n", a, b)
	fmt.Printf("pointerAddress -> &a=%p, &b=%p\n", &a, &b)

	// 6) 使用 new(T) 创建 *T 指针，初始指向类型零值。
	ps := new(string)
	*ps = "Go pointer"
	fmt.Printf("new(string): addr=%p, value=%s\n", ps, *ps)
}

// swapByValue 交换两个指针指向的“值”。
func swapByValue(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

// swapPointerOnly 仅交换形参指针，不会修改调用方变量。
func swapPointerOnly(a, b *int) {
	a, b = b, a
	_, _ = a, b
}
