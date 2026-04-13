package grammar

import (
	"fmt"
	"unicode"
)

// ByteDemo 演示 byte 与 rune 的基础用法。
func ByteDemo() {
	// byte 是 uint8 的别名，常用于 ASCII 字符。
	var b1 byte = 'A'
	var b2 byte = 65
	var b3 byte = '\x41'
	fmt.Printf("byte: %c %c %c | %d %d %d\n", b1, b2, b3, b1, b2, b3)

	// rune 是 int32 的别名，表示一个 Unicode 码点。
	var r1 rune = '\u0041'   // A
	var r2 rune = '\u03B2'   // β
	var r3 rune = '\U00004F60' // 你
	fmt.Printf("rune as char: %c %c %c\n", r1, r2, r3)
	fmt.Printf("rune as int : %d %d %d\n", r1, r2, r3)
	fmt.Printf("rune as U+  : %U %U %U\n", r1, r2, r3)
	fmt.Printf("rune as HEX : %X %X %X\n", r1, r2, r3)

	// %c 打印字符，%d 打印整数值，%U 打印 Unicode 码点格式。
	ch := '9'
	fmt.Printf("ch=%c, int=%d, code=%U\n", ch, ch, ch)

	// unicode 包可判断字符类别。
	fmt.Printf("IsLetter('A')=%t\n", unicode.IsLetter('A'))
	fmt.Printf("IsDigit('9')=%t\n", unicode.IsDigit('9'))
	fmt.Printf("IsSpace(' ')=%t\n", unicode.IsSpace(' '))

}
