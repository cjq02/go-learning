// Package function 演示 Go 语言闭包的使用
package function

import "fmt"

// ========== 1.10.2 闭包 ==========

// demonstrateBasicAnonymousFunction 演示最基本的匿名函数
func demonstrateBasicAnonymousFunction() {
	fmt.Println("=== 1. 最基本的匿名函数 ===")

	// 声明并立即执行
	func() {
		fmt.Println("这是一个匿名函数")
	}()

	fmt.Println()
}

// demonstrateFunctionVariable 演示函数变量的使用
func demonstrateFunctionVariable() {
	fmt.Println("=== 2. 函数变量 ===")

	// 声明函数变量
	var add func(int, int) int

	// 分配匿名函数给变量
	add = func(a, b int) int {
		return a + b
	}

	result := add(5, 3)
	fmt.Printf("add(5, 3) = %d\n", result)

	fmt.Println()
}

// demonstrateClosureCapture 演示闭包捕获外部变量
func demonstrateClosureCapture() {
	fmt.Println("=== 3. 闭包捕获外部变量（重要！）===")

	// 外部变量
	x := 10

	// 匿名函数可以直接访问外部变量 x
	printX := func() {
		fmt.Printf("外部变量 x = %d\n", x)
	}

	printX()

	// 修改外部变量
	x = 20
	fmt.Printf("修改 x 为 %d 后：\n", x)
	printX() // 输出新值

	fmt.Println()
}

// demonstrateClosureModifyExternal 演示闭包修改外部变量
func demonstrateClosureModifyExternal() {
	fmt.Println("=== 4. 闭包修改外部变量 ===")

	// 外部变量（计数器）
	count := 0

	// 增加计数的匿名函数
	increment := func() {
		count++
	}

	fmt.Printf("初始 count = %d\n", count)
	increment()
	fmt.Printf("执行 increment() 后，count = %d\n", count)
	increment()
	fmt.Printf("再执行一次，count = %d\n", count)

	fmt.Println()
}

// demonstrateMultipleClosures 演示多个匿名函数共享同一变量
func demonstrateMultipleClosures() {
	fmt.Println("=== 5. 多个匿名函数共享同一变量 ===")

	message := "Hello"

	// 匿名函数1：打印消息
	print := func() {
		fmt.Printf("消息: %s\n", message)
	}

	// 匿名函数2：修改消息
	change := func(newMsg string) {
		message = newMsg
	}

	print()         // 输出: Hello
	change("World") // 修改消息
	print()         // 输出: World

	fmt.Println()
}

// demonstrateAnonymousAsParameter 演示匿名函数作为参数
func demonstrateAnonymousAsParameter() {
	fmt.Println("=== 6. 匿名函数作为参数 ===")

	// 接受函数作为参数的函数
	apply := func(x, y int, op func(int, int) int) int {
		return op(x, y)
	}

	// 传递不同的匿名函数
	result1 := apply(10, 5, func(a, b int) int {
		return a + b // 加法
	})
	fmt.Printf("10 + 5 = %d\n", result1)

	result2 := apply(10, 5, func(a, b int) int {
		return a - b // 减法
	})
	fmt.Printf("10 - 5 = %d\n", result2)

	fmt.Println()
}

// demonstrateAnonymousAsReturnValue 演示匿名函数作为返回值
func demonstrateAnonymousAsReturnValue() {
	fmt.Println("=== 7. 匿名函数作为返回值 ===")

	// 创建一个返回匿名函数的函数
	makeAdder := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}

	// 创建 "加5" 的函数
	add5 := makeAdder(5)
	fmt.Printf("add5(3) = %d\n", add5(3))   // 5 + 3 = 8
	fmt.Printf("add5(10) = %d\n", add5(10)) // 5 + 10 = 15

	// 创建 "加100" 的函数
	add100 := makeAdder(100)
	fmt.Printf("add100(1) = %d\n", add100(1)) // 100 + 1 = 101

	fmt.Println()
}

// demonstrateCounterFactory 演示计数器工厂（常用模式）
func demonstrateCounterFactory() {
	fmt.Println("=== 8. 计数器工厂（常用模式）===")

	// 创建一个计数器
	makeCounter := func() func() int {
		count := 0 // 每个计数器都有自己的 count
		return func() int {
			count++
			return count
		}
	}

	// 创建两个独立的计数器
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println("计数器1的调用结果:")
	fmt.Printf("  第1次: %d\n", counter1())
	fmt.Printf("  第2次: %d\n", counter1())
	fmt.Printf("  第3次: %d\n", counter1())

	fmt.Println("计数器2的调用结果（独立计数）:")
	fmt.Printf("  第1次: %d\n", counter2())
	fmt.Printf("  第2次: %d\n", counter2())

	fmt.Println()
}

// demonstrateLoopClosureProblem 演示循环中的闭包陷阱
func demonstrateLoopClosureProblem() {
	fmt.Println("=== 9. 循环中的闭包陷阱 ===")

	fmt.Println("问题：所有闭包都引用同一个变量")
	funcs := make([]func(), 0)
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Printf("  输出: %d\n", i)
		})
	}

	fmt.Println("执行结果（所有输出都是3！）:")
	for _, f := range funcs {
		f()
	}

	fmt.Println()
	fmt.Println("解决方案1：传递参数")
	funcs1 := make([]func(), 0)
	for i := 0; i < 3; i++ {
		funcs1 = append(funcs1, func(val int) func() {
			return func() {
				fmt.Printf("  输出: %d\n", val)
			}
		}(i))
	}

	fmt.Println("执行结果（正确输出 0, 1, 2）:")
	for _, f := range funcs1 {
		f()
	}

	fmt.Println()
	fmt.Println("解决方案2：创建局部变量")
	funcs2 := make([]func(), 0)
	for i := 0; i < 3; i++ {
		val := i // 创建局部变量
		funcs2 = append(funcs2, func() {
			fmt.Printf("  输出: %d\n", val)
		})
	}

	fmt.Println("执行结果（正确输出 0, 1, 2）:")
	for _, f := range funcs2 {
		f()
	}

	fmt.Println()
}

// ClosureDemo 闭包完整演示主函数
func ClosureDemo() {
	fmt.Println("========== 1.10.2 闭包 ==========")
	fmt.Println()
	fmt.Println("闭包，也被称为匿名函数，顾名思义，即没有函数名。")
	fmt.Println("通常在函数内或方法内定义，或作为参数、返回值传递。")
	fmt.Println()
	fmt.Println("闭包的优势：可以直接访问和修改外部作用域中的变量。")
	fmt.Println()

	demonstrateBasicAnonymousFunction()
	demonstrateFunctionVariable()
	demonstrateClosureCapture()
	demonstrateClosureModifyExternal()
	demonstrateMultipleClosures()
	demonstrateAnonymousAsParameter()
	demonstrateAnonymousAsReturnValue()
	demonstrateCounterFactory()
	demonstrateLoopClosureProblem()
}
