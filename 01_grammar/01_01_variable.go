package grammar

import "fmt"

func Variable() {
	// ========================================
	// 1. 标准格式：var 变量名 变量类型
	// ========================================
	var age int
	var name string
	var isStudent bool
	var height float32

	fmt.Println("=== 标准格式声明变量 ===")
	fmt.Printf("age (int 零值): %d\n", age)
	fmt.Printf("name (string 零值): '%s'\n", name)
	fmt.Printf("isStudent (bool 零值): %t\n", isStudent)
	fmt.Printf("height (float32 零值): %f\n", height)

	// 声明并初始化
	var score int = 100
	var message string = "Hello, Go!"
	fmt.Printf("\n初始化后的变量:\n")
	fmt.Printf("score: %d\n", score)
	fmt.Printf("message: %s\n", message)

	// ========================================
	// 2. 批量格式：使用 var 和括号将多个变量定义放在一起
	// ========================================
	var (
		a       int
		b       string
		c       []float32
		d       func() bool
		e       struct{ x int }
		pi      float64 = 3.14159
		version string  = "1.0"
	)

	fmt.Println("\n=== 批量格式声明变量 ===")
	fmt.Printf("a (int): %d\n", a)
	fmt.Printf("b (string): '%s'\n", b)
	fmt.Printf("c (slice): %v\n", c)
	fmt.Printf("d (func): %v\n", d)
	fmt.Printf("e (struct): %+v\n", e)
	fmt.Printf("pi: %f\n", pi)
	fmt.Printf("version: %s\n", version)

	// ========================================
	// 3. 简短格式：:= 表达式（只能在函数内部使用）
	// ========================================
	x := 100
	y, z := 1, 2
	name2, city := "张三", "北京"

	fmt.Println("\n=== 简短格式声明变量 ===")
	fmt.Printf("x: %d\n", x)
	fmt.Printf("y: %d, z: %d\n", y, z)
	fmt.Printf("name2: %s, city: %s\n", name2, city)

	// ========================================
	// 4. Go 语言基本类型示例
	// ========================================
	fmt.Println("\n=== Go 语言基本类型 ===")

	// bool 类型
	var isActive bool = true
	fmt.Printf("bool (isActive): %t\n", isActive)

	// string 类型
	var greeting string = "你好，Go 语言！"
	fmt.Printf("string (greeting): %s\n", greeting)

	// int 系列类型
	var intVal int = 42
	var int8Val int8 = 8
	var int16Val int16 = 16
	var int32Val int32 = 32
	var int64Val int64 = 64
	fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n",
		intVal, int8Val, int16Val, int32Val, int64Val)

	// uint 系列类型（无符号整数）
	var uintVal uint = 42
	var uint8Val uint8 = 8
	var uint16Val uint16 = 16
	var uint32Val uint32 = 32
	var uint64Val uint64 = 64
	fmt.Printf("uint: %d, uint8: %d, uint16: %d, uint32: %d, uint64: %d\n",
		uintVal, uint8Val, uint16Val, uint32Val, uint64Val)

	// byte (uint8 的别名)
	var byteVal byte = 'A'
	fmt.Printf("byte (uint8 别名): %c (ASCII: %d)\n", byteVal, byteVal)

	// rune (int32 的别名，代表一个 Unicode 码)
	var runeVal rune = '中'
	fmt.Printf("rune (int32 别名): %c (Unicode: %U)\n", runeVal, runeVal)

	// float 系列类型
	var float32Val float32 = 3.14
	var float64Val float64 = 3.141592653589793
	fmt.Printf("float32: %f, float64: %f\n", float32Val, float64Val)

	// complex 系列类型（复数）
	var complex64Val complex64 = complex(1.0, 2.0)
	var complex128Val complex128 = complex(1.0, 2.0)
	fmt.Printf("complex64: %v, complex128: %v\n", complex64Val, complex128Val)

	// uintptr 类型（无符号整数，可以存储指针的地址）
	fmt.Printf("uintptr: 用于存储指针地址的无符号整数类型\n")

	// ========================================
	// 5. 变量的零值示例
	// ========================================
	fmt.Println("\n=== 变量的零值 ===")
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroPtr *int

	fmt.Printf("int 零值: %d\n", zeroInt)
	fmt.Printf("float64 零值: %f\n", zeroFloat)
	fmt.Printf("bool 零值: %t\n", zeroBool)
	fmt.Printf("string 零值: '%s'\n", zeroString)
	fmt.Printf("pointer 零值: %v\n", zeroPtr)

	// ========================================
	// 6. 骆驼命名法示例
	// ========================================
	fmt.Println("\n=== 骆驼命名法示例 ===")
	numShips := 10
	startDate := "2024-01-01"
	totalPrice := 99.99
	fmt.Printf("numShips: %d\n", numShips)
	fmt.Printf("startDate: %s\n", startDate)
	fmt.Printf("totalPrice: %f\n", totalPrice)

	// ========================================
	// 7. var 形式 vs 简短形式的选择
	// ========================================
	fmt.Println("\n=== var 形式 vs 简短形式 ===")

	// 需要显式指定类型时使用 var
	var explicitType float64 = 42 // 42 会被转换为 float64
	fmt.Printf("显式指定类型 (float64): %f\n", explicitType)

	// 变量稍后会被重新赋值，初始值无关紧要
	var counter int
	counter = 10
	counter = 20
	fmt.Printf("counter (稍后赋值): %d\n", counter)

	// 大部分局部变量使用简短形式
	localVar := "局部变量"
	fmt.Printf("简短形式声明: %s\n", localVar)

	// ========================================
	// 8. 指针类型声明示例
	// ========================================
	fmt.Println("\n=== 指针类型声明 ===")
	var num1, num2 int = 10, 20
	var ptr1, ptr2 *int = &num1, &num2
	fmt.Printf("num1: %d, ptr1 (指向 num1): %p -> %d\n", num1, ptr1, *ptr1)
	fmt.Printf("num2: %d, ptr2 (指向 num2): %p -> %d\n", num2, ptr2, *ptr2)
}
