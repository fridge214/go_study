package collector

import "fmt"

// CopyDemo 演示切片的引用赋值与 copy() 复制行为。
func CopyDemo() {
	// 1) copy(dst, src) 会按较短切片长度复制，并返回复制数量。
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{9, 9, 9}
	n1 := copy(slice2, slice1)
	fmt.Printf("copy(slice2, slice1) -> n=%d, slice1=%v, slice2=%v\n", n1, slice1, slice2)

	n2 := copy(slice1, []int{7, 8})
	fmt.Printf("copy(slice1, []int{7,8}) -> n=%d, slice1=%v\n", n2, slice1)

	// 2) 赋值只是共享底层数组；copy 才是数据拷贝。
	const elementCount = 10
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	refData := srcData // 引用同一底层数组
	copyData := make([]int, elementCount)
	copied := copy(copyData, srcData)
	fmt.Printf("copy(copyData, srcData) -> n=%d\n", copied)

	srcData[0] = 999
	fmt.Printf("after srcData[0]=999, refData[0]=%d, copyData[0]=%d\n", refData[0], copyData[0])

	// 3) 复制局部切片到目标切片前缀。
	n3 := copy(copyData, srcData[4:6]) // 仅复制 2 个元素
	fmt.Printf("copy(copyData, srcData[4:6]) -> n=%d, copyData[:5]=%v\n", n3, copyData[:5])
}
