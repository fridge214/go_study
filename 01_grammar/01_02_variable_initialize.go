package grammar

import "fmt"

// VariableInitialize 演示 Go 语言变量初始化的各种方式
func VariableInitialize() {
	fmt.Println("=== Go 语言变量初始化示例 ===\n")

	// ========================================
	// 1. 变量的自动零值初始化
	// Go 语言在声明变量时会自动初始化为类型的零值
	// ========================================
	fmt.Println("--- 1. 变量的自动零值初始化 ---")

	var intVar int
	var floatVar float64
	var stringVar string
	var boolVar bool
	var sliceVar []int
	var funcVar func()
	var ptrVar *int

	fmt.Printf("int 零值: %d\n", intVar)
	fmt.Printf("float64 零值: %f\n", floatVar)
	fmt.Printf("string 零值: '%s'\n", stringVar)
	fmt.Printf("bool 零值: %t\n", boolVar)
	fmt.Printf("slice 零值: %v\n", sliceVar)
	fmt.Printf("func 零值: %v\n", funcVar)
	fmt.Printf("pointer 零值: %v\n", ptrVar)

	// ========================================
	// 2. 标准格式：var 变量名 类型 = 表达式
	// ========================================
	fmt.Println("\n--- 2. 标准格式初始化 ---")

	// 游戏中玩家血量初始值为 100
	var hp int = 100
	fmt.Printf("玩家血量 (hp): %d\n", hp)

	// 声明并初始化字符串变量
	var playerName string = "玩家1"
	fmt.Printf("玩家名称: %s\n", playerName)

	// 声明并初始化布尔变量
	var isAlive bool = true
	fmt.Printf("玩家状态 (isAlive): %t\n", isAlive)

	// ========================================
	// 3. 编译器推导类型格式：var 变量名 = 表达式
	// 编译器会根据右值自动推导变量类型
	// ========================================

	fmt.Println("\n--- 3. 编译器推导类型格式 ---")

	// 右值为整型，编译器推导为 int
	var attack = 40
	var defence = 20
	fmt.Printf("攻击力 (attack): %d, 防御力 (defence): %d\n", attack, defence)

	// 右值为浮点数，编译器会推导为 float64
	// 但我们可以显式指定为 float32
	var damageRate float32 = 0.17
	fmt.Printf("伤害系数 (damageRate): %f\n", damageRate)

	// 复杂表达式：包含类型转换
	var damage = float32(attack-defence) * damageRate
	fmt.Printf("伤害值 (damage): %.1f\n", damage)

	// ========================================
	// 4. 短变量声明并初始化：:= 表达式
	// 更加简洁的写法，编译器自动推导类型
	// ========================================
	fmt.Println("\n--- 4. 短变量声明并初始化 ---")

	// 使用短变量声明
	hp2 := 100
	fmt.Printf("玩家血量 (hp2): %d\n", hp2)

	// 同时声明多个变量
	name, level := "玩家2", 5
	fmt.Printf("玩家名称: %s, 等级: %d\n", name, level)

	// ========================================
	// 5. 短变量声明在实际开发中的应用
	// 例如网络连接、文件操作等返回多个值的场景
	// ========================================
	fmt.Println("\n--- 5. 短变量声明的实际应用 ---")

	// 注意：这里使用注释演示，因为无法实际连接
	// conn, err := net.Dial("tcp", "127.0.0.1:8080")
	// if err != nil {
	//     fmt.Printf("连接失败: %v\n", err)
	//     return
	// }
	// defer conn.Close()
	// fmt.Printf("连接成功: %v\n", conn)

	// 模拟一个返回多个值的函数
	result, success := mockNetworkCall()
	fmt.Printf("网络调用结果: %s, 成功: %t\n", result, success)

	// ========================================
	// 6. 短变量声明的限制
	// ========================================
	fmt.Println("\n--- 6. 短变量声明的限制 ---")

	// 6.1 变量必须是新声明的
	// 以下代码会报错：no new variables on left side of :=
	// var hp3 int
	// hp3 := 10  // 错误：hp3 已经声明过

	// 正确的做法：先声明，再赋值
	var hp3 int
	hp3 = 10
	fmt.Printf("hp3 (先声明后赋值): %d\n", hp3)

	// 6.2 在多个短变量声明中，至少有一个新变量
	// 即使其他变量重复声明，编译器也不会报错
	fmt.Println("\n--- 6.2 多个短变量声明示例 ---")
	// 假设 err 已经在其他地方声明过
	// conn, err := net.Dial("tcp", "127.0.0.1:8080")
	// conn2, err := net.Dial("tcp", "127.0.0.1:8080")  // err 重复声明不会报错

	// 模拟演示
	_, err1 := mockNetworkCall()
	_, err2 := mockNetworkCall()
	fmt.Printf("两次调用的错误值: %v, %v\n", err1, err2)

	// ========================================
	// 7. 类型转换示例
	// ========================================
	fmt.Println("\n--- 7. 类型转换示例 ---")

	// 整型转浮点型
	var intNum int = 10
	var floatNum float64 = float64(intNum)
	fmt.Printf("int 转 float64: %d -> %f\n", intNum, floatNum)

	// 浮点型转整型（会丢失小数部分）
	var pi float64 = 3.14159
	var piInt int = int(pi)
	fmt.Printf("float64 转 int: %f -> %d\n", pi, piInt)

	// 字符串转整型
	var strNum string = "123"
	var num int
	fmt.Sscanf(strNum, "%d", &num)
	fmt.Printf("字符串转 int: '%s' -> %d\n", strNum, num)

	// ========================================
	// 8. 变量初始化的最佳实践
	// ========================================
	fmt.Println("\n--- 8. 变量初始化最佳实践 ---")

	// 8.1 需要显式指定类型时使用 var
	var maxConnections int = 1000
	fmt.Printf("最大连接数 (需要指定类型): %d\n", maxConnections)

	// 8.2 简单初始化使用短变量声明
	username := "admin"
	fmt.Printf("用户名 (简单初始化): %s\n", username)

	// 8.3 复杂表达式使用 var 便于阅读
	var totalScore = (attack * 10) + (level * 100)
	fmt.Printf("总分 (复杂表达式): %d\n", totalScore)

	// 8.4 多变量初始化
	x, y, z := 1, 2, 3
	fmt.Printf("多变量初始化: x=%d, y=%d, z=%d\n", x, y, z)
}

// mockNetworkCall 模拟网络调用，返回结果和错误信息
func mockNetworkCall() (string, error) {
	return "请求成功", nil
}
