package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// ========== 1.18.3 锁与 channel ==========
//
// 在 Go 中，当需要 goroutine 之间协作的地方，更常见的方式是使用 channel，
// 而不是 sync 包中的 Mutex 或 RWMutex 的互斥锁。但其实它们各有侧重。
//
// 大部分时候，流程是根据数据驱动的，channel 会被使用得更频繁。

// LockAndChannelDemo 演示锁与 channel 的使用场景
func LockAndChannelDemo() {
	fmt.Println("========== 1.18.3 锁与 channel ==========")
	fmt.Println()
	fmt.Println("在 Go 中，当需要 goroutine 之间协作的地方，更常见的方式是使用 channel，")
	fmt.Println("而不是 sync 包中的 Mutex 或 RWMutex 的互斥锁。但其实它们各有侧重。")
	fmt.Println()
	fmt.Println("大部分时候，流程是根据数据驱动的，channel 会被使用得更频繁。")
	fmt.Println()

	demonstrateChannelScenarios()
	demonstrateLockScenarios()
	demonstrateComparison()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ channel 擅长数据流动的场景：传递数据所有权、分发任务、交流异步结果")
	fmt.Println("✅ 锁适合的场景：访问缓存、管理状态、保护共享资源")
	fmt.Println("✅ Go 的哲学：'通过通信共享内存，而不是通过共享内存通信'")
	fmt.Println("✅ 优先使用 channel，当 channel 不合适时再考虑锁")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - channel 用于 goroutine 之间的通信和协调")
	fmt.Println("   - 锁用于保护共享资源的并发访问")
	fmt.Println("   - 两者可以结合使用，但要注意避免死锁")
	fmt.Println()
}

// demonstrateChannelScenarios 演示 channel 的使用场景
func demonstrateChannelScenarios() {
	fmt.Println("=== 1.18.3.1 channel 的使用场景 ===")
	fmt.Println()
	fmt.Println("channel 擅长的是数据流动的场景：")
	fmt.Println()

	fmt.Println("1. 传递数据的所有权")
	fmt.Println("   即把某个数据发送给其他协程。")
	fmt.Println()
	demonstrateDataOwnership()

	fmt.Println("2. 分发任务")
	fmt.Println("   每个任务都是一个数据。")
	fmt.Println()
	demonstrateTaskDistribution()

	fmt.Println("3. 交流异步结果")
	fmt.Println("   结果是一个数据。")
	fmt.Println()
	demonstrateAsyncResult()
}

// demonstrateDataOwnership 演示传递数据所有权
func demonstrateDataOwnership() {
	fmt.Println("   场景1：传递数据的所有权")
	fmt.Println()

	ch := make(chan string, 3)

	// 生产者：创建数据并发送
	go func() {
		data := []string{"数据1", "数据2", "数据3"}
		for _, d := range data {
			ch <- d
			fmt.Printf("   [生产者] 发送数据所有权: %s\n", d)
		}
		close(ch)
	}()

	// 消费者：接收数据并处理
	fmt.Println("   [消费者] 接收数据：")
	for data := range ch {
		fmt.Printf("   [消费者] 接收到数据: %s (现在拥有这个数据)\n", data)
		// 处理数据...
	}
	fmt.Println()
}

// demonstrateTaskDistribution 演示分发任务
func demonstrateTaskDistribution() {
	fmt.Println("   场景2：分发任务")
	fmt.Println()

	taskCh := make(chan int, 10)
	var wg sync.WaitGroup

	// 创建任务
	tasks := []int{1, 2, 3, 4, 5}
	go func() {
		for _, task := range tasks {
			taskCh <- task
			fmt.Printf("   [任务分发] 分发任务: %d\n", task)
		}
		close(taskCh)
	}()

	// 多个工作 goroutine 处理任务
	workerCount := 3
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range taskCh {
				fmt.Printf("   [工作协程 %d] 处理任务: %d\n", workerID, task)
				time.Sleep(50 * time.Millisecond) // 模拟处理时间
			}
		}(i)
	}

	wg.Wait()
	fmt.Println()
}

// demonstrateAsyncResult 演示交流异步结果
func demonstrateAsyncResult() {
	fmt.Println("   场景3：交流异步结果")
	fmt.Println()

	resultCh := make(chan int, 3)

	// 启动多个异步任务
	for i := 1; i <= 3; i++ {
		go func(id int) {
			// 模拟异步操作
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			result := id * 10
			resultCh <- result
			fmt.Printf("   [任务 %d] 完成，结果: %d\n", id, result)
		}(i)
	}

	// 收集结果
	fmt.Println("   [主程序] 收集异步结果：")
	for i := 0; i < 3; i++ {
		result := <-resultCh
		fmt.Printf("   [主程序] 收到结果: %d\n", result)
	}
	fmt.Println()
}

