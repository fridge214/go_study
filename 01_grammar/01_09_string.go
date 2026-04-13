package grammar

import (
	"fmt"
	"unicode/utf8"
)

// StringDemo 演示 Go 字符串的基本用法。
func StringDemo() {
	// 字符串是值类型，底层是 UTF-8 字节序列；创建后不可变。
	s := "Go语言"
	fmt.Printf("s = %s\n", s)

	// 双引号字符串支持转义字符（如 \n）。
	escaped := "第一行\n第二行"
	fmt.Println("escaped:")
	fmt.Println(escaped)

	// 反引号字符串会原样保留内容（换行也会保留，转义不生效）。
	raw := `raw line1
raw line2
\n will not escape`
	fmt.Println("raw:")
	fmt.Println(raw)

	// len(s) 返回字节长度；索引取到的是“字节”。
	ascii := "hello"
	fmt.Printf("ascii=%s, len=%d, ascii[0]=%c\n", ascii, len(ascii), ascii[0])

	// UTF-8 下中文通常占多个字节：len 是字节数，RuneCountInString 是字符数。
	text := "Go语言"
	fmt.Printf("text=%s, bytes(len)=%d, runes=%d\n", text, len(text), utf8.RuneCountInString(text))

	// 字符串比较按字节字典序进行。
	fmt.Printf("\"abc\" < \"abd\" -> %t\n", "abc" < "abd")

	// 字符串拼接可用 + 和 +=。
	part := "hello"
	part += ", "
	part = part + "world"
	fmt.Printf("concat: %s\n", part)

	// 字符串不可直接修改单个字节（下面写法会编译错误，故仅注释说明）。
	// s[0] = 'g'
}
