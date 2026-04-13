package grammar

import "fmt"

// OperatorPriorityDemo 演示 Go 运算符优先级与结合性。
func OperatorPriorityDemo() {
	// 1) 乘除优先于加减。
	a, b, c := 16, 4, 2
	r1 := a + b*c   // 等价于 a + (b*c)
	r2 := (a + b) * c
	fmt.Printf("a + b*c = %d, (a+b)*c = %d\n", r1, r2)

	// 2) 同级运算符通常从左到右结合。
	// 例如减法：20-5-3 == (20-5)-3
	r3 := 20 - 5 - 3
	r4 := 20 - (5 - 3)
	fmt.Printf("20-5-3 = %d, 20-(5-3) = %d\n", r3, r4)

	// 3) 关系运算高于 &&，&& 高于 ||。
	x := 7
	logic := x > 5 && x < 10 || x == 100
	fmt.Printf("x>5 && x<10 || x==100 -> %t\n", logic)

	// 4) 位运算示例：移位、按位与、按位或。
	// 先算移位，再按位与/或（按 Go 规则从高到低结合）。
	bit := 1<<3 | 2 // (1<<3) | 2 = 8 | 2 = 10
	mask := bit & 6
	fmt.Printf("bit=%d, bit&6=%d\n", bit, mask)

	// 5) 一元运算优先级较高；括号始终最清晰。
	n := -a + b // (-a) + b
	fmt.Printf("-a + b = %d\n", n)

	// 实战建议：拿不准优先级时优先加括号，代码可读性更高。
}