// Cache 简单的缓存结构（使用锁保护）
type Cache struct { // type 关键字定义类型，Cache 是类型名，struct 关键字定义结构体
	mu    sync.RWMutex           // mu 是结构体字段名，sync.RWMutex 是字段类型（读写互斥锁）
	items map[string]interface{} // items 是字段名，map[string]interface{} 是字段类型（键为字符串，值为任意类型）
} // 结构体定义结束

// NewCache 创建新的缓存
func NewCache() *Cache { // func 关键字定义函数，NewCache 是函数名，返回 *Cache 指针类型
	return &Cache{ // return 语句返回 &Cache{} 结构体字面量，& 是取地址运算符
		items: make(map[string]interface{}), // items 是结构体字段名，make() 创建 map，key 类型 string，value 类型 interface{}
	} // 结构体字面量结束大括号
} // 函数体结束大括号

// Get 获取缓存值（读锁）
func (c *Cache) Get(key string) (interface{}, bool) { // 方法接收者 c *Cache，参数 key string，返回 (interface{}, bool)
	c.mu.RLock()                  // 获取读锁，RLock() 允许多个 goroutine 同时读取
	defer c.mu.RUnlock()          // defer 延迟执行解锁，确保函数结束时释放锁
	value, exists := c.items[key] // 从 map 中获取值，map 读取操作需要锁保护
	return value, exists          // 返回两个值：缓存值和是否存在标志
} // 函数结束

// Set 设置缓存值（写锁）
func (c *Cache) Set(key string, value interface{}) { // 方法接收者 c *Cache，参数 key string, value interface{}
	c.mu.Lock()          // 获取写锁，Lock() 阻塞直到获取独占锁
	defer c.mu.Unlock()  // defer 延迟执行解锁，确保函数结束时释放锁
	c.items[key] = value // 设置 map 中的值，map 写入操作需要锁保护
} // 函数结束

// Counter 计数器（使用锁保护状态）
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment 增加计数
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Get 获取当前计数
func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// demonstrateLockScenarios 演示锁的使用场景
func demonstrateLockScenarios() {
	fmt.Println("=== 1.18.3.2 锁的使用场景 ===")
	fmt.Println()
	fmt.Println("锁使用的场景更偏向同一时间只给一个协程访问数据的权限：")
	fmt.Println()

	fmt.Println("1. 访问缓存")
	fmt.Println()
	demonstrateCacheAccess()

	fmt.Println("2. 管理状态")
	fmt.Println()
	demonstrateStateManagement()
}

