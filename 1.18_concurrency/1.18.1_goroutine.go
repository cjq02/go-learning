package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// ========== 1.18.1 goroutine ==========
//
// goroutine 是轻量线程，创建一个 goroutine 所需的资源开销很小，
// 所以可以创建非常多的 goroutine 来并发工作。
//
// 它们是由 Go 运行时调度的。调度过程就是 Go 运行时把 goroutine 任务
// 分配给 CPU 执行的过程。
//
// 但是 goroutine 不是通常理解的线程，线程是操作系统调度的。
//
// 在 Go 中，想让某个任务并发或者异步执行，只需把任务封装为一个函数或闭包，
// 交给 goroutine 执行即可。

// GoroutineDemo 演示 goroutine 基本使用
func GoroutineDemo() {
	fmt.Println("========== 1.18.1 goroutine ==========")
	fmt.Println()
	fmt.Println("goroutine 是轻量线程，创建一个 goroutine 所需的资源开销很小，")
	fmt.Println("所以可以创建非常多的 goroutine 来并发工作。")
	fmt.Println()
	fmt.Println("它们是由 Go 运行时调度的。调度过程就是 Go 运行时把 goroutine 任务")
	fmt.Println("分配给 CPU 执行的过程。")
	fmt.Println()
	fmt.Println("但是 goroutine 不是通常理解的线程，线程是操作系统调度的。")
	fmt.Println()
	fmt.Println("在 Go 中，想让某个任务并发或者异步执行，只需把任务封装为一个函数或闭包，")
	fmt.Println("交给 goroutine 执行即可。")
	fmt.Println()

	demonstrateGoroutineBasic()
	demonstrateGoroutineClosure()
	demonstrateThreadSafety()
	demonstrateThreadSafeCounter()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ goroutine 是轻量级线程，资源开销小")
	fmt.Println("✅ 由 Go 运行时调度，不是操作系统线程")
	fmt.Println("✅ 使用 go 关键字启动 goroutine")
	fmt.Println("✅ 可以传递函数或闭包给 goroutine")
	fmt.Println("✅ Go 并发存在线程安全问题，需要使用 sync.Mutex 等机制保护")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - goroutine 是异步执行的，主程序不会等待 goroutine 完成")
	fmt.Println("   - 需要使用 sync.WaitGroup 或 channel 来等待 goroutine 完成")
	fmt.Println("   - 共享数据需要加锁保护，避免竞态条件")
	fmt.Println("   - Go 标准库大多数数据结构默认非线程安全")
	fmt.Println()
}

// say 打印字符串的函数
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// demonstrateGoroutineBasic 演示 goroutine 基本使用
func demonstrateGoroutineBasic() {
	fmt.Println("=== 1.18.1.1 goroutine 基本使用 ===")
	fmt.Println()

	fmt.Println("1. 声明方式 1：把方法或函数交给 goroutine 执行：")
	fmt.Println("   go <method_name>(<method_params>...)")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   go say(\"in goroutine: world\")")
	fmt.Println("   say(\"hello\")")
	fmt.Println()

	fmt.Println("2. 声明方式 2：把闭包交给 goroutine 执行：")
	fmt.Println("   go func(<method_params>...){")
	fmt.Println("       <statement_or_expression>")
	fmt.Println("       ...")
	fmt.Println("   }(<params>...)")
	fmt.Println()
	fmt.Println("   示例：")
	fmt.Println("   go func() {")
	fmt.Println("       fmt.Println(\"run goroutine in closure\")")
	fmt.Println("   }()")
	fmt.Println()
	fmt.Println("   go func(s string) {")
	fmt.Println("       fmt.Println(s)")
	fmt.Println("   }(\"goroutine: closure params\")")
	fmt.Println()

	fmt.Println("3. 实际运行示例：")
	fmt.Println()

	// 启动 goroutine
	go func() {
		fmt.Println("   [goroutine] run goroutine in closure")
	}()

	go func(s string) {
		fmt.Println("   [goroutine]", s)
	}("goroutine: closure params")

	go say("   [goroutine] in goroutine: world")

	// 主程序继续执行
	say("   [main] hello")

	// 等待 goroutine 完成（简单示例，实际应该用 WaitGroup）
	time.Sleep(600 * time.Millisecond)
	fmt.Println()
}

