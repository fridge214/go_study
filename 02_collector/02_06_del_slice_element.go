package collector

import "fmt"

// DelSliceElementDemo 演示从切片头部、中间、尾部删除元素的常见写法。
func DelSliceElementDemo() {
	// 1) 从头部删除
	head1 := []int{1, 2, 3, 4, 5}
	head1 = head1[1:] // 删除头部 1 个
	fmt.Printf("head remove by slice: %v\n", head1)

	head2 := []int{1, 2, 3, 4, 5}
	head2 = append(head2[:0], head2[2:]...) // 删除头部 2 个
	fmt.Printf("head remove by append: %v\n", head2)

	head3 := []int{1, 2, 3, 4, 5}
	n := copy(head3, head3[2:]) // 删除头部 2 个
	head3 = head3[:n]
	fmt.Printf("head remove by copy: %v\n", head3)

	// 2) 从中间删除
	mid1 := []string{"a", "b", "c", "d", "e"}
	i := 2 // 删除索引 2 的 "c"
	mid1 = append(mid1[:i], mid1[i+1:]...)
	fmt.Printf("mid remove by append: %v\n", mid1)

	mid2 := []string{"a", "b", "c", "d", "e"}
	i = 1 // 删除索引 1 开始的 2 个元素: b,c
	n2 := i + copy(mid2[i:], mid2[i+2:])
	mid2 = mid2[:n2]
	fmt.Printf("mid remove by copy: %v\n", mid2)

	// 3) 从尾部删除
	tail := []int{1, 2, 3, 4, 5}
	tail = tail[:len(tail)-2] // 删除尾部 2 个
	fmt.Printf("tail remove: %v\n", tail)
}
