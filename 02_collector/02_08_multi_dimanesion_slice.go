package collector

import "fmt"

// MultiDimanesionSliceDemo 演示二维切片的声明、遍历和追加元素。
func MultiDimanesionSliceDemo() {
	// 1) 声明并初始化二维切片（切片的切片）。
	slice := [][]int{{10}, {100, 200}}
	fmt.Printf("init slice=%v\n", slice)
	fmt.Printf("outer len=%d cap=%d\n", len(slice), cap(slice))
	fmt.Printf("inner[0] len=%d cap=%d, inner[1] len=%d cap=%d\n",
		len(slice[0]), cap(slice[0]), len(slice[1]), cap(slice[1]))

	// 2) 访问二维切片元素。
	fmt.Printf("slice[0][0]=%d, slice[1][1]=%d\n", slice[0][0], slice[1][1])

	// 3) 给某个内层切片追加元素。
	slice[0] = append(slice[0], 20)
	slice[1] = append(slice[1], 300, 400)
	fmt.Printf("after append=%v\n", slice)

	// 4) 遍历二维切片。
	for i, row := range slice {
		for j, v := range row {
			fmt.Printf("slice[%d][%d]=%d ", i, j, v)
		}
		fmt.Println()
	}
}
