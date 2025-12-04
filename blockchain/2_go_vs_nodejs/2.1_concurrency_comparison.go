package comparison

import (
	"fmt"
	"sync"
	"time"
)

// ========== 2.1 Go vs Node.js 并发模型对比 ==========
//
// 本文件对比 Go 的 goroutine/channel 和 Node.js 的并发模型
// 帮助理解两种语言的并发编程差异

// GoVsNodejsConcurrencyDemo 演示 Go 和 Node.js 并发模型的对比
func GoVsNodejsConcurrencyDemo() {
	fmt.Println("========== 2.1 Go vs Node.js 并发模型对比 ==========")
	fmt.Println()

	compareGoroutineVsEventLoop()
	compareChannelVsCallback()
	compareConcurrencyModels()
	comparePerformance()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ Go 的 goroutine 是真正的并发，Node.js 的事件循环是异步单线程")
	fmt.Println("✅ Go 的 channel 是类型安全的通信机制，Node.js 使用回调/Promise/EventEmitter")
	fmt.Println("✅ Go 适合 CPU 密集型任务，Node.js 适合 I/O 密集型任务")
	fmt.Println("✅ Go 可以充分利用多核 CPU，Node.js 需要 Worker Threads 才能利用多核")
	fmt.Println()
}

// ========== 1. Goroutine vs 事件循环 ==========

