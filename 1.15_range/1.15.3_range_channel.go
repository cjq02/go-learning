package rangeiteration

import (
	"fmt"
	"time"
)

// ========== 1.15.3 对通道迭代 ==========
//
// 通道除了可以使用 for 循环配合 select 关键字获取数据以外，
// 也可以使用 for 循环配合 range 关键字获取数据
//
// 因为通道结构的特殊性，当使用 range 遍历通道时，
// 只给一个迭代变量赋值，而不像数组或字符串一样能够使用 index 索引
//
// 当通道被关闭时，在 range 关键字迭代完通道中所有值后，循环就会自动退出

// addData 向通道添加数据（辅助函数）
func addData(ch chan int) {
	size := cap(ch)
	for i := 0; i < size; i++ {
		ch <- i
		fmt.Printf("  发送: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 缩短等待时间以便演示
	}
	close(ch)
	fmt.Println("  通道已关闭")
}

// demonstrateChannelRangeBasic 演示通道 range 迭代的基本用法
func demonstrateChannelRangeBasic() {
	fmt.Println("=== 1.15.3.1 通道 range 迭代的基本用法 ===")

	ch := make(chan int, 10)

	// 启动 goroutine 向通道发送数据
	go addData(ch)

	// 使用 range 遍历通道
	fmt.Println("使用 range 遍历通道：")
	for i := range ch {
		fmt.Printf("  接收到值: %d\n", i)
	}
	fmt.Println("循环自动退出（通道已关闭且所有值已读取）")
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - range 遍历通道时，只有一个迭代变量（值），没有索引")
	fmt.Println("  - 当通道关闭且所有值被读取后，循环自动退出")
	fmt.Println("  - 不需要手动检查通道是否关闭")
	fmt.Println()
}

// demonstrateChannelRangeVsSelect 对比 range 和 select 的使用
//
//nolint:all // 此函数用于演示 select 用法，故意不使用 range
func demonstrateChannelRangeVsSelect() {
	fmt.Println("=== range vs select 对比 ===")

	ch := make(chan int, 5)

	// 方式1：使用 range（推荐，简洁）
	fmt.Println("--- 方式1：使用 range（推荐）---")
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	fmt.Println("使用 range 遍历：")
	for value := range ch {
		fmt.Printf("  接收到值: %d\n", value)
	}
	fmt.Println()

	// 方式2：使用 select（需要手动检查）
	fmt.Println("--- 方式2：使用 select（需要手动检查）---")
	ch2 := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	fmt.Println("使用 select 遍历：")
	// 这里使用 select 是为了演示对比，实际场景中推荐使用 range
	//nolint // 演示 select 用法，不使用 range
	for {
		select {
		case value, ok := <-ch2:
			if !ok {
				fmt.Println("  通道已关闭，退出循环")
				goto end
			}
			fmt.Printf("  接收到值: %d\n", value)
		}
	}
end:
	fmt.Println()

	fmt.Println("对比总结：")
	fmt.Println("  - range：简洁，自动处理通道关闭")
	fmt.Println("  - select：灵活，可以处理多个通道，但需要手动检查关闭")
	fmt.Println()
}

// demonstrateChannelRangeWithMultipleChannels 演示多个通道的情况
func demonstrateChannelRangeWithMultipleChannels() {
	fmt.Println("=== 多个通道的情况 ===")

	ch1 := make(chan int, 3)
	ch2 := make(chan string, 3)

	// 向通道发送数据
	go func() {
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		close(ch1)
	}()

	go func() {
		ch2 <- "a"
		ch2 <- "b"
		ch2 <- "c"
		close(ch2)
	}()

	// 使用 range 遍历多个通道（需要分别遍历）
	fmt.Println("遍历 ch1：")
	for value := range ch1 {
		fmt.Printf("  ch1: %d\n", value)
	}

	fmt.Println("遍历 ch2：")
	for value := range ch2 {
		fmt.Printf("  ch2: %s\n", value)
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - range 一次只能遍历一个通道")
	fmt.Println("  - 如果需要处理多个通道，需要使用 select")
	fmt.Println()
}

// demonstrateChannelRangeBlocking 演示通道 range 的阻塞特性
func demonstrateChannelRangeBlocking() {
	fmt.Println("=== 通道 range 的阻塞特性 ===")

	ch := make(chan int, 2)

	// 先发送一些数据
	fmt.Println("先发送一些数据")
	fmt.Println("发送数据：ch <- 1")
	ch <- 1
	fmt.Println("发送数据：ch <- 2")
	ch <- 2

	fmt.Println("启动 goroutine 继续发送数据")
	// 启动 goroutine 继续发送数据
	go func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("发送数据：ch <- 3")
		ch <- 3
		fmt.Println("发送数据：ch <- 4")
		ch <- 4
		time.Sleep(200 * time.Millisecond)
		fmt.Println("关闭通道：close(ch)")
		close(ch)
	}()

	fmt.Println("range 遍历通道（会阻塞等待数据）：")
	for value := range ch {
		fmt.Printf("  接收到值: %d\n", value)
	}
	fmt.Println("循环退出（通道已关闭）")
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - range 遍历通道时会阻塞，等待数据")
	fmt.Println("  - 如果通道中没有数据，会一直等待")
	fmt.Println("  - 只有当通道关闭且所有数据读取完毕，循环才会退出")
	fmt.Println()
}

// demonstrateChannelRangeUnbuffered 演示无缓冲通道的 range 迭代
func demonstrateChannelRangeUnbuffered() {
	fmt.Println("=== 无缓冲通道的 range 迭代 ===")

	ch := make(chan int) // 无缓冲通道

	// 启动 goroutine 发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("  发送: %d\n", i)
		}
		close(ch)
	}()

	fmt.Println("遍历无缓冲通道：")
	for value := range ch {
		fmt.Printf("  接收: %d\n", value)
		time.Sleep(50 * time.Millisecond) // 模拟处理时间
	}
	fmt.Println("循环退出")
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 无缓冲通道：发送和接收必须同时准备好")
	fmt.Println("  - range 遍历无缓冲通道时，会阻塞等待发送方")
	fmt.Println("  - 通道关闭后，range 循环自动退出")
	fmt.Println()
}

