// Package controlflow 演示 Go 语言流程控制语句的使用
package controlflow

import "fmt"

// ========== 1.8.1 if 语句 ==========

// demonstrateBasicIf 演示基本的 if 语句
func demonstrateBasicIf() {
	fmt.Println("=== 1. 基本 if 语句 ===")

	a := 10

	// 基本 if 语句
	if a > 5 {
		fmt.Println("a 大于 5")
	}

	// if-else 语句
	if a > 15 {
		fmt.Println("a 大于 15")
	} else {
		fmt.Println("a 不大于 15")
	}

	fmt.Println()
}

// demonstrateIfElseIf 演示 if-else if-else 语句
func demonstrateIfElseIf() {
	fmt.Println("=== 2. if-else if-else 语句 ===")

	score := 85

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	fmt.Println()
}

// demonstrateNestedIf 演示嵌套的 if 语句
func demonstrateNestedIf() {
	fmt.Println("=== 3. 嵌套的 if 语句 ===")

	x, y := 10, 20

	if x > 5 {
		if y > 15 {
			fmt.Println("x 大于 5 且 y 大于 15")
		} else {
			fmt.Println("x 大于 5 但 y 不大于 15")
		}
	} else {
		if y > 15 {
			fmt.Println("x 不大于 5 但 y 大于 15")
		} else {
			fmt.Println("x 不大于 5 且 y 不大于 15")
		}
	}

	fmt.Println()
}

// demonstrateIfWithAssignment 演示带初始化语句的 if 语句
func demonstrateIfWithAssignment() {
	fmt.Println("=== 4. 带初始化语句的 if 语句 ===")

	// 在 if 语句中声明并初始化变量
	// 变量的作用域仅限于 if-else 块内
	if num := 15; num > 10 {
		fmt.Printf("num(%d) 大于 10\n", num)
	} else {
		fmt.Printf("num(%d) 不大于 10\n", num)
	}

	// 注意：num 变量在此处不可访问
	// fmt.Println(num) // 这行代码会编译错误

	fmt.Println()
}

// demonstrateComplexIf 演示复杂的 if 语句组合
func demonstrateComplexIf() {
	fmt.Println("=== 5. 复杂的 if 语句组合 ===")

	a, b := 10, 5

	// 带初始化语句的 if-else if-else
	if x := a + b; x > 10 {
		fmt.Printf("x(%d) = a(%d) + b(%d) 大于 10\n", x, a, b)

		// 嵌套的 if 语句
		if y := a - b; y > 0 {
			fmt.Printf("y(%d) = a(%d) - b(%d) 大于 0\n", y, a, b)
		} else {
			fmt.Printf("y(%d) = a(%d) - b(%d) 不大于 0\n", y, a, b)
		}
	} else if x == 10 {
		fmt.Printf("x(%d) 等于 10\n", x)
	} else {
		fmt.Printf("x(%d) 小于 10\n", x)
	}

	fmt.Println()
}

// demonstrateIfWithFunctionCall 演示 if 语句中的函数调用
func demonstrateIfWithFunctionCall() {
	fmt.Println("=== 6. if 语句中的函数调用 ===")

	// 在条件表达式中调用函数
	if result := calculate(5, 3); result > 7 {
		fmt.Printf("calculate(5, 3) = %d 大于 7\n", result)
	} else {
		fmt.Printf("calculate(5, 3) = %d 不大于 7\n", result)
	}

	fmt.Println()
}

// calculate 辅助函数用于演示
func calculate(a, b int) int {
	return a + b
}

// demonstrateBooleanExpressions 演示布尔表达式
func demonstrateBooleanExpressions() {
	fmt.Println("=== 7. 布尔表达式 ===")

	isTrue := true
	isFalse := false

	// 逻辑与 &&
	if isTrue && !isFalse {
		fmt.Println("isTrue 为 true 且 isFalse 为 false")
	}

	// 逻辑或 ||
	if isTrue || isFalse {
		fmt.Println("isTrue 或 isFalse 至少有一个为 true")
	}

	// 复杂的布尔表达式
	a, b, c := 5, 10, 15
	if a < b && (b > c || a+c > 20) {
		fmt.Println("复杂的布尔表达式成立")
	}

	fmt.Println()
}

// demonstrateIfScope 演示 if 语句的作用域
func demonstrateIfScope() {
	fmt.Println("=== 8. if 语句的作用域 ===")

	value := 100

	// 在 if 语句中声明的变量只在该 if-else 块中有效
	if temp := value / 2; temp > 30 {
		fmt.Printf("temp(%d) 大于 30，在 if 块中可用\n", temp)
		// temp 变量在这里可用
	} else {
		fmt.Printf("temp 在 else 块中也可用，值为: %d\n", temp)
		// temp 变量在这里也可用
	}

	// temp 变量在这里不可用
	// fmt.Println(temp) // 编译错误：undefined: temp

	fmt.Println()
}

// IfStatementDemo if 语句完整演示主函数
func IfStatementDemo() {
	fmt.Println("========== 1.8.1 if 语句 ==========")
	fmt.Println()
	fmt.Println("if 语句由一个或多个布尔表达式组成，且布尔表达式可以不加括号。")
	fmt.Println()
	fmt.Println("基本语法:")
	fmt.Println("if <expression> {")
	fmt.Println("    <do something>")
	fmt.Println("} else {")
	fmt.Println("    <do something else>")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("关键概念:")
	fmt.Println("- if 语句不需要括号包围条件表达式")
	fmt.Println("- 可以在条件表达式前添加初始化语句")
	fmt.Println("- 初始化语句中声明的变量作用域仅限于整个 if-else 块")
	fmt.Println("- 支持嵌套的 if 语句")
	fmt.Println()

	demonstrateBasicIf()
	demonstrateIfElseIf()
	demonstrateNestedIf()
	demonstrateIfWithAssignment()
	demonstrateComplexIf()
	demonstrateIfWithFunctionCall()
	demonstrateBooleanExpressions()
	demonstrateIfScope()
}
