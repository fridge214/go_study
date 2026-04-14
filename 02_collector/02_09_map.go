package collector

import "fmt"

// MapDemo 演示 map 的声明、初始化、访问、删除和引用语义。
func MapDemo() {
	// 1) 字面量创建 map。
	mapLit := map[string]int{"one": 1, "two": 2}
	fmt.Printf("mapLit=%v, len=%d\n", mapLit, len(mapLit))

	// 2) make 创建 map（可指定初始容量）。
	mapCreated := make(map[string]float64, 4)
	mapCreated["pi"] = 3.14159
	mapCreated["e"] = 2.71828
	fmt.Printf("mapCreated=%v, len=%d\n", mapCreated, len(mapCreated))

	// 3) map 赋值是引用语义：两个变量指向同一底层结构。
	mapAssigned := mapLit
	mapAssigned["two"] = 3
	fmt.Printf("after assigned write, mapLit[\"two\"]=%d\n", mapLit["two"])

	// 4) 读取不存在的 key：返回 value 类型零值，可用 ok 判断是否存在。
	v1 := mapLit["ten"]
	v2, ok := mapLit["ten"]
	fmt.Printf("mapLit[\"ten\"]=%d\n", v1)
	fmt.Printf("mapLit[\"ten\"] => value=%d, ok=%t\n", v2, ok)

	// 5) 删除 key。
	delete(mapLit, "one")
	fmt.Printf("after delete one, mapLit=%v\n", mapLit)

	// 6) 以切片作为 map 的 value。
	children := make(map[int][]int)
	children[1] = []int{11, 12}
	children[2] = append(children[2], 21, 22)
	fmt.Printf("children=%v\n", children)
}
