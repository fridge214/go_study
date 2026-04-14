package collector

import (
	"fmt"
	"sync"
)

// SyncMapDemo 演示 sync.Map 的基本操作与并发安全读写。
func SyncMapDemo() {
	var scene sync.Map

	// 1) Store / Load
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	v, ok := scene.Load("london")
	fmt.Printf("load london => value=%v ok=%t\n", v, ok)

	// 2) Delete
	scene.Delete("london")
	_, ok = scene.Load("london")
	fmt.Printf("after delete london => ok=%t\n", ok)

	// 3) Range 遍历（顺序不保证）。
	scene.Range(func(k, v any) bool {
		fmt.Printf("range: %v=%v\n", k, v)
		return true
	})

	// 4) 简单并发读写示例。
	var counter sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			counter.Store(n, n*n)
			if val, ok := counter.Load(n); ok {
				fmt.Printf("concurrent load: key=%d value=%v\n", n, val)
			}
		}(i)
	}
	wg.Wait()
}
