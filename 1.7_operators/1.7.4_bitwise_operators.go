package operators

import "fmt"

// ========== 1.7.4 位运算符 ==========

// demonstrateBitwiseOperators 位运算符演示
func demonstrateBitwiseOperators() {
	fmt.Println("=== 位运算符 ===")

	fmt.Println("位运算符对整数的二进制位进行操作：")
	fmt.Println("&  (AND)     - 按位与")
	fmt.Println("|  (OR)      - 按位或")
	fmt.Println("^  (XOR)     - 按位异或")
	fmt.Println("&^ (AND NOT) - 按位清除")
	fmt.Println("<< (左移)    - 左移位")
	fmt.Println(">> (右移)    - 右移位")
	fmt.Println()

	// 演示基本的位运算
	fmt.Println("=== 基本位运算 ===")
	fmt.Printf("0 & 0 = %d\n", 0&0)
	fmt.Printf("0 | 0 = %d\n", 0|0)
	fmt.Printf("0 ^ 0 = %d\n", 0^0)
	fmt.Println()

	fmt.Printf("0 & 1 = %d\n", 0&1)
	fmt.Printf("0 | 1 = %d\n", 0|1)
	fmt.Printf("0 ^ 1 = %d\n", 0^1)
	fmt.Println()

	fmt.Printf("1 & 1 = %d\n", 1&1)
	fmt.Printf("1 | 1 = %d\n", 1|1)
	fmt.Printf("1 ^ 1 = %d\n", 1^1)
	fmt.Println()

	fmt.Printf("1 & 0 = %d\n", 1&0)
	fmt.Printf("1 | 0 = %d\n", 1|0)
	fmt.Printf("1 ^ 0 = %d\n", 1^0)
}

// demonstrateBitwiseWithBinary 二进制表示的位运算
func demonstrateBitwiseWithBinary() {
	fmt.Println("\n=== 二进制表示的位运算 ===")

	a, b := uint8(60), uint8(13) // 60 = 0011 1100, 13 = 0000 1101

	fmt.Printf("a = %d (%08b)\n", a, a)
	fmt.Printf("b = %d (%08b)\n", b, b)
	fmt.Println()

	// 按位与
	and := a & b // 0000 1100 = 12
	fmt.Printf("a & b  = %d (%08b) - 按位与\n", and, and)

	// 按位或
	or := a | b // 0011 1101 = 61
	fmt.Printf("a | b  = %d (%08b) - 按位或\n", or, or)

	// 按位异或
	xor := a ^ b // 0011 0001 = 49
	fmt.Printf("a ^ b  = %d (%08b) - 按位异或\n", xor, xor)

	// 按位取反
	notA := ^a // 1100 0011 = 195 (注意：这是补码表示)
	fmt.Printf("^a     = %d (%08b) - 按位取反\n", notA, notA)

	// 按位清除 (AND NOT)
	andNot := a &^ b // 0011 0000 = 48
	fmt.Printf("a &^ b = %d (%08b) - 按位清除\n", andNot, andNot)
}

// demonstrateShiftOperators 移位运算符演示
func demonstrateShiftOperators() {
	fmt.Println("\n=== 移位运算符 ===")

	value := uint8(12) // 0000 1100

	fmt.Printf("原始值: %d (%08b)\n", value, value)
	fmt.Println()

	// 左移
	left1 := value << 1 // 0001 1000 = 24
	left2 := value << 2 // 0011 0000 = 48
	left3 := value << 3 // 0110 0000 = 96

	fmt.Printf("左移 1 位: %d (%08b) - 相当于乘以 2\n", left1, left1)
	fmt.Printf("左移 2 位: %d (%08b) - 相当于乘以 4\n", left2, left2)
	fmt.Printf("左移 3 位: %d (%08b) - 相当于乘以 8\n", left3, left3)
	fmt.Println()

	// 右移
	right1 := value >> 1 // 0000 0110 = 6
	right2 := value >> 2 // 0000 0011 = 3
	right3 := value >> 3 // 0000 0001 = 1

	fmt.Printf("右移 1 位: %d (%08b) - 相当于除以 2\n", right1, right1)
	fmt.Printf("右移 2 位: %d (%08b) - 相当于除以 4\n", right2, right2)
	fmt.Printf("右移 3 位: %d (%08b) - 相当于除以 8\n", right3, right3)

	// 负数移位
	fmt.Println("\n=== 负数移位（有符号整数）===")
	negative := int8(-12) // 1111 0100 (二进制补码)
	fmt.Printf("负数: %d (%08b)\n", negative, uint8(negative))

	rightShiftNegative := negative >> 1
	fmt.Printf("右移 1 位: %d (%08b) - 算术右移（符号位不变）\n",
		rightShiftNegative, uint8(rightShiftNegative))
}

// demonstrateBitwiseApplications 位运算应用演示
func demonstrateBitwiseApplications() {
	fmt.Println("\n=== 位运算应用场景 ===")

	// 1. 权限系统
	fmt.Println("1. 权限系统:")
	const (
		Read    = 1 << iota // 001
		Write               // 010
		Execute             // 100
	)

	userPerm := Read | Write // 011
	fmt.Printf("用户权限: %03b (%d)\n", userPerm, userPerm)
	fmt.Printf("可读: %v\n", userPerm&Read != 0)
	fmt.Printf("可写: %v\n", userPerm&Write != 0)
	fmt.Printf("可执行: %v\n", userPerm&Execute != 0)

	// 2. 状态标志
	fmt.Println("\n2. 状态标志:")
	const (
		Connected = 1 << iota
		Encrypted
		Compressed
	)

	status := Connected | Encrypted // 已连接且加密
	fmt.Printf("连接状态: %03b\n", status)
	fmt.Printf("已连接: %v\n", status&Connected != 0)
	fmt.Printf("已加密: %v\n", status&Encrypted != 0)
	fmt.Printf("已压缩: %v\n", status&Compressed != 0)

	// 3. RGB 颜色操作
	fmt.Println("\n3. RGB 颜色操作:")
	// 假设颜色格式: 0x00RRGGBB
	color := uint32(0x00FF8040) // 红色: FF, 绿色: 80, 蓝色: 40

	red := (color >> 16) & 0xFF
	green := (color >> 8) & 0xFF
	blue := color & 0xFF

	fmt.Printf("颜色: 0x%06X\n", color&0xFFFFFF)
	fmt.Printf("红色: 0x%02X (%d)\n", red, red)
	fmt.Printf("绿色: 0x%02X (%d)\n", green, green)
	fmt.Printf("蓝色: 0x%02X (%d)\n", blue, blue)
}

// BitwiseOperatorsDemo 位运算符演示主函数
func BitwiseOperatorsDemo() {
	fmt.Println("========== 1.7.4 位运算符 ==========")

	demonstrateBitwiseOperators()
	demonstrateBitwiseWithBinary()
	demonstrateShiftOperators()
	demonstrateBitwiseApplications()

	fmt.Println("\n=== 位运算符总结 ===")
	fmt.Println("✅ & 按位与 (AND)")
	fmt.Println("✅ | 按位或 (OR)")
	fmt.Println("✅ ^ 按位异或 (XOR)")
	fmt.Println("✅ &^ 按位清除 (AND NOT)")
	fmt.Println("✅ << 左移位")
	fmt.Println("✅ >> 右移位")
	fmt.Println("✅ ^ 一元运算符：按位取反")
	fmt.Println("✅ 应用场景：权限控制、状态标志、位操作算法")
}