// demonstrateChannelRangeErrorHandling 演示通道 range 的错误处理
func demonstrateChannelRangeErrorHandling() {
	fmt.Println("=== 通道 range 的错误处理 ===")

	ch := make(chan int, 3)

	// 发送数据并关闭
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	fmt.Println("使用 range 遍历（自动处理关闭）：")
	for value := range ch {
		fmt.Printf("  值: %d\n", value)
	}
	fmt.Println("循环正常退出")

	// 尝试从已关闭的通道读取
	fmt.Println("\n尝试从已关闭的通道读取：")
	value, ok := <-ch
	fmt.Printf("value = %d, ok = %v (通道已关闭，ok=false)\n", value, ok)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - range 会自动处理通道关闭，无需手动检查")
	fmt.Println("  - 如果手动读取已关闭的通道，ok=false 表示通道已关闭")
	fmt.Println("  - 已关闭的通道会返回零值")
	fmt.Println()
}

// RangeChannelDemo range 迭代通道完整演示
func RangeChannelDemo() {
	fmt.Println("========== 1.15.3 对通道迭代 ==========")
	fmt.Println()
	fmt.Println("通道除了可以使用 for 循环配合 select 关键字获取数据以外，")
	fmt.Println("也可以使用 for 循环配合 range 关键字获取数据。")
	fmt.Println()
	fmt.Println("因为通道结构的特殊性，当使用 range 遍历通道时，")
	fmt.Println("只给一个迭代变量赋值，而不像数组或字符串一样能够使用 index 索引。")
	fmt.Println()
	fmt.Println("当通道被关闭时，在 range 关键字迭代完通道中所有值后，")
	fmt.Println("循环就会自动退出。")
	fmt.Println()

	demonstrateChannelRangeBasic()
	demonstrateChannelRangeVsSelect()
	demonstrateChannelRangeWithMultipleChannels()
	demonstrateChannelRangeBlocking()
	demonstrateChannelRangeUnbuffered()
	demonstrateChannelRangeErrorHandling()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 通道使用 range 迭代：for value := range ch")
	fmt.Println("✅ 特点：")
	fmt.Println("   - 只有一个迭代变量（值），没有索引")
	fmt.Println("   - 通道关闭后，循环自动退出")
	fmt.Println("   - 会阻塞等待数据")
	fmt.Println()
	fmt.Println("✅ 适用场景：")
	fmt.Println("   - 单个通道的遍历（推荐）")
	fmt.Println("   - 需要自动处理通道关闭")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - range 一次只能遍历一个通道")
	fmt.Println("   - 多个通道需要使用 select")
	fmt.Println("   - 无缓冲通道会阻塞等待发送方")
	fmt.Println()
}
