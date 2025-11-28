package mapcollection

import (
	"fmt"
	"sync"
	"time"
)

// ========== 1.14.4 并发时使用 map 集合 ==========
//
// map 集合在并发场景下的使用
// 重要：Go 的 map 不是并发安全的
// 当多个 goroutine 同时读写 map 时，会导致 panic
//
// 解决方案：
// 1. 使用互斥锁（sync.Mutex）保护 map
// 2. 使用 sync.Map（适用于读多写少的场景）

// demonstrateConcurrentMapError 演示并发读写 map 会报错
func demonstrateConcurrentMapError() {
	fmt.Println("=== 1. 并发读写 map 会报错 ===")
	fmt.Println()
	fmt.Println("⚠️ 警告：以下代码会触发 panic，仅用于演示错误情况")
	fmt.Println("实际运行时会报错：")
	fmt.Println("  fatal error: concurrent map writes")
	fmt.Println("  或")
	fmt.Println("  fatal error: concurrent map read and map write")
	fmt.Println()

	fmt.Println("错误示例代码：")
	fmt.Println("  m := make(map[string]int)")
	fmt.Println("  go func() { for { m[\"a\"]++ } }()")
	fmt.Println("  go func() { for { m[\"a\"]++; fmt.Println(m[\"a\"]) } }()")
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - Go 的 map 不是并发安全的")
	fmt.Println("  - 多个 goroutine 同时读写 map 会导致 panic")
	fmt.Println("  - 必须使用同步机制保护 map")
	fmt.Println()
}

// demonstrateConcurrentMapWithMutex 演示使用互斥锁保护 map
func demonstrateConcurrentMapWithMutex() {
	fmt.Println("=== 2. 使用互斥锁保护 map（推荐方式）===")

	m := make(map[string]int)
	var wg sync.WaitGroup
	var lock sync.Mutex

	wg.Add(2)

	// goroutine 1：写入
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			lock.Lock()
			m["a"]++
			lock.Unlock()
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// goroutine 2：读写
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			lock.Lock()
			m["a"]++
			value := m["a"]
			lock.Unlock()
			fmt.Printf("  goroutine 2: m[\"a\"] = %d\n", value)
			time.Sleep(time.Millisecond * 10)
		}
	}()

	wg.Wait()
	fmt.Printf("最终 m[\"a\"] = %d\n", m["a"])
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 使用 sync.Mutex 互斥锁保护 map")
	fmt.Println("  - 读写操作前加锁，操作后解锁")
	fmt.Println("  - 保证同一时间只有一个 goroutine 访问 map")
	fmt.Println("  - 这是推荐的方式，适用于大多数场景")
	fmt.Println()
}

// demonstrateConcurrentMapWithSyncMap 演示使用 sync.Map
func demonstrateConcurrentMapWithSyncMap() {
	fmt.Println("=== 3. 使用 sync.Map（读多写少场景）===")

	var m sync.Map
	var wg sync.WaitGroup

	wg.Add(2)

	// goroutine 1：写入
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			value, _ := m.Load("a")
			var count int
			if value != nil {
				count = value.(int)
			}
			m.Store("a", count+1)
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// goroutine 2：读写
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			value, _ := m.Load("a")
			var count int
			if value != nil {
				count = value.(int)
			}
			m.Store("a", count+1)
			fmt.Printf("  goroutine 2: m[\"a\"] = %d\n", count+1)
			time.Sleep(time.Millisecond * 10)
		}
	}()

	wg.Wait()
	finalValue, _ := m.Load("a")
	fmt.Printf("最终 m[\"a\"] = %d\n", finalValue)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - sync.Map 是并发安全的，无需加锁")
	fmt.Println("  - 使用 Load() 读取，Store() 写入")
	fmt.Println("  - 适用于读多写少的场景")
	fmt.Println("  - 内存开销比普通 map 更大")
	fmt.Println("  - 类型需要断言（interface{}）")
	fmt.Println()
}

// demonstrateSyncMapMethods 演示 sync.Map 的常用方法
func demonstrateSyncMapMethods() {
	fmt.Println("=== 4. sync.Map 的常用方法 ===")

	var m sync.Map

	// Store：存储键值对
	m.Store("name", "John")
	m.Store("age", 25)
	fmt.Println("Store 后：")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true
	})

	// Load：读取值
	value, ok := m.Load("name")
	if ok {
		fmt.Printf("\nLoad(\"name\") = %v\n", value)
	}

	// LoadOrStore：读取或存储
	actual, loaded := m.LoadOrStore("city", "Beijing")
	fmt.Printf("\nLoadOrStore(\"city\", \"Beijing\"): actual=%v, loaded=%v\n", actual, loaded)

	// Delete：删除键值对
	m.Delete("age")
	fmt.Println("\nDelete(\"age\") 后：")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true
	})

	// Range：遍历
	fmt.Println("\nRange 遍历：")
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true // 返回 true 继续遍历，false 停止
	})
	fmt.Println()
}

