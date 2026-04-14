package collector

import "fmt"

// SliceDemo 演示 Go 切片的创建、截取、nil/empty 区别，以及 make 的长度和容量。
func SliceDemo() {
	// 1) 从数组生成切片：左闭右开 [start:end)。
	var highRiseBuilding [30]int
	for i := 0; i < len(highRiseBuilding); i++ {
		highRiseBuilding[i] = i + 1
	}

	fmt.Printf("highRiseBuilding[10:15] = %v\n", highRiseBuilding[10:15])
	fmt.Printf("highRiseBuilding[20:]   = %v\n", highRiseBuilding[20:])
	fmt.Printf("highRiseBuilding[:2]    = %v\n", highRiseBuilding[:2])

	// 2) 对已有切片再切片：[:] 表示原切片本身，0:0 可用于重置为“空视图”。
	a := []int{1, 2, 3}
	fmt.Printf("a[:]   = %v\n", a[:])
	fmt.Printf("a[0:0] = %v\n", a[0:0])

	// 3) 直接声明切片：nil 切片与空切片都 len=0，但 nil 判断结果不同。
	var strList []string       // nil
	var numList []int          // nil
	numListEmpty := []int{}    // 非 nil 空切片
	fmt.Printf("strList=%v, numList=%v, numListEmpty=%v\n", strList, numList, numListEmpty)
	fmt.Printf("len(strList)=%d, len(numList)=%d, len(numListEmpty)=%d\n",
		len(strList), len(numList), len(numListEmpty))
	fmt.Printf("strList==nil -> %t\n", strList == nil)
	fmt.Printf("numList==nil -> %t\n", numList == nil)
	fmt.Printf("numListEmpty==nil -> %t\n", numListEmpty == nil)

	// 4) make 创建切片：cap 是预留容量，len 是当前可见元素个数。使用make创建切片一定发生了内存分配
	m1 := make([]int, 2)
	m2 := make([]int, 2, 10)
	fmt.Printf("m1=%v, len=%d, cap=%d\n", m1, len(m1), cap(m1))
	fmt.Printf("m2=%v, len=%d, cap=%d\n", m2, len(m2), cap(m2))
}
