package collector

import (
	"fmt"
	"sort"
)

// MapForDemo 演示 map 的遍历方式和有序输出方法。
func MapForDemo() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	// 1) 同时遍历 key 和 value（无序）。
	for k, v := range scene {
		fmt.Printf("k=%s v=%d\n", k, v)
	}

	// 2) 只遍历 value。
	for _, v := range scene {
		fmt.Printf("only value=%d\n", v)
	}

	// 3) 只遍历 key。
	for k := range scene {
		fmt.Printf("only key=%s\n", k)
	}

	// 4) 需要固定顺序时：取出 key -> 排序 -> 按序读取 value。
	var keyList []string
	for k := range scene {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	fmt.Printf("sorted keys=%v\n", keyList)
	for _, k := range keyList {
		fmt.Printf("sorted k=%s v=%d\n", k, scene[k])
	}
}