// demonstrateMutexVsSyncMap 对比互斥锁和 sync.Map
func demonstrateMutexVsSyncMap() {
	fmt.Println("=== 5. 互斥锁 vs sync.Map 对比 ===")

	fmt.Println("┌─────────────┬──────────────┬──────────────┐")
	fmt.Println("│   特性      │  互斥锁      │  sync.Map    │")
	fmt.Println("├─────────────┼──────────────┼──────────────┤")
	fmt.Println("│ 并发安全    │ ✅ 是        │ ✅ 是        │")
	fmt.Println("│ 性能        │ 写多读少更好 │ 读多写少更好 │")
	fmt.Println("│ 类型安全    │ ✅ 是        │ ❌ 否        │")
	fmt.Println("│ 内存开销    │ 较小         │ 较大         │")
	fmt.Println("│ 使用复杂度  │ 简单         │ 稍复杂       │")
	fmt.Println("│ 适用场景    │ 大多数场景   │ 读多写少     │")
	fmt.Println("└─────────────┴──────────────┴──────────────┘")
	fmt.Println()

	fmt.Println("推荐使用场景：")
	fmt.Println("  ✅ 互斥锁（sync.Mutex）：")
	fmt.Println("     - 大多数并发场景")
	fmt.Println("     - 需要类型安全")
	fmt.Println("     - 写操作较多")
	fmt.Println()
	fmt.Println("  ✅ sync.Map：")
	fmt.Println("     - 读多写少的场景")
	fmt.Println("     - 可以接受类型断言的开销")
	fmt.Println("     - 可以接受更大的内存开销")
	fmt.Println()
}

// demonstrateConcurrentMapBestPractice 演示并发 map 的最佳实践
func demonstrateConcurrentMapBestPractice() {
	fmt.Println("=== 6. 并发 map 的最佳实践 ===")

	// 方式1：使用互斥锁（推荐）
	fmt.Println("--- 方式1：使用互斥锁（推荐）---")
	type SafeMap struct {
		mu sync.RWMutex
		m  map[string]int
	}

	safeMap := &SafeMap{
		m: make(map[string]int),
	}

	// 写入
	safeMap.mu.Lock()
	safeMap.m["key"] = 1
	safeMap.mu.Unlock()

	// 读取（使用读锁，允许多个读操作并发）
	safeMap.mu.RLock()
	value := safeMap.m["key"]
	safeMap.mu.RUnlock()

	fmt.Printf("SafeMap value = %d\n", value)
	fmt.Println()

	// 方式2：使用 sync.Map（读多写少）
	fmt.Println("--- 方式2：使用 sync.Map（读多写少）---")
	var syncMap sync.Map
	syncMap.Store("key", 1)
	val, _ := syncMap.Load("key")
	fmt.Printf("sync.Map value = %d\n", val)
	fmt.Println()

	fmt.Println("最佳实践建议：")
	fmt.Println("  1. 大多数场景使用互斥锁保护普通 map")
	fmt.Println("  2. 读多写少场景考虑使用 sync.Map")
	fmt.Println("  3. 使用 RWMutex 可以提高读性能（多个读可以并发）")
	fmt.Println("  4. 避免在锁内执行耗时操作")
	fmt.Println()
}

// MapConcurrentDemo map 并发使用完整演示
func MapConcurrentDemo() {
	fmt.Println("========== 1.14.4 并发时使用 map 集合 ==========")
	fmt.Println()
	fmt.Println("map 集合在并发场景下的使用")
	fmt.Println("重要：Go 的 map 不是并发安全的")
	fmt.Println("当多个 goroutine 同时读写 map 时，会导致 panic")
	fmt.Println()

	demonstrateConcurrentMapError()
	demonstrateConcurrentMapWithMutex()
	demonstrateConcurrentMapWithSyncMap()
	demonstrateSyncMapMethods()
	demonstrateMutexVsSyncMap()
	demonstrateConcurrentMapBestPractice()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ Go 的 map 不是并发安全的")
	fmt.Println("✅ 并发读写 map 会导致 panic")
	fmt.Println("✅ 解决方案：")
	fmt.Println("   1. 使用互斥锁（sync.Mutex）保护 map（推荐）")
	fmt.Println("   2. 使用 sync.Map（适用于读多写少的场景）")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - sync.Map 内存开销更大")
	fmt.Println("   - sync.Map 需要类型断言")
	fmt.Println("   - 推荐使用互斥锁保护普通 map")
	fmt.Println("   - 可以使用 RWMutex 提高读性能")
	fmt.Println()
}

