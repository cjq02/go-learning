package operators

import "fmt"

// ========== 1.7.3 逻辑运算符 ==========

// demonstrateLogicalOperators 逻辑运算符演示
func demonstrateLogicalOperators() {
	fmt.Println("=== 逻辑运算符 ===")

	a := true
	b := false

	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Println()

	// 逻辑与 (AND)
	fmt.Printf("a && b = %v (逻辑与: 都为true才为true)\n", a && b)
	fmt.Printf("a && a = %v\n", a && a)
	fmt.Printf("b && b = %v\n", b && b)

	// 逻辑或 (OR)
	fmt.Printf("a || b = %v (逻辑或: 只要一个为true就为true)\n", a || b)
	fmt.Printf("a || a = %v\n", a || a)
	fmt.Printf("b || b = %v\n", b || b)

	// 逻辑非 (NOT)
	fmt.Printf("!a = %v (逻辑非: 取反)\n", !a)
	fmt.Printf("!b = %v\n", !b)
	fmt.Printf("!(a && b) = %v\n", !(a && b))
	fmt.Printf("!(a || b) = %v\n", !(a || b))

	// 真值表演示
	fmt.Println("\n=== 逻辑运算符真值表 ===")
	fmt.Println("A     B     | A && B | A || B | !A   | !B")
	fmt.Println("-------------|--------|--------|------|------")

	truthTable := []struct{ a, b bool }{
		{false, false},
		{false, true},
		{true, false},
		{true, true},
	}

	for _, row := range truthTable {
		and := row.a && row.b
		or := row.a || row.b
		notA := !row.a
		notB := !row.b
		fmt.Printf("%-5v %-5v | %-6v | %-6v | %-4v | %-4v\n",
			row.a, row.b, and, or, notA, notB)
	}
}

// demonstrateShortCircuit 短路求值演示
func demonstrateShortCircuit() {
	fmt.Println("\n=== 短路求值 (Short-circuit Evaluation) ===")
	fmt.Println("Go 中的逻辑运算符使用短路求值：")

	fmt.Println("对于 && (AND)：如果左操作数为false，右操作数不会被求值")
	fmt.Println("对于 || (OR)：如果左操作数为true，右操作数不会被求值")

	fmt.Println()

	// 演示短路求值
	fmt.Println("演示 && 短路求值:")
	result1 := false && expensiveOperation("不会执行")
	fmt.Printf("false && expensiveOperation() = %v\n", result1)

	result2 := true && expensiveOperation("会执行")
	fmt.Printf("true && expensiveOperation() = %v\n", result2)

	fmt.Println("\n演示 || 短路求值:")
	result3 := true || expensiveOperation("不会执行")
	fmt.Printf("true || expensiveOperation() = %v\n", result3)

	result4 := false || expensiveOperation("会执行")
	fmt.Printf("false || expensiveOperation() = %v\n", result4)
}

// expensiveOperation 模拟耗时操作
func expensiveOperation(msg string) bool {
	fmt.Printf("执行耗时操作: %s\n", msg)
	return true
}

// demonstrateComplexLogicalExpressions 复杂逻辑表达式演示
func demonstrateComplexLogicalExpressions() {
	fmt.Println("\n=== 复杂逻辑表达式 ===")

	x, y, z := 10, 5, 15

	// 复合条件
	condition1 := x > y && y < z
	condition2 := x == 10 || y == 10
	condition3 := !(x < y) && z >= 15

	fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
	fmt.Printf("x > y && y < z = %v\n", condition1)
	fmt.Printf("x == 10 || y == 10 = %v\n", condition2)
	fmt.Printf("!(x < y) && z >= 15 = %v\n", condition3)

	// 逻辑运算符优先级：! > && > ||
	// 使用括号提高可读性
	complexExpr := (x > y && y < z) || (x == 10 && !(y > z))
	fmt.Printf("\n复杂表达式: (x > y && y < z) || (x == 10 && !(y > z)) = %v\n", complexExpr)
}

// LogicalOperatorsDemo 逻辑运算符演示主函数
func LogicalOperatorsDemo() {
	fmt.Println("========== 1.7.3 逻辑运算符 ==========")

	demonstrateLogicalOperators()
	demonstrateShortCircuit()
	demonstrateComplexLogicalExpressions()

	fmt.Println("\n=== 逻辑运算符总结 ===")
	fmt.Println("✅ && 逻辑与 (AND)")
	fmt.Println("✅ || 逻辑或 (OR)")
	fmt.Println("✅ !  逻辑非 (NOT)")
	fmt.Println("✅ 短路求值：提高性能，避免不必要的计算")
	fmt.Println("✅ 优先级：! > && > ||")
	fmt.Println("✅ 结果类型：bool")
}
