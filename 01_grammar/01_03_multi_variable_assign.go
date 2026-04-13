package grammar

import "fmt"

// 多个变量同时初始化
func MultiVariableAssign() {
	fmt.Println("=== Go 多变量同时赋值 ===")

	var a int = 100
	var b int = 200
	fmt.Printf("original value: a: %d, b: %d\n", a, b)
	
	b, a = a, b
	fmt.Printf("changed value: a: %d, b: %d\n", a, b)
}