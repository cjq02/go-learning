package concurrency

import (
	"fmt"
	"time"
)

// ========== 1.18.2 channel ==========
//
// channel 是 Go 中定义的一种类型，专门用来在多个 goroutine 之间通信的
// 线程安全的数据结构。
//
// 可以在一个 goroutine 中向一个 channel 中发送数据，
// 从另外一个 goroutine 中接收数据。
//
// channel 类似队列，满足先进先出原则。

// ChannelDemo 演示 channel 基本使用
func ChannelDemo() {
	fmt.Println("========== 1.18.2 channel ==========")
	fmt.Println()
	fmt.Println("channel 是 Go 中定义的一种类型，专门用来在多个 goroutine 之间通信的")
	fmt.Println("线程安全的数据结构。")
	fmt.Println()
	fmt.Println("可以在一个 goroutine 中向一个 channel 中发送数据，")
	fmt.Println("从另外一个 goroutine 中接收数据。")
	fmt.Println()
	fmt.Println("channel 类似队列，满足先进先出原则。")
	fmt.Println()

	demonstrateChannelDefinition()
	demonstrateChannelOperations()
	demonstrateBufferedChannel()
	demonstrateChannelDirection()
	demonstrateChannelSelect()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ channel 是线程安全的数据结构，用于 goroutine 间通信")
	fmt.Println("✅ channel 类似队列，满足先进先出（FIFO）原则")
	fmt.Println("✅ 支持发送、接收、关闭三种操作")
	fmt.Println("✅ 可以限制 channel 的方向（只发送或只接收）")
	fmt.Println("✅ 使用 select 可以实现多路复用")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 无缓冲 channel 是同步的，发送和接收会阻塞")
	fmt.Println("   - 有缓冲 channel 是异步的，缓冲区满时发送会阻塞")
	fmt.Println("   - 向已关闭的 channel 发送数据会 panic")
	fmt.Println("   - 从已关闭的 channel 接收数据会立即返回零值和 false")
	fmt.Println("   - 关闭已关闭的 channel 会 panic")
	fmt.Println()
}

// demonstrateChannelDefinition 演示 channel 定义
func demonstrateChannelDefinition() {
	fmt.Println("=== 1.18.2.1 channel 定义 ===")
	fmt.Println()

	fmt.Println("1. 定义方式：")
	fmt.Println()
	fmt.Println("   // 仅声明（未初始化，值为 nil）")
	fmt.Println("   var <channel_name> chan <type_name>")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   var ch chan int")
	fmt.Println()

	fmt.Println("   // 初始化无缓冲 channel（同步 channel）")
	fmt.Println("   <channel_name> := make(chan <type_name>)")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   ch := make(chan int)")
	fmt.Println()

	fmt.Println("   // 初始化有缓冲 channel（异步 channel）")
	fmt.Println("   <channel_name> := make(chan <type_name>, buffer_size)")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   ch := make(chan int, 3)  // 缓冲区大小为 3")
	fmt.Println()

	fmt.Println("2. 无缓冲 vs 有缓冲 channel：")
	fmt.Println()
	fmt.Println("   无缓冲 channel（同步）：")
	fmt.Println("   - 发送操作会阻塞，直到有 goroutine 接收")
	fmt.Println("   - 接收操作会阻塞，直到有 goroutine 发送")
	fmt.Println("   - 保证发送和接收同时发生（同步）")
	fmt.Println()
	fmt.Println("   有缓冲 channel（异步）：")
	fmt.Println("   - 缓冲区未满时，发送不会阻塞")
	fmt.Println("   - 缓冲区不为空时，接收不会阻塞")
	fmt.Println("   - 缓冲区满时，发送会阻塞")
	fmt.Println("   - 缓冲区空时，接收会阻塞")
	fmt.Println()

	fmt.Println("3. 实际示例：")
	fmt.Println()

	// 无缓冲 channel
	unbufferedCh := make(chan int)
	fmt.Printf("   无缓冲 channel: %v\n", unbufferedCh)

	// 有缓冲 channel
	bufferedCh := make(chan int, 3)
	fmt.Printf("   有缓冲 channel (容量3): %v\n", bufferedCh)
	fmt.Println()
}