// demonstrateCacheAccess 演示访问缓存
func demonstrateCacheAccess() {
	fmt.Println("   场景1：访问缓存")
	fmt.Println()

	cache := NewCache()

	var wg sync.WaitGroup

	// 多个 goroutine 同时读写缓存
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 写入缓存
			key := fmt.Sprintf("key%d", id)
			value := fmt.Sprintf("value%d", id)
			cache.Set(key, value)
			fmt.Printf("   [协程 %d] 写入缓存: %s = %s\n", id, key, value)

			// 读取缓存
			if v, ok := cache.Get(key); ok {
				fmt.Printf("   [协程 %d] 读取缓存: %s = %s\n", id, key, v)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 使用 RWMutex 保护缓存，支持并发读取")
	fmt.Println("   - 写入时使用写锁，保证数据一致性")
	fmt.Println("   - 这是典型的锁使用场景：保护共享资源")
	fmt.Println()
}

// demonstrateStateManagement 演示管理状态
func demonstrateStateManagement() {
	fmt.Println("   场景2：管理状态")
	fmt.Println()

	counter := &Counter{}

	var wg sync.WaitGroup

	// 多个 goroutine 同时修改状态
	fmt.Println("   多个 goroutine 同时增加计数：")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	fmt.Printf("   最终计数: %d (预期: 10)\n", counter.Get())
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 使用 Mutex 保护共享状态（计数器）")
	fmt.Println("   - 确保同一时间只有一个 goroutine 能修改状态")
	fmt.Println("   - 这是典型的锁使用场景：管理共享状态")
	fmt.Println()
}

// demonstrateComparison 演示 channel 和锁的对比
func demonstrateComparison() {
	fmt.Println("=== 1.18.3.3 channel 与锁的对比 ===")
	fmt.Println()

	fmt.Println("1. 设计理念对比：")
	fmt.Println()
	fmt.Println("   Channel（通信）：")
	fmt.Println("   - Go 的哲学：'通过通信共享内存，而不是通过共享内存通信'")
	fmt.Println("   - 关注数据流动和 goroutine 之间的协调")
	fmt.Println("   - 更适合数据驱动的场景")
	fmt.Println()
	fmt.Println("   锁（共享内存）：")
	fmt.Println("   - 传统的并发控制方式")
	fmt.Println("   - 关注保护共享资源的访问")
	fmt.Println("   - 更适合保护共享状态和缓存")
	fmt.Println()

	fmt.Println("2. 使用场景对比：")
	fmt.Println()
	fmt.Println("   Channel 适合：")
	fmt.Println("   ✅ 传递数据所有权")
	fmt.Println("   ✅ 分发任务")
	fmt.Println("   ✅ 交流异步结果")
	fmt.Println("   ✅ goroutine 之间的协调和同步")
	fmt.Println("   ✅ 流水线处理")
	fmt.Println()
	fmt.Println("   锁适合：")
	fmt.Println("   ✅ 保护共享资源（如缓存、数据库连接池）")
	fmt.Println("   ✅ 管理共享状态（如计数器、配置）")
	fmt.Println("   ✅ 需要细粒度控制的场景")
	fmt.Println("   ✅ 性能要求极高的场景（锁的开销可能更小）")
	fmt.Println()

	fmt.Println("3. 选择建议：")
	fmt.Println()
	fmt.Println("   📌 优先使用 channel：")
	fmt.Println("      - 大部分情况下，channel 是更好的选择")
	fmt.Println("      - 代码更清晰，更符合 Go 的哲学")
	fmt.Println("      - 更容易理解和维护")
	fmt.Println()
	fmt.Println("   📌 使用锁的情况：")
	fmt.Println("      - channel 不适合的场景（如保护缓存）")
	fmt.Println("      - 需要保护共享资源时")
	fmt.Println("      - 性能要求极高，锁的开销更小时")
	fmt.Println()
	fmt.Println("   📌 可以结合使用：")
	fmt.Println("      - channel 用于 goroutine 之间的通信")
	fmt.Println("      - 锁用于保护共享资源")
	fmt.Println("      - 两者可以很好地配合使用")
	fmt.Println()

	fmt.Println("4. 实际示例对比：")
	fmt.Println()
	fmt.Println("   使用 Channel 实现计数器：")
	demonstrateCounterWithChannel()

	fmt.Println("   使用锁实现计数器：")
	demonstrateCounterWithLock()
}

// demonstrateCounterWithChannel 使用 channel 实现计数器
func demonstrateCounterWithChannel() {
	fmt.Println()
	fmt.Println("   type ChannelCounter struct {")
	fmt.Println("       ch chan int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 通过 channel 发送增量请求")
	fmt.Println("   - 单个 goroutine 处理所有增量请求")
	fmt.Println("   - 保证线程安全，但可能性能不如锁")
	fmt.Println()

	// 简化的 channel 计数器实现
	type ChannelCounter struct {
		ch       chan int
		resultCh chan int
		doneCh   chan struct{}
	}

	counter := &ChannelCounter{
		ch:       make(chan int),
		resultCh: make(chan int),
		doneCh:   make(chan struct{}),
	}

	count := 0

	// 启动处理 goroutine
	go func() {
		for {
			select {
			case <-counter.ch:
				count++
			case <-counter.doneCh:
				counter.resultCh <- count
				return
			}
		}
	}()

	// 模拟增加计数
	for i := 0; i < 5; i++ {
		counter.ch <- 1
	}

	// 停止处理 goroutine 并获取结果
	close(counter.doneCh)
	result := <-counter.resultCh
	fmt.Printf("   Channel 计数器结果: %d\n", result)
	fmt.Println()
}

// demonstrateCounterWithLock 使用锁实现计数器
func demonstrateCounterWithLock() {
	fmt.Println()
	fmt.Println("   type LockCounter struct {")
	fmt.Println("       mu    sync.Mutex")
	fmt.Println("       count int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 使用 Mutex 保护共享状态")
	fmt.Println("   - 直接修改共享变量")
	fmt.Println("   - 性能通常更好，代码更简单")
	fmt.Println()

	counter := &Counter{}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("   锁计数器结果: %d\n", counter.Get())
	fmt.Println()
}
