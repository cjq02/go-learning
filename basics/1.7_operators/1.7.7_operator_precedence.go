package operators

import "fmt"

// ========== 1.7.7 运算优先级 ==========

// demonstrateArithmeticPrecedence 算术运算优先级演示
func demonstrateArithmeticPrecedence() {
	fmt.Println("=== 算术运算优先级 ===")

	var a int = 21
	var b int = 10
	var c int = 16
	var d int = 5
	var e int

	fmt.Printf("初始值: a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)

	// 优先级：乘除 > 加减
	e = (a + b) * c / d // (31 * 16) / 5
	fmt.Printf("(a + b) * c / d = %d\n", e)

	e = ((a + b) * c) / d // (31 * 16) / 5
	fmt.Printf("((a + b) * c) / d = %d\n", e)

	e = (a + b) * (c / d) // 31 * (16/5)
	fmt.Printf("(a + b) * (c / d) = %d\n", e)

	e = a + (b*c)/d // 21 + (160/5)
	fmt.Printf("a + (b * c) / d = %d\n", e)

	// 整数除法会截断小数部分
	fmt.Printf("注意：16/5 = %d (整数除法向下取整)\n", 16/5)
}

// demonstrateComplexPrecedence 复杂运算优先级演示
func demonstrateComplexPrecedence() {
	fmt.Println("\n=== 复杂运算优先级 ===")

	// 复杂的优先级示例
	// 计算顺序：括号 > 乘除 > 加减 > 移位 > 关系 > 相等 > 位与 > 位异或 > 位或 > 逻辑与 > 逻辑或

	// 21 + (160/5) = 21 + 32 = 53
	result1 := 21 + 160/5
	fmt.Printf("21 + 160/5 = %d (先算除法，再算加法)\n", result1)

	// 2 & 2 = 2; 2 * 3 = 6; 6 << 1 = 12; 3 + 4 = 7; 7 ^ 3 = 4; 4 | 12 = 12
	result2 := 3 + 4 ^ 3 | 2&2*3<<1
	fmt.Printf("3 + 4 ^ 3 | 2&2*3<<1 = %d\n", result2)
	fmt.Printf("计算步骤解析:\n")
	fmt.Printf("  1. 2&2 = %d (位与)\n", 2&2)
	fmt.Printf("  2. %d*3 = %d (乘法)\n", 2&2, (2&2)*3)
	fmt.Printf("  3. %d<<1 = %d (左移)\n", (2&2)*3, ((2&2)*3)<<1)
	fmt.Printf("  4. 3+4 = %d (加法)\n", 3+4)
	fmt.Printf("  5. %d ^ 3 = %d (异或)\n", 3+4, (3+4)^3)
	fmt.Printf("  6. %d | %d = %d (或运算)\n", (3+4)^3, ((2&2)*3)<<1, ((3+4)^3)|(((2&2)*3)<<1))

	fmt.Printf("最终结果: %d == 12 ? %v\n", result2, result2 == 12)
}

// demonstrateLogicalPrecedence 逻辑运算优先级演示
func demonstrateLogicalPrecedence() {
	fmt.Println("\n=== 逻辑运算优先级 ===")

	a, b, c := true, false, true

	// 逻辑运算符优先级：! > && > ||
	result1 := a && b || c
	result2 := (a && b) || c  // 明确使用括号
	result3 := a && (b || c)

	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)
	fmt.Printf("a && b || c = %v\n", result1)
	fmt.Printf("(a && b) || c = %v\n", result2)
	fmt.Printf("a && (b || c) = %v\n", result3)

	// 复杂的逻辑表达式
	complexExpr := a && !b || c && a
	fmt.Printf("a && !b || c && a = %v\n", complexExpr)
	fmt.Printf("等价于: ((a && (!b)) || (c && a))\n", complexExpr)
}

