package collector

import (
	"container/list"
	"fmt"
)

// ListDemo 演示 container/list（双向链表）的常见操作。
func ListDemo() {
	l := list.New()
	fmt.Printf("init len=%d\n", l.Len())

	// 1) 头尾插入。
	front := l.PushFront("b")
	back := l.PushBack("d")
	l.PushFront("a")
	l.PushBack("e")
	fmt.Printf("after push len=%d\n", l.Len())

	// 2) 在指定元素前后插入。
	l.InsertAfter("c", front)  // a b c d e
	l.InsertBefore("tail", back) // a b c tail d e

	// 3) 遍历链表。
	fmt.Print("iterate: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()

	// 4) 移动元素。
	l.MoveToFront(back) // d a b c tail e
	l.MoveToBack(front) // d a c tail e b
	fmt.Print("after move: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()

	// 5) 删除元素。
	removed := l.Remove(front) // 删除 b
	fmt.Printf("removed=%v, len=%d\n", removed, l.Len())
	fmt.Print("final: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
