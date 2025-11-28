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
// 这是一个经典的 Go 并发模式：带有超时控制和外部停止信号的后台任务循环
// 核心结构：无限循环 (for {}) - "只要没叫我停，我就一直干活"
// 三道"刹车"机制：
// 1. 超时刹车 (Context) - 运行时间太长，强制停
// 2. 任务完成刹车 (Atomic) - 后台任务干完了，主动停
// 3. 强制兜底刹车 (loopCount) - 防止死循环
func demonstrateComplexForLoop() {
	fmt.Println("=== 10. 复杂的 for 循环示例（带注释详解）===")

	// 1. 设置超时控制 (Context)
	// 这里设置了一个 3秒后 会自动过期的上下文（"闹钟"）
	// WithTimeout 返回一个 Context 和一个取消函数 CancelFunc
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel() // 习惯用法：函数结束时清理资源

	var started bool
	// Atomic Bool 用于线程安全地在不同协程间共享状态
	// 这里用来作为"任务完成"的信号旗
	var stopped atomic.Bool

	// 2. 启动后台任务 (Goroutine)
	// 模拟一个异步工作，比如去下载文件或处理数据
	go func() {
		// 模拟干活干了2秒
		time.Sleep(time.Second * 2)
		// 告诉主线程：我干完了！
		stopped.Store(true)
		fmt.Println("【后台】任务完成，设置停止标志")
	}()

	// 3. 主循环 (Main Loop)
	loopCount := 0

	// 定义标签 LoopLabel，用于从 select 中直接跳出外层循环
LoopLabel:
	for {
		loopCount++
		if !started {
			started = true
			fmt.Println("【主循环】开始执行...")
		}

		fmt.Printf("【主循环】第 %d 次轮询检查...\n", loopCount)

		// 检查上下文是否超时 ("瞄一眼闹钟")
		// 使用 select + default 实现非阻塞检查
		select {
		case <-ctx.Done():
			// 3秒时间到了，Context 会发送信号
			fmt.Println("【退出】上下文超时 (Context Timeout)")
			// ⚠️ 注意：在 select 中直接使用 break 只会跳出 select，不会跳出 for 循环
			// ✅ 正确做法：配合标签 (Label) 跳出指定循环
			break LoopLabel
		default:
			// default 分支确保了如果 Context 没超时，代码不会卡在这里，而是继续往下走
			// 这就是"非阻塞"的关键
		}

		// 检查停止标志 ("瞄一眼队友")
		if stopped.Load() {
			fmt.Println("【退出】收到停止信号 (Task Completed)")
			break // 这里在 if 中，可以直接 break 跳出 for 循环
		}

		// 模拟一些工作间隔，避免循环跑得太快占满 CPU
		time.Sleep(time.Millisecond * 500)

		// 强制兜底刹车 (安全措施)
		if loopCount >= 10 {
			fmt.Println("【退出】达到最大循环次数 (Safety Limit)")
			break
		}
	}

	fmt.Println("循环结束")
	fmt.Println()
}

// demonstrateBlankIdentifier 演示空白标识符（_）在 for 循环中的使用
// 空白标识符用于忽略不需要的值，避免"声明但未使用"的编译错误
func demonstrateBlankIdentifier() {
	fmt.Println("=== 11. 空白标识符（_）在 for 循环中的使用 ===")
	fmt.Println("说明：空白标识符 _ 用于忽略不需要的值，避免编译错误")

	// 示例1：遍历数组，只获取值，忽略索引
	fmt.Println("\n--- 示例1：只获取值，忽略索引 ---")
	arr := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	for _, value := range arr {
		fmt.Printf("值: %s\n", value)
	}

	// 示例2：遍历数组，只获取索引，忽略值
	fmt.Println("\n--- 示例2：只获取索引，忽略值 ---")
	for index := range arr {
		fmt.Printf("索引: %d\n", index)
	}

	// 示例3：遍历 map，只获取 key，忽略 value
	fmt.Println("\n--- 示例3：遍历 map，只获取 key，忽略 value ---")
	m := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}
	for key := range m {
		fmt.Printf("key: %s\n", key)
	}

	// 示例4：遍历 map，只获取 value，忽略 key
	fmt.Println("\n--- 示例4：遍历 map，只获取 value，忽略 key ---")
	for _, value := range m {
		fmt.Printf("value: %d\n", value)
	}

	// 示例5：遍历字符串，只获取字符，忽略位置
	fmt.Println("\n--- 示例5：遍历字符串，只获取字符，忽略位置 ---")
	str := "Go语言"
	for _, char := range str {
		fmt.Printf("字符: %c\n", char)
	}

	// 示例6：遍历切片，只获取值，忽略索引
	fmt.Println("\n--- 示例6：遍历切片，只获取值，忽略索引 ---")
	slice := []int{10, 20, 30, 40, 50}
	sum := 0
	for _, num := range slice {
		sum += num
	}
	fmt.Printf("切片元素之和: %d\n", sum)

	// ⚠️ 注意事项
	fmt.Println("\n⚠️ 注意事项：")
	fmt.Println("  - 空白标识符 _ 不能作为变量使用")
	fmt.Println("  - 使用 _ 可以避免'声明但未使用'的编译错误")
	fmt.Println("  - 在 range 循环中，_ 用于忽略不需要的返回值")
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
	fmt.Println("- 空白标识符 _: 用于忽略不需要的值")
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
	demonstrateBlankIdentifier()
}