// demonstrateChannelOperations 演示 channel 操作
func demonstrateChannelOperations() {
	fmt.Println("=== 1.18.2.2 channel 操作 ===")
	fmt.Println()

	fmt.Println("channel 的三种操作：发送数据、接收数据、关闭通道")
	fmt.Println()

	fmt.Println("1. 发送数据：")
	fmt.Println("   <channel_name> <- <variable_name_or_value>")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   ch <- 10        // 发送值 10")
	fmt.Println("   ch <- value    // 发送变量 value")
	fmt.Println()

	fmt.Println("2. 接收数据：")
	fmt.Println("   // 方式1：接收值和一个标志（表示 channel 是否关闭）")
	fmt.Println("   value_name, ok_flag := <- <channel_name>")
	fmt.Println()
	fmt.Println("   // 方式2：只接收值")
	fmt.Println("   value_name := <- <channel_name>")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   value, ok := <-ch  // ok 为 true 表示 channel 未关闭")
	fmt.Println("   value := <-ch      // 只接收值")
	fmt.Println()

	fmt.Println("3. 关闭 channel：")
	fmt.Println("   close(<channel_name>)")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   close(ch)")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 关闭 channel 后，不能再发送数据（会 panic）")
	fmt.Println("   - 可以继续接收数据，直到 channel 为空")
	fmt.Println("   - 接收完所有数据后，再接收会返回零值和 false")
	fmt.Println()

	fmt.Println("4. 实际运行示例：")
	fmt.Println()

	ch := make(chan int, 3)

	// 启动发送 goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("   [发送] 发送数据: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
		fmt.Println("   [发送] channel 已关闭")
	}()

	// 等待一下，确保发送开始
	time.Sleep(50 * time.Millisecond)

	// 接收数据
	fmt.Println("   [接收] 开始接收数据：")
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("   [接收] channel 已关闭，接收完成")
			break
		}
		fmt.Printf("   [接收] 接收到数据: %d\n", value)
	}
	fmt.Println()
}

// demonstrateBufferedChannel 演示有缓冲 channel
func demonstrateBufferedChannel() {
	fmt.Println("=== 1.18.2.3 有缓冲 channel ===")
	fmt.Println()

	fmt.Println("1. 有缓冲 channel 的特点：")
	fmt.Println("   - 缓冲区未满时，发送不会阻塞")
	fmt.Println("   - 缓冲区不为空时，接收不会阻塞")
	fmt.Println("   - 可以实现异步通信")
	fmt.Println()

	fmt.Println("2. 实际运行示例：")
	fmt.Println()

	ch := make(chan int, 3)

	fmt.Println("   创建容量为 3 的缓冲 channel")
	fmt.Println()

	// 发送数据（不会阻塞，因为缓冲区未满）
	fmt.Println("   发送数据到缓冲区：")
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Printf("   发送: %d (缓冲区未满，不阻塞)\n", i)
	}
	fmt.Println()

	// 尝试再发送一个（会阻塞，因为缓冲区已满）
	fmt.Println("   尝试发送第 4 个数据（缓冲区已满，会阻塞）：")
	go func() {
		ch <- 4
		fmt.Println("   发送: 4 (缓冲区有空间了)")
	}()

	// 接收数据
	fmt.Println("   开始接收数据：")
	for i := 0; i < 4; i++ {
		value := <-ch
		fmt.Printf("   接收: %d\n", value)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

// receiveOnly 只接收 channel 的函数
func receiveOnly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("   [只接收函数] 接收到: %d\n", v)
	}
	fmt.Println("   [只接收函数] channel 已关闭，接收完成")
}

