package grammar

import "fmt"

// getData 模拟一个返回两个值的函数
func getData() (int, int) {
	return 100, 200
}

// AnonymousVariable 演示匿名变量（空白标识符 _）的用法
func AnonymousVariable() {
	fmt.Println("=== Go 匿名变量示例 ===")

	// 只接收第一个返回值，第二个用 _ 丢弃
	a, _ := getData()

	// 只接收第二个返回值，第一个用 _ 丢弃
	_, b := getData()

	fmt.Printf("a = %d, b = %d\n", a, b)
}
