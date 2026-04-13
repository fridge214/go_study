package collector

import "fmt"

// MultiDimensionArrayDemo 演示 Go 多维数组（以二维数组为主）。
func MultiDimensionArrayDemo() {
	// 1) 声明一个 2x2 的二维数组，默认值都是 0。
	var a [2][2]int
	fmt.Printf("default a = %v\n", a)

	// 2) 使用索引为每个元素赋值。
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 40
	fmt.Printf("after set a = %v\n", a)

	// 3) 使用字面量初始化二维数组。
	b := [2][2]int{{1, 2}, {3, 4}}
	fmt.Printf("literal b = %v\n", b)

	// 4) 按索引指定初始化（未指定位置使用零值）。
	c := [4][2]int{
		1: {20, 21},
		3: {40, 41},
	}
	fmt.Printf("indexed literal c = %v\n", c)

	// 5) 遍历二维数组：外层行、内层列。
	for i, row := range b {
		for j, v := range row {
			fmt.Printf("b[%d][%d]=%d ", i, j, v)
		}
		fmt.Println()
	}

	// 6) 同类型二维数组可以直接赋值（值拷贝）。
	var d [2][2]int
	d = b
	fmt.Printf("copy d = %v\n", d)

	// 7) 可以单独取出某一行（该行本质是一维数组）。
	row1 := d[1] // type: [2]int
	val := d[1][0]
	fmt.Printf("row1=%v, val=%d\n", row1, val)

	// 说明：多维数组类型由“每一维长度 + 元素类型”共同决定。
	// 例如 [2][2]int 与 [3][2]int 是不同类型。
}