// sendOnly 只发送 channel 的函数
func sendOnly(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("   [只发送函数] 发送: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
	fmt.Println("   [只发送函数] channel 已关闭")
}

// demonstrateChannelDirection 演示 channel 方向性
func demonstrateChannelDirection() {
	fmt.Println("=== 1.18.2.4 channel 方向性 ===")
	fmt.Println()

	fmt.Println("channel 还有两个变种，可以把 channel 作为参数传递时，")
	fmt.Println("限制 channel 在函数或方法中能够执行的操作。")
	fmt.Println()

	fmt.Println("1. 只发送 channel（chan<- type）：")
	fmt.Println("   func <method_name>(<channel_name> chan<- <type>)")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 函数只能向 channel 发送数据")
	fmt.Println("   - 不能从 channel 接收数据")
	fmt.Println("   - 可以关闭 channel")
	fmt.Println()

	fmt.Println("2. 只接收 channel（<-chan type）：")
	fmt.Println("   func <method_name>(<channel_name> <-chan <type>)")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 函数只能从 channel 接收数据")
	fmt.Println("   - 不能向 channel 发送数据")
	fmt.Println("   - 不能关闭 channel")
	fmt.Println()

	fmt.Println("3. 实际运行示例：")
	fmt.Println()

	// 创建一个带缓冲的 channel
	ch := make(chan int, 3)

	// 启动发送 goroutine（使用只发送函数）
	go sendOnly(ch)

	// 启动接收 goroutine（使用只接收函数）
	go receiveOnly(ch)

	// 等待 goroutine 完成
	time.Sleep(800 * time.Millisecond)
	fmt.Println()

	fmt.Println("4. 方向性的优势：")
	fmt.Println("   ✅ 类型安全：编译时检查，防止误操作")
	fmt.Println("   ✅ 代码清晰：明确函数的职责")
	fmt.Println("   ✅ 接口设计：可以设计更清晰的 API")
	fmt.Println()
}

// demonstrateChannelSelect 演示 select 多路复用
func demonstrateChannelSelect() {
	fmt.Println("=== 1.18.2.5 select 多路复用 ===")
	fmt.Println()

	fmt.Println("select 语句可以让 goroutine 同时等待多个 channel 操作。")
	fmt.Println("它会阻塞直到某个 case 可以执行，然后执行该 case。")
	fmt.Println()

	fmt.Println("1. select 语法：")
	fmt.Println("   select {")
	fmt.Println("   case value := <-ch1:")
	fmt.Println("       // 处理 ch1 的数据")
	fmt.Println("   case value := <-ch2:")
	fmt.Println("       // 处理 ch2 的数据")
	fmt.Println("   case ch3 <- value:")
	fmt.Println("       // 向 ch3 发送数据")
	fmt.Println("   case <-timeout:")
	fmt.Println("       // 超时处理")
	fmt.Println("   default:")
	fmt.Println("       // 所有 case 都不满足时执行（非阻塞）")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("2. 实际运行示例：")
	fmt.Println()

	ch := make(chan int, 3)

	// 启动发送 goroutine
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			fmt.Printf("   [发送] 发送: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		close(ch)
	}()

	// 使用 select 进行多路复用
	timeout := time.After(2 * time.Second)
	receivedCount := 0

	fmt.Println("   使用 select 接收数据（带超时）：")
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("   [select] Channel 已关闭")
				fmt.Println()
				goto endSelect
			}
			fmt.Printf("   [select] 接收到: %d\n", v)
			receivedCount++
			if receivedCount >= 3 {
				goto endSelect
			}
		case <-timeout:
			fmt.Println("   [select] 操作超时")
			goto endSelect
		default:
			fmt.Println("   [select] 没有数据，等待中...")
			time.Sleep(100 * time.Millisecond)
		}
	}

endSelect:
	fmt.Println("3. select 的特点：")
	fmt.Println("   ✅ 可以同时监听多个 channel")
	fmt.Println("   ✅ 随机选择一个就绪的 case 执行")
	fmt.Println("   ✅ 如果没有 case 就绪，执行 default（如果有）")
	fmt.Println("   ✅ 如果没有 default，会阻塞直到某个 case 就绪")
	fmt.Println("   ✅ 常用于超时控制和多路复用")
	fmt.Println()
}
