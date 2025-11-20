// Package controlflow 演示 Go 语言 switch 语句的使用
package controlflow

import "fmt"

// CustomType 自定义类型用于演示类型 switch
type CustomType struct {
	Name string
}

// ========== 1.8.2 switch 语句 ==========

// demonstrateBasicSwitch 演示基本的 switch 语句
func demonstrateBasicSwitch() {
	fmt.Println("=== 1. 基本 switch 语句 ===")

	a := "test string"

	// 基本用法
	switch a {
	case "test":
		fmt.Println("a = ", a)
	case "s":
		fmt.Println("a = ", a)
	case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
		fmt.Println("catch in a test, a = ", a)
	case "n":
		fmt.Println("a = not")
	default:
		fmt.Println("default case")
	}

	fmt.Println()
}

// demonstrateSwitchWithAssignment 演示带初始化语句的 switch 语句
func demonstrateSwitchWithAssignment() {
	fmt.Println("=== 2. 带初始化语句的 switch 语句 ===")

	// 变量b仅在当前switch代码块内有效
	switch b := 5; b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	case 3, 4:
		fmt.Println("b = 3 or 4")
	case 5:
		fmt.Println("b = 5")
	default:
		fmt.Println("b = ", b)
	}

	fmt.Println()
}

// demonstrateSwitchWithoutCondition 演示不带条件的 switch 语句
func demonstrateSwitchWithoutCondition() {
	fmt.Println("=== 3. 不带条件的 switch 语句 ===")

	a := "test string"
	b := 5

	// 不指定判断变量，直接在case中添加判定条件
	switch {
	case a == "t":
		fmt.Println("a = t")
	case b == 3:
		fmt.Println("b = 5")
	case b == 5, a == "test string":
		fmt.Println("a = test string; or b = 5")
	default:
		fmt.Println("default case")
	}

	fmt.Println()
}

// demonstrateTypeSwitch 演示类型 switch 语句
func demonstrateTypeSwitch() {
	fmt.Println("=== 4. 类型 switch 语句 ===")

	var d interface{}
	d = 1

	// 类型 switch，用于判断接口变量的实际类型
	switch t := d.(type) {
	case byte:
		fmt.Println("d is byte type, ", t)
	case int:
		fmt.Println("d is int type, ", t)
	case string:
		fmt.Println("d is string type, ", t)
	case CustomType:
		fmt.Println("d is CustomType type, ", t)
	default:
		fmt.Println("d is unknown type, ", t)
	}

	// 演示指针类型
	var e interface{}
	custom := CustomType{Name: "test"}
	e = &custom

	switch t := e.(type) {
	case *CustomType:
		fmt.Println("e is CustomType pointer type, ", t.Name)
	default:
		fmt.Println("e is unknown type, ", t)
	}

	fmt.Println()
}

// demonstrateSwitchFallthrough 演示 fallthrough 关键字
func demonstrateSwitchFallthrough() {
	fmt.Println("=== 5. fallthrough 关键字 ===")

	// fallthrough 会继续执行下一个 case，即使条件不匹配
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
		fallthrough // 继续执行下一个 case
	case "Tuesday":
		fmt.Println("Today is Tuesday")
		fallthrough // 继续执行下一个 case
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Another day")
	}

	fmt.Println()
}

// demonstrateSwitchInLoops 演示循环中的 switch 语句
func demonstrateSwitchInLoops() {
	fmt.Println("=== 6. 循环中的 switch 语句 ===")

	// 在循环中使用 switch
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("i is zero")
		case 1, 2:
			fmt.Println("i is one or two")
		case 3:
			fmt.Println("i is three")
		default:
			fmt.Printf("i is %d\n", i)
		}
	}

	fmt.Println()
}

// demonstrateSwitchWithFunctions 演示 switch 中调用函数
func demonstrateSwitchWithFunctions() {
	fmt.Println("=== 7. switch 中调用函数 ===")

	// 在 switch 条件中调用函数
	switch getDay() {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Friday":
		fmt.Println("End of the work week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Midweek day")
	}

	fmt.Println()
}

// getDay 辅助函数用于演示
func getDay() string {
	return "Monday"
}

// demonstrateSwitchScope 演示 switch 作用域
func demonstrateSwitchScope() {
	fmt.Println("=== 8. switch 作用域 ===")

	// 在 switch 初始化语句中声明的变量只在 switch 块内有效
	switch x := calculateValue(); x {
	case 1:
		fmt.Printf("x is %d in case 1\n", x)
	case 2:
		fmt.Printf("x is %d in case 2\n", x)
	case 3:
		fmt.Printf("x is %d in case 3\n", x)
	default:
		fmt.Printf("x is %d in default case\n", x)
	}

	// x 变量在此处不可访问
	// fmt.Println(x) // 这行代码会编译错误

	fmt.Println()
}

// calculateValue 辅助函数用于演示
func calculateValue() int {
	return 2
}

// SwitchStatementDemo switch 语句完整演示主函数
func SwitchStatementDemo() {
	fmt.Println("========== 1.8.2 switch 语句 ==========")
	fmt.Println()
	fmt.Println("switch 语句用于基于不同条件执行不同的动作。")
	fmt.Println()
	fmt.Println("每个 case 分支都是唯一的，从上往下逐一判断，直到匹配为止。")
	fmt.Println("如果某些 case 条件重复，编译时会报错。")
	fmt.Println()
	fmt.Println("默认情况下 case 分支自带 break 效果，无需在每个 case 中声明 break。")
	fmt.Println()
	fmt.Println("基本语法:")
	fmt.Println("switch <variable> {")
	fmt.Println("case <value1>:")
	fmt.Println("    <do something1>")
	fmt.Println("case <value2>:")
	fmt.Println("    <do something2>")
	fmt.Println("default:")
	fmt.Println("    <do something>")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("关键概念:")
	fmt.Println("- 可以匹配多个值: case <value1>, <value2>")
	fmt.Println("- 支持初始化语句: switch <var> := <expr>; <var> { ... }")
	fmt.Println("- 不带条件的 switch: switch { ... }")
	fmt.Println("- 类型 switch: switch v := x.(type) { ... }")
	fmt.Println("- fallthrough 关键字: 继续执行下一个 case")
	fmt.Println()

	demonstrateBasicSwitch()
	demonstrateSwitchWithAssignment()
	demonstrateSwitchWithoutCondition()
	demonstrateTypeSwitch()
	demonstrateSwitchFallthrough()
	demonstrateSwitchInLoops()
	demonstrateSwitchWithFunctions()
	demonstrateSwitchScope()
}