// demonstrateBitwisePrecedence 位运算优先级演示
func demonstrateBitwisePrecedence() {
	fmt.Println("\n=== 位运算优先级 ===")

	// 位运算符优先级：<< >> > & > ^ > |

	a, b := 5, 3 // 5=101, 3=011

	// 移位运算符优先级最高
	result1 := a<<1 | b    // (a<<1) | b
	result2 := a | b<<1    // a | (b<<1)
	result3 := a<<1 ^ b<<2 // (a<<1) ^ (b<<2)

	fmt.Printf("a=%d (%03b), b=%d (%03b)\n", a, a, b, b)
	fmt.Printf("a<<1 | b = %d (%04b)\n", result1, result1)
	fmt.Printf("a | b<<1 = %d (%04b)\n", result2, result2)
	fmt.Printf("a<<1 ^ b<<2 = %d (%04b)\n", result3, result3)
}

// demonstrateMixedPrecedence 混合运算优先级演示
func demonstrateMixedPrecedence() {
	fmt.Println("\n=== 混合运算优先级 ===")

	// 综合示例：算术 + 位运算 + 比较 + 逻辑
	x, y, z := 10, 5, 3

	// 完整的优先级顺序：
	// 1. 算术运算符 (* / % + -)
	// 2. 位运算符 (<< >> & ^ |)
	// 3. 比较运算符 (== != < <= > >=)
	// 4. 逻辑运算符 (&& ||)

	result1 := x + y*z == 25 && x > y
	result2 := (x + y)*z == 45 && x > y
	result3 := x + (y*z) == 25 && x > y

	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)
	fmt.Printf("x + y*z == 25 && x > y = %v\n", result1)
	fmt.Printf("  计算: %d + %d*%d == 25 && %d > %d\n", x, y, z, x, y)
	fmt.Printf("  等价: (%d + (%d*%d)) == 25 && (%d > %d)\n", x, y, z, x, y)

	fmt.Printf("\n(x + y)*z == 45 && x > y = %v\n", result2)
	fmt.Printf("x + (y*z) == 25 && x > y = %v\n", result3)
}

// demonstrateParenthesesImportance 括号重要性演示
func demonstrateParenthesesImportance() {
	fmt.Println("\n=== 括号的重要性 ===")

	a, b, c, d := 2, 3, 4, 5

	// 不使用括号
	result1 := a + b*c - d/2

	// 使用括号明确优先级
	result2 := (a + b) * c - d/2
	result3 := a + (b*c) - d/2
	result4 := a + b*c - (d/2)

	fmt.Printf("变量: a=%d, b=%d, c=%d, d=%d\n", a, b, c, d)
	fmt.Printf("a + b*c - d/2 = %d\n", result1)
	fmt.Printf("  计算顺序: a + (b*c) - (d/2)\n")

	fmt.Printf("\n使用括号明确优先级:\n")
	fmt.Printf("(a + b) * c - d/2 = %d\n", result2)
	fmt.Printf("a + (b*c) - d/2 = %d\n", result3)
	fmt.Printf("a + b*c - (d/2) = %d\n", result4)

	fmt.Println("\n注：可以使用小括号，提高部分计算的优先级。也可以提高表达式的可读性。")
}

// OperatorPrecedenceDemo 运算优先级演示主函数
func OperatorPrecedenceDemo() {
	fmt.Println("========== 1.7.7 运算优先级 ==========")

	demonstrateArithmeticPrecedence()
	demonstrateComplexPrecedence()
	demonstrateLogicalPrecedence()
	demonstrateBitwisePrecedence()
	demonstrateMixedPrecedence()
	demonstrateParenthesesImportance()

	fmt.Println("\n=== 运算优先级总结 ===")
	fmt.Println("优先级从高到低：")
	fmt.Println("1. 括号 ()")
	fmt.Println("2. 一元运算符 ! & * <-")
	fmt.Println("3. 算术运算符 * / % << >> & &^")
	fmt.Println("4. 算术运算符 + - | ^")
	fmt.Println("5. 比较运算符 == != < <= > >=")
	fmt.Println("6. 逻辑运算符 &&")
	fmt.Println("7. 逻辑运算符 ||")
	fmt.Println("8. 赋值运算符 = += -= *= /= %= <<= >>= &= ^= |= &^=")
	fmt.Println()
	fmt.Println("💡 提示：使用括号可以明确控制计算顺序，提高代码可读性")
}
