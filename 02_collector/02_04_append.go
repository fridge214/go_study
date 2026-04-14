package collector

import "fmt"

// AppendDemo 演示 append 的常见用法与切片扩容行为。
func AppendDemo() {
	// 1) 末尾追加：单个元素、多个元素、追加另一个切片。
	var a []int
	a = append(a, 1)
	a = append(a, 2, 3, 4)
	a = append(a, []int{5, 6, 7}...)
	fmt.Printf("append result: %v\n", a)

	// 2) 观察扩容：len 增长，cap 按策略扩展，底层数组地址可能变化。
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("i=%d len=%d cap=%d ptr=%p values=%v\n",
			i, len(numbers), cap(numbers), sliceDataPtr(numbers), numbers)
	}

	// 3) 头部添加元素（通常会触发拷贝，性能弱于尾部追加）。
	head := []int{1, 2, 3}
	head = append([]int{0}, head...)
	head = append([]int{-2, -1}, head...)
	fmt.Printf("prepend result: %v\n", head)

	// 4) 在中间插入元素/切片（链式 append）。
	insert := []int{1, 2, 5, 6}
	i := 2
	x := 3
	insert = append(insert[:i], append([]int{x}, insert[i:]...)...)
	insert = append(insert[:i+1], append([]int{4}, insert[i+1:]...)...)
	fmt.Printf("insert one by one: %v\n", insert)

	insert2 := []int{1, 2, 6, 7}
	insert2 = append(insert2[:2], append([]int{3, 4, 5}, insert2[2:]...)...)
	fmt.Printf("insert slice: %v\n", insert2)
}

func sliceDataPtr(s []int) *int {
	if len(s) == 0 {
		return nil
	}
	return &s[0]
}
