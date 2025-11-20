// Package function 演示 Go 语言函数的使用
package function

import "fmt"

// ========== 1.10.1 函数基础 ==========

// demonstrateBasicFunction 演示没有参数和返回值的简单函数
func demonstrateBasicFunction() {
	fmt.Println("=== 函数的三个主要部分 ===")
	fmt.Println("函数定义的一般形式：")
	fmt.Println("func <function_name>(<parameter list>) (<return types>) {")
	fmt.Println("    <expressions>")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("其中：")
	fmt.Println("- 名称是必须的")
	fmt.Println("- 参数列表是可选的")
	fmt.Println("- 返回类型列表是可选的")
	fmt.Println()

	fmt.Println("=== 示例 1: 无参数无返回值的函数 ===")
	custom()
}

// custom 演示没有参数和返回值的简单函数
func custom() {
	fmt.Println("你好，世界！")
}

// demonstrateFunctionWithParameters 演示带参数的函数
func demonstrateFunctionWithParameters() {
	fmt.Println("\n=== 示例 2: 有参数的函数 ===")
	fmt.Println("函数可以接收多个参数，参数类型必须明确指定。")
	fmt.Println()

	a, b := 5, 3
	sum := add(a, b)
	fmt.Printf("add(%d, %d) = %d\n", a, b, sum)

	fmt.Println()
	fmt.Println("=== 示例 3: 参数但无返回值的函数 ===")
	greet("Alice")
	greet("Bob")
}

// add 演示带参数和返回值的函数
func add(a, b int) int {
	return a + b
}

// greet 演示带参数但没有返回值的函数
func greet(name string) {
	fmt.Printf("你好，%s!\n", name)
}

// demonstrateFunctionWithReturnValues 演示有返回值的函数
func demonstrateFunctionWithReturnValues() {
	fmt.Println("\n=== 示例 4: 有返回值的函数 ===")
	fmt.Println("函数可以返回一个或多个值。")
	fmt.Println()

	// 单个返回值
	fmt.Println("1. 单个返回值：")
	sum := add(10, 20)
	fmt.Printf("add(10, 20) 的返回值: %d\n", sum)

	// 多个返回值
	fmt.Println("\n2. 多个返回值（通常用于返回值和错误）：")
	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("divide(10.0, 2.0) = %.2f\n", result)
	}

	result, err = divide(10.0, 0.0)
	if err != nil {
		fmt.Printf("divide(10.0, 0.0) 错误: %v\n", err)
	} else {
		fmt.Printf("divide(10.0, 0.0) = %.2f\n", result)
	}
}

// divide 演示有多个返回值的函数
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("不能除以零")
	}
	return a / b, nil
}

// demonstrateMultipleReturnValues 演示多个返回值的交换函数
func demonstrateMultipleReturnValues() {
	fmt.Println("\n=== 示例 5: 交换两个值 ===")
	fmt.Println("Go 语言的特色是可以方便地返回多个值。")
	fmt.Println()

	a, b := "hello", "world"
	fmt.Printf("交换前: a = %s, b = %s\n", a, b)

	a, b = swap(a, b)
	fmt.Printf("交换后: a = %s, b = %s\n", a, b)

	fmt.Println("\n=== 利用多返回值进行变量交换 ===")
	x, y := 10, 20
	fmt.Printf("交换前: x = %d, y = %d\n", x, y)
	x, y = y, x // 直接交换
	fmt.Printf("交换后: x = %d, y = %d\n", x, y)
}

// swap 演示返回多个值的函数
func swap(a, b string) (string, string) {
	return b, a
}

// demonstrateNamedReturnValues 演示命名返回值
func demonstrateNamedReturnValues() {
	fmt.Println("\n=== 示例 6: 命名返回值 ===")
	fmt.Println("函数可以给返回值起名字，这样可以直接使用，不需要 return 后跟具体值。")
	fmt.Println()

	area := calculateArea(5.0, 3.0)
	fmt.Printf("calculateArea(5.0, 3.0) = %.2f\n", area)

	area = calculateArea(10.0, 8.0)
	fmt.Printf("calculateArea(10.0, 8.0) = %.2f\n", area)

	fmt.Println("\n说明：")
	fmt.Println("- 命名返回值自动初始化为零值")
	fmt.Println("- 直接使用 return 时会返回命名的返回值（裸返回）")
	fmt.Println("- 虽然方便，但过多使用会降低代码可读性")
}

// calculateArea 演示带有命名返回值的函数
func calculateArea(length, width float64) (area float64) {
	area = length * width
	return // 裸返回
}

// demonstrateFunctionBestPractices 演示函数使用最佳实践
func demonstrateFunctionBestPractices() {
	fmt.Println("\n=== 函数使用最佳实践 ===")
	fmt.Println()

	fmt.Println("1. 清晰的函数名字：")
	fmt.Println("   - 使用有意义的名字")
	fmt.Println("   - calculateArea 比 calc 更清晰")
	fmt.Println("   - 首字母大写表示可导出（公开）函数")
	fmt.Println()

	fmt.Println("2. 合理的参数个数：")
	fmt.Println("   - 建议参数不超过 3-4 个")
	fmt.Println("   - 如果参数过多，考虑使用结构体")
	fmt.Println()

	fmt.Println("3. 正确的错误处理：")
	fmt.Println("   - 多返回值时，通常最后一个是 error")
	fmt.Println("   - 示例: (result, error)")
	fmt.Println()

	fmt.Println("4. 一致的返回值类型：")
	fmt.Println("   - 避免混淆，保持返回类型一致")
	fmt.Println("   - 使用接口来提高灵活性")
	fmt.Println()

	fmt.Println("5. 避免副作用：")
	fmt.Println("   - 函数应该尽可能是纯函数")
	fmt.Println("   - 尽量减少修改全局变量")
}

// FunctionsDemo 函数完整演示主函数
func FunctionsDemo() {
	fmt.Println("========== 1.10.1 函数 ==========")
	fmt.Println()
	fmt.Println("函数是 Go 语言中重要的编程单元。")
	fmt.Println("函数只有三个主要部分：")
	fmt.Println("1. 函数名 - 必需")
	fmt.Println("2. 参数列表 - 可选")
	fmt.Println("3. 返回类型列表 - 可选")
	fmt.Println()

	demonstrateBasicFunction()
	demonstrateFunctionWithParameters()
	demonstrateFunctionWithReturnValues()
	demonstrateMultipleReturnValues()
	demonstrateNamedReturnValues()
	demonstrateFunctionBestPractices()
}
