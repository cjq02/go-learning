// Package loopcontrol 演示 Go 语言 for 循环的使用
package loopcontrol

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// ========== 1.9.1 for 循环 ==========

// demonstrateBasicForLoop 演示基本的 for 循环
func demonstrateBasicForLoop() {
	fmt.Println("=== 1. 基本的 for 循环 ===")

	// 方式1: 初始化语句; 条件表达式; 后置语句
	for i := 0; i < 5; i++ {
		fmt.Printf("方式1，第 %d 次循环\n", i+1)
	}

	fmt.Println()
}

// demonstrateConditionOnlyForLoop 演示仅有条件表达式的 for 循环
func demonstrateConditionOnlyForLoop() {
	fmt.Println("=== 2. 仅有条件表达式的 for 循环 ===")

	// 方式2: 仅有条件表达式（类似 while 循环）
	counter := 1
	for counter <= 3 {
		fmt.Printf("方式2，第 %d 次循环\n", counter)
		counter++
	}

	fmt.Println()
}

// demonstrateInfiniteForLoop 演示无限 for 循环
func demonstrateInfiniteForLoop() {
	fmt.Println("=== 3. 无限 for 循环 ===")

	// 方式3: 无限循环（需要使用 break 语句退出）
	count := 0
	for {
		count++
		fmt.Printf("无限循环第 %d 次\n", count)

		if count >= 3 {
			fmt.Println("退出无限循环")
			break // 使用 break 退出循环
		}
	}

	fmt.Println()
}

// demonstrateArrayRangeLoop 演示数组的 range 循环
func demonstrateArrayRangeLoop() {
	fmt.Println("=== 4. 数组的 range 循环 ===")

	// 创建并初始化数组
	var arr [5]string
	arr[0] = "Hello"
	arr[1] = "World"
	arr[2] = "Go"
	arr[3] = "Language"
	arr[4] = "!"

	// 仅遍历下标
	fmt.Println("仅遍历下标:")
	for i := range arr {
		fmt.Printf("arr[%d] = %s\n", i, arr[i])
	}

	fmt.Println()

	// 同时遍历下标和元素
	fmt.Println("同时遍历下标和元素:")
	for i, element := range arr {
		fmt.Printf("arr[%d] = %s\n", i, element)
	}

	fmt.Println()
}

// demonstrateSliceRangeLoop 演示切片的 range 循环
func demonstrateSliceRangeLoop() {
	fmt.Println("=== 5. 切片的 range 循环 ===")

	// 创建并初始化切片
	slice := make([]string, 5)
	slice[0] = "Apple"
	slice[1] = "Banana"
	slice[2] = "Cherry"
	slice[3] = "Date"
	slice[4] = "Elderberry"

	// 仅遍历下标
	fmt.Println("仅遍历下标:")
	for i := range slice {
		fmt.Printf("slice[%d] = %s\n", i, slice[i])
	}

	fmt.Println()

	// 同时遍历下标和元素
	fmt.Println("同时遍历下标和元素:")
	for i, element := range slice {
		fmt.Printf("slice[%d] = %s\n", i, element)
	}

	fmt.Println()
}

// demonstrateMapRangeLoop 演示 map 的 range 循环
func demonstrateMapRangeLoop() {
	fmt.Println("=== 6. map 的 range 循环 ===")

	// 创建并初始化 map
	m := make(map[string]string)
	m["a"] = "Hello, a"
	m["b"] = "Hello, b"
	m["c"] = "Hello, c"
	m["d"] = "Hello, d"

	// 仅遍历 key
	fmt.Println("仅遍历 key:")
	for key := range m {
		fmt.Printf("key: %s\n", key)
	}

	fmt.Println()

	// 同时遍历 key 和 value
	fmt.Println("同时遍历 key 和 value:")
	for key, value := range m {
		fmt.Printf("m[%s] = %s\n", key, value)
	}

	fmt.Println()
}