// compareGoroutineVsEventLoop 对比 goroutine 和事件循环
func compareGoroutineVsEventLoop() {
	fmt.Println("=== 2.1.1 Goroutine vs Node.js 事件循环 ===")
	fmt.Println()

	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│ Go 的 Goroutine                                          │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│ ✅ 真正的并发：多个 goroutine 可以同时运行在不同 CPU 核心上 │")
	fmt.Println("│ ✅ 轻量级：每个 goroutine 约 2KB 内存                   │")
	fmt.Println("│ ✅ 可以创建数百万个 goroutine                           │")
	fmt.Println("│ ✅ 由 Go 运行时调度，充分利用多核 CPU                   │")
	fmt.Println("│ ✅ 适合 CPU 密集型任务                                  │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│ Node.js 的事件循环                                      │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│ ⚠️  单线程：所有代码在单个线程中执行                    │")
	fmt.Println("│ ✅ 异步 I/O：I/O 操作不会阻塞事件循环                   │")
	fmt.Println("│ ⚠️  CPU 密集型任务会阻塞整个事件循环                    │")
	fmt.Println("│ ✅ 适合 I/O 密集型任务（网络、文件、数据库）            │")
	fmt.Println("│ ⚠️  需要 Worker Threads 才能利用多核 CPU                │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()

	// Go 示例：真正的并发
	fmt.Println("Go 示例：真正的并发执行")
	demonstrateGoroutineConcurrency()

	// Node.js 模拟：单线程事件循环
	fmt.Println("Node.js 模拟：单线程事件循环（在 Go 中模拟）")
	demonstrateEventLoopSimulation()

	fmt.Println()
}

// demonstrateGoroutineConcurrency 演示 Go 的并发
func demonstrateGoroutineConcurrency() {
	fmt.Println("   Go: 多个 goroutine 可以真正并行执行")
	fmt.Println()

	var wg sync.WaitGroup
	start := time.Now()

	// 启动 5 个 goroutine，每个执行 CPU 密集型任务
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟 CPU 密集型任务
			sum := 0
			for j := 0; j < 1000000; j++ {
				sum += j
			}
			fmt.Printf("   [Goroutine %d] 完成计算，结果: %d\n", id, sum)
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("   ✅ 5 个 goroutine 并行执行完成，耗时: %v\n", elapsed)
	fmt.Println()
}

// demonstrateEventLoopSimulation 模拟 Node.js 的事件循环
func demonstrateEventLoopSimulation() {
	fmt.Println("   Node.js: 单线程事件循环，任务按顺序执行")
	fmt.Println()

	start := time.Now()

	// 模拟 Node.js 的事件循环：任务按顺序执行
	for i := 1; i <= 5; i++ {
		// 模拟 CPU 密集型任务（会阻塞）
		sum := 0
		for j := 0; j < 1000000; j++ {
			sum += j
		}
		fmt.Printf("   [事件循环] 任务 %d 完成，结果: %d\n", i, sum)
	}

	elapsed := time.Since(start)
	fmt.Printf("   ⚠️  5 个任务串行执行完成，耗时: %v（比 Go 慢）\n", elapsed)
	fmt.Println()
}

// ========== 2. Channel vs 回调/Promise/EventEmitter ==========

// compareChannelVsCallback 对比 channel 和 Node.js 的通信机制
func compareChannelVsCallback() {
	fmt.Println("=== 2.1.2 Channel vs Node.js 通信机制 ===")
	fmt.Println()

	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│ Go 的 Channel                                            │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│ ✅ 类型安全：编译时检查类型                              │")
	fmt.Println("│ ✅ 线程安全：内置的并发安全机制                          │")
	fmt.Println("│ ✅ 阻塞式：发送/接收会阻塞，直到数据准备好              │")
	fmt.Println("│ ✅ 同步语义：无缓冲 channel 保证同步                     │")
	fmt.Println("│ ✅ select 多路复用：同时监听多个 channel                │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()

	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│ Node.js 的通信机制                                      │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│ 1. 回调函数（Callback）                                 │")
	fmt.Println("│    - 容易产生回调地狱                                   │")
	fmt.Println("│    - 错误处理复杂                                       │")
	fmt.Println("│                                                         │")
	fmt.Println("│ 2. Promise/async-await                                  │")
	fmt.Println("│    - 解决回调地狱                                       │")
	fmt.Println("│    - 但仍然是单线程，不能真正并行                      │")
	fmt.Println("│                                                         │")
	fmt.Println("│ 3. EventEmitter                                         │")
	fmt.Println("│    - 发布-订阅模式                                      │")
	fmt.Println("│    - 类似 channel，但非阻塞                             │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()

	// Go Channel 示例
	fmt.Println("Go 示例：使用 Channel 通信")
	demonstrateChannelCommunication()

	// Node.js 模拟：Promise/EventEmitter
	fmt.Println("Node.js 模拟：Promise/EventEmitter（在 Go 中模拟）")
	demonstrateNodejsCommunication()

	fmt.Println()
}

// demonstrateChannelCommunication 演示 Go 的 channel 通信
func demonstrateChannelCommunication() {
	fmt.Println("   Go: 使用 channel 进行 goroutine 间通信")
	fmt.Println()

	// 创建 channel
	ch := make(chan string, 3)

	// 生产者 goroutine
	go func() {
		messages := []string{"消息1", "消息2", "消息3"}
		for _, msg := range messages {
			ch <- msg
			fmt.Printf("   [生产者] 发送: %s\n", msg)
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	// 消费者：接收消息
	fmt.Println("   [消费者] 接收消息：")
	for msg := range ch {
		fmt.Printf("   [消费者] 接收: %s\n", msg)
	}
	fmt.Println()
}

// demonstrateNodejsCommunication 模拟 Node.js 的通信
func demonstrateNodejsCommunication() {
	fmt.Println("   Node.js: 使用 Promise/EventEmitter（模拟）")
	fmt.Println()

	// 模拟 Promise
	fmt.Println("   1. Promise 示例（模拟）：")
	fmt.Println("      fetchData()")
	fmt.Println("        .then(data => processData(data))")
	fmt.Println("        .then(result => console.log(result))")
	fmt.Println("        .catch(error => console.error(error))")
	fmt.Println()

	// 模拟 EventEmitter
	fmt.Println("   2. EventEmitter 示例（模拟）：")
	fmt.Println("      emitter.on('data', (data) => {")
	fmt.Println("          console.log('收到数据:', data)")
	fmt.Println("      })")
	fmt.Println("      emitter.emit('data', '消息1')")
	fmt.Println()

	fmt.Println("   ⚠️  注意：Node.js 的通信是异步的，但不是并行的")
	fmt.Println()
}

// ========== 3. 并发模型对比 ==========

// compareConcurrencyModels 对比并发模型
func compareConcurrencyModels() {
	fmt.Println("=== 2.1.3 并发模型详细对比 ===")
	fmt.Println()

	fmt.Println("┌──────────────┬──────────────────┬──────────────────────┐")
	fmt.Println("│ 特性          │ Go                │ Node.js              │")
	fmt.Println("├──────────────┼──────────────────┼──────────────────────┤")
	fmt.Println("│ 并发模型      │ goroutine (协程)  │ 事件循环 (单线程)     │")
	fmt.Println("│ 内存占用      │ ~2KB/goroutine    │ ~1MB/Worker Thread   │")
	fmt.Println("│ 最大并发数    │ 数百万            │ 数千 (受内存限制)     │")
	fmt.Println("│ CPU 密集型    │ ✅ 优秀           │ ❌ 会阻塞事件循环     │")
	fmt.Println("│ I/O 密集型    │ ✅ 优秀           │ ✅ 优秀               │")
	fmt.Println("│ 多核利用      │ ✅ 自动           │ ⚠️  需要 Worker       │")
	fmt.Println("│ 通信机制      │ channel           │ Promise/EventEmitter│")
	fmt.Println("│ 类型安全      │ ✅ 编译时检查     │ ⚠️  运行时检查         │")
	fmt.Println("│ 错误处理      │ 多返回值          │ try-catch/Promise    │")
	fmt.Println("└──────────────┴──────────────────┴──────────────────────┘")
	fmt.Println()

	// 实际场景对比
	fmt.Println("实际场景对比：")
	fmt.Println()

	// 场景 1: 并发处理多个 HTTP 请求
	fmt.Println("场景 1: 并发处理 1000 个 HTTP 请求")
	compareHTTPRequests()

	// 场景 2: CPU 密集型任务
	fmt.Println("场景 2: 处理 CPU 密集型任务")
	compareCPUTasks()

	fmt.Println()
}

// compareHTTPRequests 对比处理 HTTP 请求
func compareHTTPRequests() {
	fmt.Println("   Go: 使用 goroutine 并发处理")
	fmt.Println("   ```go")
	fmt.Println("   for _, url := range urls {")
	fmt.Println("       go fetch(url)  // 每个请求一个 goroutine")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println("   ✅ 可以轻松处理数万个并发请求")
	fmt.Println("   ✅ 内存占用低（每个 goroutine ~2KB）")
	fmt.Println()

	fmt.Println("   Node.js: 使用事件循环处理")
	fmt.Println("   ```javascript")
	fmt.Println("   for (const url of urls) {")
	fmt.Println("       fetch(url).then(...)  // 异步，但不并行")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println("   ✅ I/O 操作不阻塞，性能好")
	fmt.Println("   ⚠️  但受限于单线程，CPU 任务会阻塞")
	fmt.Println()
}

// compareCPUTasks 对比 CPU 密集型任务
func compareCPUTasks() {
	fmt.Println("   Go: CPU 密集型任务")
	fmt.Println("   ```go")
	fmt.Println("   for i := 0; i < 10; i++ {")
	fmt.Println("       go heavyComputation()  // 真正并行执行")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println("   ✅ 充分利用多核 CPU")
	fmt.Println("   ✅ 10 个任务可以同时在 10 个 CPU 核心上运行")
	fmt.Println()

	fmt.Println("   Node.js: CPU 密集型任务")
	fmt.Println("   ```javascript")
	fmt.Println("   for (let i = 0; i < 10; i++) {")
	fmt.Println("       heavyComputation()  // 会阻塞事件循环")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println("   ❌ 会阻塞事件循环，其他请求无法处理")
	fmt.Println("   ⚠️  需要使用 Worker Threads:")
	fmt.Println("   ```javascript")
	fmt.Println("   const { Worker } = require('worker_threads')")
	fmt.Println("   const worker = new Worker('./worker.js')")
	fmt.Println("   ```")
	fmt.Println("   ⚠️  Worker Threads 开销大（~1MB 内存）")
	fmt.Println()
}

// ========== 4. 性能对比 ==========

// comparePerformance 对比性能
func comparePerformance() {
	fmt.Println("=== 2.1.4 性能对比 ===")
	fmt.Println()

	fmt.Println("测试场景：处理 10,000 个并发任务")
	fmt.Println()

	// Go 性能测试
	fmt.Println("Go 性能测试：")
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟轻量级任务
			_ = id * 2
		}(i)
	}
	wg.Wait()
	goElapsed := time.Since(start)
	fmt.Printf("   ✅ Go: 10,000 个 goroutine 完成，耗时: %v\n", goElapsed)
	fmt.Printf("   ✅ 内存占用: ~20MB (10,000 × 2KB)\n")
	fmt.Println()

	// Node.js 模拟（在 Go 中模拟单线程）
	fmt.Println("Node.js 模拟（单线程事件循环）：")
	start = time.Now()
	for i := 0; i < 10000; i++ {
		// 模拟事件循环：任务按顺序执行
		_ = i * 2
	}
	nodejsElapsed := time.Since(start)
	fmt.Printf("   ⚠️  Node.js: 10,000 个任务串行执行，耗时: %v\n", nodejsElapsed)
	fmt.Printf("   ⚠️  实际 Node.js 会更快（因为异步 I/O），但 CPU 任务会阻塞\n")
	fmt.Println()

	fmt.Printf("   性能对比: Go 并发执行 vs Node.js 串行执行\n")
	fmt.Printf("   Go 优势: %.2fx 更快（在 CPU 密集型任务中）\n",
		float64(nodejsElapsed)/float64(goElapsed))
	fmt.Println()
}

// ========== 5. 实际应用场景 ==========

// demonstrateRealWorldScenarios 演示实际应用场景
func demonstrateRealWorldScenarios() {
	fmt.Println("=== 2.1.5 实际应用场景选择 ===")
	fmt.Println()

	fmt.Println("选择 Go 的场景：")
	fmt.Println("  ✅ 区块链节点（需要处理大量并发连接）")
	fmt.Println("  ✅ 微服务 API（需要高性能）")
	fmt.Println("  ✅ 实时数据处理（需要低延迟）")
	fmt.Println("  ✅ CPU 密集型任务（图像处理、加密计算）")
	fmt.Println("  ✅ 需要充分利用多核 CPU 的应用")
	fmt.Println()

	fmt.Println("选择 Node.js 的场景：")
	fmt.Println("  ✅ Web 前端开发（React、Vue 等）")
	fmt.Println("  ✅ 快速原型开发")
	fmt.Println("  ✅ I/O 密集型应用（API 网关、代理服务器）")
	fmt.Println("  ✅ 实时应用（WebSocket 服务器）")
	fmt.Println("  ✅ 需要丰富 npm 生态的项目")
	fmt.Println()

	fmt.Println("混合架构（推荐）：")
	fmt.Println("  - 前端: Node.js/React")
	fmt.Println("  - API 网关: Go（高性能）")
	fmt.Println("  - 核心业务: Go（高并发）")
	fmt.Println("  - 工具脚本: Node.js（快速开发）")
	fmt.Println()
}

// ========== 6. Node.js 的替代方案 ==========

// demonstrateNodejsAlternatives 演示 Node.js 的替代方案
func demonstrateNodejsAlternatives() {
	fmt.Println("=== 2.1.6 Node.js 中类似 goroutine/channel 的方案 ===")
	fmt.Println()

	fmt.Println("1. Worker Threads（类似 goroutine，但开销大）")
	fmt.Println("   ```javascript")
	fmt.Println("   const { Worker, isMainThread } = require('worker_threads')")
	fmt.Println("   if (isMainThread) {")
	fmt.Println("       const worker = new Worker(__filename)")
	fmt.Println("   } else {")
	fmt.Println("       // Worker 线程中的代码")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println("   ⚠️  每个 Worker 约 1MB 内存（vs goroutine 的 2KB）")
	fmt.Println("   ⚠️  创建开销大，不适合大量并发")
	fmt.Println()

	fmt.Println("2. EventEmitter（类似 channel，但非阻塞）")
	fmt.Println("   ```javascript")
	fmt.Println("   const EventEmitter = require('events')")
	fmt.Println("   const emitter = new EventEmitter()")
	fmt.Println("   emitter.on('data', (data) => { ... })")
	fmt.Println("   emitter.emit('data', 'message')")
	fmt.Println("   ```")
	fmt.Println("   ⚠️  非阻塞，不保证同步")
	fmt.Println("   ⚠️  没有类型安全")
	fmt.Println()

	fmt.Println("3. Promise/async-await（类似异步 channel）")
	fmt.Println("   ```javascript")
	fmt.Println("   async function fetchData() {")
	fmt.Println("       const data = await fetch(url)")
	fmt.Println("       return data")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println("   ✅ 解决回调地狱")
	fmt.Println("   ⚠️  仍然是单线程，不能真正并行")
	fmt.Println()

	fmt.Println("4. 第三方库：类似 channel 的实现")
	fmt.Println("   - `async-channel`: 提供类似 Go channel 的 API")
	fmt.Println("   - `p-queue`: 提供队列和并发控制")
	fmt.Println("   ⚠️  但性能不如 Go 的原生 channel")
	fmt.Println()
}