// demonstrateGoroutineClosure 演示 goroutine 中使用闭包
func demonstrateGoroutineClosure() {
	fmt.Println("=== 1.18.1.2 goroutine 中使用闭包 ===")
	fmt.Println()

	fmt.Println("1. 闭包捕获变量：")
	fmt.Println("   - goroutine 中的闭包可以捕获外部变量")
	fmt.Println("   - 多个 goroutine 共享同一个变量时需要注意线程安全")
	fmt.Println()

	fmt.Println("2. 示例代码：")
	fmt.Println("   for i := 0; i < 3; i++ {")
	fmt.Println("       go func() {")
	fmt.Println("           fmt.Println(i)  // 可能打印 0, 1, 2 或 3, 3, 3")
	fmt.Println("       }()")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   问题：闭包捕获的是变量 i 的引用，不是值")
	fmt.Println("   解决：通过参数传递值")
	fmt.Println()

	fmt.Println("3. 正确的做法：")
	fmt.Println("   for i := 0; i < 3; i++ {")
	fmt.Println("       go func(n int) {")
	fmt.Println("           fmt.Println(n)  // 正确打印 0, 1, 2")
	fmt.Println("       }(i)  // 传递值")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("4. 实际运行示例：")
	fmt.Println("   错误示例（可能有问题）：")
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("   [错误] goroutine %d: i = %d\n", i, i)
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	fmt.Println("   正确示例：")
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Printf("   [正确] goroutine %d: n = %d\n", n, n)
		}(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

// UnsafeCounter 非线程安全的计数器
type UnsafeCounter struct {
	count int
}

// Increment 增加计数（非线程安全）
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// GetCount 获取当前计数（非线程安全）
func (c *UnsafeCounter) GetCount() int {
	return c.count
}

// SafeCounter 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Increment 增加计数（线程安全）
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// GetCount 获取当前计数（线程安全）
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// demonstrateThreadSafety 演示线程安全问题
func demonstrateThreadSafety() {
	fmt.Println("=== 1.18.1.3 线程安全问题 ===")
	fmt.Println()

	fmt.Println("1. 问题说明：")
	fmt.Println("   Go 中并发同样存在线程安全问题，因为 Go 也是使用共享内存")
	fmt.Println("   让多个 goroutine 之间通信。并且大部分时候为了性能，")
	fmt.Println("   所以 Go 的大多数标准库的数据结构默认是非线程安全的。")
	fmt.Println()

	fmt.Println("2. 非线程安全的计数器示例：")
	fmt.Println("   type UnsafeCounter struct {")
	fmt.Println("       count int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func (c *UnsafeCounter) Increment() {")
	fmt.Println("       c.count += 1  // 非原子操作，存在竞态条件")
	fmt.Println("   }")
	fmt.Println()

	counter := UnsafeCounter{}

	fmt.Println("3. 启动 1000 个 goroutine 同时增加计数：")
	fmt.Println("   for i := 0; i < 1000; i++ {")
	fmt.Println("       go func() {")
	fmt.Println("           for j := 0; j < 100; j++ {")
	fmt.Println("               counter.Increment()")
	fmt.Println("           }")
	fmt.Println("       }()")
	fmt.Println("   }")
	fmt.Println()

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 启动1000个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 输出最终计数
	fmt.Printf("4. 最终计数（非线程安全）: %d\n", counter.GetCount())
	fmt.Println("   预期结果: 100000 (1000 * 100)")
	fmt.Println("   实际结果: 小于 100000（因为存在竞态条件）")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 多个 goroutine 同时修改 count 变量")
	fmt.Println("   - c.count += 1 不是原子操作，包含读取、计算、写入三个步骤")
	fmt.Println("   - 多个 goroutine 可能同时读取到相同的值，导致丢失更新")
	fmt.Println("   - 这就是竞态条件（race condition）")
	fmt.Println()
}

// demonstrateThreadSafeCounter 演示线程安全的计数器
func demonstrateThreadSafeCounter() {
	fmt.Println("=== 1.18.1.4 线程安全的计数器 ===")
	fmt.Println()

	fmt.Println("1. 线程安全的计数器实现：")
	fmt.Println("   type SafeCounter struct {")
	fmt.Println("       mu    sync.Mutex  // 互斥锁")
	fmt.Println("       count int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func (c *SafeCounter) Increment() {")
	fmt.Println("       c.mu.Lock()         // 加锁")
	fmt.Println("       defer c.mu.Unlock() // 确保解锁")
	fmt.Println("       c.count++")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - sync.Mutex 是互斥锁，保证同一时间只有一个 goroutine 能访问")
	fmt.Println("   - Lock() 获取锁，Unlock() 释放锁")
	fmt.Println("   - defer 确保即使发生 panic 也能释放锁")
	fmt.Println()

	safeCounter := SafeCounter{}

	fmt.Println("2. 使用线程安全的计数器：")
	fmt.Println("   启动 1000 个 goroutine 同时增加计数")
	fmt.Println()

	var wg sync.WaitGroup

	// 启动1000个goroutine同时增加计数
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				safeCounter.Increment()
			}
		}()
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 输出最终计数
	fmt.Printf("3. 最终计数（线程安全）: %d\n", safeCounter.GetCount())
	fmt.Println("   预期结果: 100000 (1000 * 100)")
	fmt.Println("   实际结果: 100000（正确！）")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 使用互斥锁保护共享数据")
	fmt.Println("   - 同一时间只有一个 goroutine 能修改 count")
	fmt.Println("   - 避免了竞态条件，保证了数据一致性")
	fmt.Println()
	fmt.Println("4. sync.WaitGroup 的使用：")
	fmt.Println("   var wg sync.WaitGroup")
	fmt.Println("   wg.Add(1)      // 增加等待计数")
	fmt.Println("   go func() {")
	fmt.Println("       defer wg.Done()  // 完成时减少计数")
	fmt.Println("       // ... 执行任务")
	fmt.Println("   }()")
	fmt.Println("   wg.Wait()      // 等待所有 goroutine 完成")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - WaitGroup 用于等待多个 goroutine 完成")
	fmt.Println("   - Add(n) 增加等待计数")
	fmt.Println("   - Done() 减少等待计数（相当于 Add(-1)）")
	fmt.Println("   - Wait() 阻塞直到计数为 0")
	fmt.Println()
}

