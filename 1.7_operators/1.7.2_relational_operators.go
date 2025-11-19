package operators

import "fmt"

// ========== 1.7.2 关系运算符 ==========

// demonstrateRelationalOperators 关系运算符演示
func demonstrateRelationalOperators() {
	fmt.Println("=== 关系运算符 ===")
	fmt.Println("关系运算符结果只会是 bool 类型。")

	a := 1
	b := 5

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a == b: %v (等于)\n", a == b)
	fmt.Printf("a != b: %v (不等于)\n", a != b)
	fmt.Printf("a > b:  %v (大于)\n", a > b)
	fmt.Printf("a < b:  %v (小于)\n", a < b)
	fmt.Printf("a >= b: %v (大于等于)\n", a >= b)
	fmt.Printf("a <= b: %v (小于等于)\n", a <= b)

	// 字符串比较
	str1 := "hello"
	str2 := "world"
	str3 := "hello"

	fmt.Printf("\n字符串比较:\n")
	fmt.Printf("str1 = %q, str2 = %q, str3 = %q\n", str1, str2, str3)
	fmt.Printf("str1 == str2: %v\n", str1 == str2)
	fmt.Printf("str1 == str3: %v\n", str1 == str3)
	fmt.Printf("str1 < str2:  %v (字典序比较)\n", str1 < str2)

	// 浮点数比较
	x := 3.14
	y := 3.14
	z := 2.71

	fmt.Printf("\n浮点数比较:\n")
	fmt.Printf("x = %.2f, y = %.2f, z = %.2f\n", x, y, z)
	fmt.Printf("x == y: %v\n", x == y)
	fmt.Printf("x != z: %v\n", x != z)
	fmt.Printf("x > z:  %v\n", x > z)

	// 注意：浮点数比较可能有精度问题
	a1 := 0.1
	a2 := 0.1
	a3 := a1 + a2 // 0.2

	fmt.Printf("\n浮点数精度问题:\n")
	fmt.Printf("0.1 + 0.1 = %.20f\n", a3)
	fmt.Printf("0.1 + 0.1 == 0.2: %v (可能为false!)\n", a3 == 0.2)

	// 布尔值比较
	bool1 := true
	bool2 := false

	fmt.Printf("\n布尔值比较:\n")
	fmt.Printf("bool1 = %v, bool2 = %v\n", bool1, bool2)
	fmt.Printf("bool1 == bool2: %v\n", bool1 == bool2)
	fmt.Printf("bool1 != bool2: %v\n", bool1 != bool2)
}

// RelationalOperatorsDemo 关系运算符演示主函数
func RelationalOperatorsDemo() {
	fmt.Println("========== 1.7.2 关系运算符 ==========")
	demonstrateRelationalOperators()

	fmt.Println("\n=== 关系运算符总结 ===")
	fmt.Println("✅ == != > < >= <= 关系运算符")
	fmt.Println("✅ 结果类型：bool")
	fmt.Println("✅ 支持数值、字符串、布尔值比较")
	fmt.Println("✅ 字符串按字典序比较")
	fmt.Println("✅ 浮点数比较可能有精度问题")
}
