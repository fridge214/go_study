package collector

import "fmt"

// RangeDemo 演示切片上 range 的常见用法与注意点。
func RangeDemo() {
	slice := []int{10, 20, 30, 40}

	// 1) range 同时返回索引和值。
	for index, value := range slice {
		fmt.Printf("index=%d value=%d\n", index, value)
	}

	// 2) value 是元素副本，不是原元素本身。
	for index, value := range slice {
		fmt.Printf("value=%d valueAddr=%p elemAddr=%p\n", value, &value, &slice[index])
	}

	// 3) 不需要索引时，用 _ 忽略。
	for _, value := range slice {
		fmt.Printf("only value=%d\n", value)
	}

	// 4) 如果要从指定位置开始遍历，用传统 for。
	for index := 2; index < len(slice); index++ {
		fmt.Printf("classic for index=%d value=%d\n", index, slice[index])
	}
}