// demonstrateStringRangeLoop 演示字符串的 range 循环
func demonstrateStringRangeLoop() {
	fmt.Println("=== 7. 字符串的 range 循环 ===")

	str := "Go语言"

	// 遍历字符串（按 rune，即 Unicode 码点）
	fmt.Println("遍历字符串:")
	for i, char := range str {
		fmt.Printf("位置 %d: %c (Unicode: %U)\n", i, char, char)
	}

	fmt.Println()
}

// demonstrateNestedForLoop 演示嵌套的 for 循环
func demonstrateNestedForLoop() {
	fmt.Println("=== 8. 嵌套的 for 循环 ===")

	// 九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println()
}

// demonstrateForLoopControlStatements 演示循环控制语句
func demonstrateForLoopControlStatements() {
	fmt.Println("=== 9. 循环控制语句 ===")

	fmt.Println("使用 break 跳出循环:")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			fmt.Println("跳出循环")
			break
		}
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println()

	fmt.Println("使用 continue 跳过当前迭代:")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("跳过偶数 %d\n", i)
			continue
		}
		fmt.Printf("奇数: %d\n", i)
	}

	fmt.Println()
}

// demonstrateComplexForLoop 演示复杂的 for 循环示例
func demonstrateComplexForLoop() {
	fmt.Println("=== 10. 复杂的 for 循环示例 ===")

	// 模拟带有超时控制的循环
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var started bool
	var stopped atomic.Bool

	// 创建一个 goroutine 来设置停止标志
	go func() {
		time.Sleep(time.Second * 2)
		stopped.Store(true)
		fmt.Println("后台任务完成，设置停止标志")
	}()

	// 主循环
	loopCount := 0
	for {
		loopCount++
		if !started {
			started = true
			fmt.Println("开始循环")
		}

		fmt.Printf("主循环第 %d 次\n", loopCount)

		// 检查上下文是否超时
		select {
		case <-ctx.Done():
			fmt.Println("上下文超时")
			break
		default:
			// 非阻塞检查
		}

		// 检查停止标志
		if stopped.Load() {
			fmt.Println("收到停止信号")
			break
		}

		// 避免循环过快
		time.Sleep(time.Millisecond * 500)

		// 防止无限循环（安全措施）
		if loopCount >= 10 {
			fmt.Println("达到最大循环次数，退出")
			break
		}
	}

	fmt.Println("循环结束")
	fmt.Println()
}

// ForLoopDemo for 循环完整演示主函数
func ForLoopDemo() {
	fmt.Println("========== 1.9.1 for 循环 ==========")
	fmt.Println()
	fmt.Println("for 循环是 Go 语言中唯一的循环结构，可以用于多种场景。")
	fmt.Println()
	fmt.Println("基本语法:")
	fmt.Println("for <init>; <condition>; <post> {")
	fmt.Println("    <expression>")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("四种声明方式:")
	fmt.Println("1. 标准 for 循环: for i := 0; i < 10; i++ { ... }")
	fmt.Println("2. 条件 for 循环: for condition { ... }")
	fmt.Println("3. 无限 for 循环: for { ... }")
	fmt.Println("4. range for 循环: for key, value := range collection { ... }")
	fmt.Println()
	fmt.Println("关键概念:")
	fmt.Println("- break: 跳出当前循环")
	fmt.Println("- continue: 跳过当前迭代，继续下一次循环")
	fmt.Println("- range: 用于遍历数组、切片、map 和字符串")
	fmt.Println()

	demonstrateBasicForLoop()
	demonstrateConditionOnlyForLoop()
	demonstrateInfiniteForLoop()
	demonstrateArrayRangeLoop()
	demonstrateSliceRangeLoop()
	demonstrateMapRangeLoop()
	demonstrateStringRangeLoop()
	demonstrateNestedForLoop()
	demonstrateForLoopControlStatements()
	demonstrateComplexForLoop()
}
