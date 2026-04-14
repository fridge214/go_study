package collector

import "fmt"

// DeleteMapDemo 演示 map 元素删除与“清空”操作。
func DeleteMapDemo() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	fmt.Printf("init scene=%v, len=%d\n", scene, len(scene))

	// 1) delete 删除指定 key。
	delete(scene, "brazil")
	fmt.Printf("after delete brazil=%v, len=%d\n", scene, len(scene))

	// 2) 删除不存在的 key 不会报错。
	delete(scene, "not-exist")
	fmt.Printf("after delete not-exist=%v, len=%d\n", scene, len(scene))

	// 3) 清空 map：重新 make 一个新的 map。
	scene = make(map[string]int)
	fmt.Printf("after remake clear=%v, len=%d\n", scene, len(scene))
}
