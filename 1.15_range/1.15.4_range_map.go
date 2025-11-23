package rangeiteration

import "fmt"

// ========== 1.15.4 对映射集合迭代 ==========
//
// 在 Go 中，使用 range 关键字迭代映射集合时：
// 1. 一种是拿到 key
// 2. 一种是拿到 key 和 value
// 3. range 关键字在迭代映射集合时，其中的 key 是乱序的（无序的）

// demonstrateMapRangeBasic 演示 map 的基本 range 迭代
func demonstrateMapRangeBasic() {
	fmt.Println("=== 1.15.4.1 map 的基本 range 迭代 ===")

	hash := map[string]int{
		"a": 1,
		"f": 2,
		"z": 3,
		"c": 4,
	}

	fmt.Printf("map hash = %v\n", hash)
	fmt.Println()

	// 方式1：只获取 key
	fmt.Println("--- 方式1：只获取 key ---")
	for key := range hash {
		fmt.Printf("key=%s, value=%d\n", key, hash[key])
	}
	fmt.Println()

	// 方式2：同时获取 key 和 value
	fmt.Println("--- 方式2：同时获取 key 和 value ---")
	for key, value := range hash {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 方式1：只获取 key，需要通过 hash[key] 访问值")
	fmt.Println("  - 方式2：同时获取 key 和 value，更高效")
	fmt.Println("  - map 的迭代顺序是随机的（无序的）")
	fmt.Println()
}

// demonstrateMapUnorderedIteration 演示 map 的无序特性
func demonstrateMapUnorderedIteration() {
	fmt.Println("=== map 的无序特性 ===")

	hash := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	fmt.Printf("map hash = %v\n", hash)
	fmt.Println()

	fmt.Println("多次遍历，顺序可能不同：")
	for i := 0; i < 3; i++ {
		fmt.Printf("第 %d 次遍历: ", i+1)
		for key := range hash {
			fmt.Printf("%s ", key)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("⚠️ 重要提示：")
	fmt.Println("  - map 的迭代顺序是随机的，每次运行可能不同")
	fmt.Println("  - 不要依赖 map 的迭代顺序")
	fmt.Println("  - 如果需要有序遍历，需要先对 key 排序")
	fmt.Println()
}

// demonstrateMapRangeWithSortedKeys 演示有序遍历 map
func demonstrateMapRangeWithSortedKeys() {
	fmt.Println("=== 有序遍历 map（需要先排序 key）===")

	hash := map[string]int{
		"z": 26,
		"a": 1,
		"f": 6,
		"c": 3,
		"b": 2,
	}

	fmt.Printf("map hash = %v\n", hash)
	fmt.Println()

	// 无序遍历
	fmt.Println("无序遍历：")
	for key, value := range hash {
		fmt.Printf("  %s: %d\n", key, value)
	}
	fmt.Println()

	// 有序遍历（需要先获取所有 key，排序后再遍历）
	fmt.Println("有序遍历（按 key 排序）：")
	// 注意：这里只是演示概念，实际需要导入 sort 包
	fmt.Println("  需要先获取所有 key，排序后再遍历")
	fmt.Println("  例如：keys := make([]string, 0, len(hash))")
	fmt.Println("       for k := range hash { keys = append(keys, k) }")
	fmt.Println("       sort.Strings(keys)")
	fmt.Println("       for _, k := range keys { fmt.Println(k, hash[k]) }")
	fmt.Println()
}

// demonstrateMapRangeIgnoreValue 演示忽略 value 的遍历
func demonstrateMapRangeIgnoreValue() {
	fmt.Println("=== 忽略 value 的遍历 ===")

	hash := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}

	fmt.Printf("map hash = %v\n", hash)
	fmt.Println()

	// 只获取 key，忽略 value
	fmt.Println("只获取 key（忽略 value）：")
	for key := range hash {
		fmt.Printf("  key: %s\n", key)
	}
	fmt.Println()

	// 只获取 value，忽略 key（使用空白标识符）
	fmt.Println("只获取 value（忽略 key）：")
	for _, value := range hash {
		fmt.Printf("  value: %d\n", value)
	}
	fmt.Println()
}

// demonstrateMapRangeModification 演示遍历时修改 map
func demonstrateMapRangeModification() {
	fmt.Println("=== 遍历时修改 map ===")

	hash := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Printf("原始 map hash = %v\n", hash)

	// 遍历时修改值
	fmt.Println("\n遍历时修改值：")
	for key, value := range hash {
		hash[key] = value * 10
		fmt.Printf("  修改 %s: %d -> %d\n", key, value, hash[key])
	}
	fmt.Printf("修改后 hash = %v\n", hash)
	fmt.Println()

	// 遍历时删除键值对
	fmt.Println("遍历时删除键值对：")
	for key := range hash {
		if key == "b" {
			delete(hash, key)
			fmt.Printf("  删除 key: %s\n", key)
		}
	}
	fmt.Printf("删除后 hash = %v\n", hash)
	fmt.Println()

	fmt.Println("⚠️ 注意事项：")
	fmt.Println("  - 可以在遍历时修改 map 的值")
	fmt.Println("  - 可以在遍历时删除键值对")
	fmt.Println("  - 但不要添加新的键值对（行为未定义）")
	fmt.Println()
}

// demonstrateMapRangeEmptyMap 演示空 map 的遍历
//nolint:SA6005 // 此函数用于演示 nil map 的遍历行为，这是安全的
func demonstrateMapRangeEmptyMap() {
	fmt.Println("=== 空 map 的遍历 ===")

	// 空 map
	emptyMap := map[string]int{}
	fmt.Printf("emptyMap = %v\n", emptyMap)
	fmt.Println("遍历空 map：")
	for key, value := range emptyMap {
		fmt.Printf("  %s: %d\n", key, value)
	}
	fmt.Println("（没有输出，因为 map 是空的）")
	fmt.Println()

	// nil map
	var nilMap map[string]int
	fmt.Printf("nilMap = %v\n", nilMap)
	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)
	fmt.Println("遍历 nil map：")
	for key, value := range nilMap { //nolint:SA6005 // 演示 nil map 的遍历行为，这是安全的
		fmt.Printf("  %s: %d\n", key, value)
	}
	fmt.Println("（没有输出，不会 panic）")
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 空 map 和 nil map 都可以安全遍历")
	fmt.Println("  - 遍历空 map 不会执行循环体")
	fmt.Println("  - 遍历 nil map 也不会 panic")
	fmt.Println()
}

// RangeMapDemo range 迭代映射集合完整演示
func RangeMapDemo() {
	fmt.Println("========== 1.15.4 对映射集合迭代 ==========")
	fmt.Println()
	fmt.Println("在 Go 中，使用 range 关键字迭代映射集合时：")
	fmt.Println("  1. 一种是拿到 key")
	fmt.Println("  2. 一种是拿到 key 和 value")
	fmt.Println("  3. range 关键字在迭代映射集合时，其中的 key 是乱序的（无序的）")
	fmt.Println()

	demonstrateMapRangeBasic()
	demonstrateMapUnorderedIteration()
	demonstrateMapRangeWithSortedKeys()
	demonstrateMapRangeIgnoreValue()
	demonstrateMapRangeModification()
	demonstrateMapRangeEmptyMap()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 遍历方式：")
	fmt.Println("   - for key := range map（只获取 key）")
	fmt.Println("   - for key, value := range map（获取 key 和 value）")
	fmt.Println("   - for _, value := range map（只获取 value）")
	fmt.Println()
	fmt.Println("✅ map 特性：")
	fmt.Println("   - 迭代顺序是随机的（无序的）")
	fmt.Println("   - 每次遍历顺序可能不同")
	fmt.Println("   - 不要依赖迭代顺序")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 可以在遍历时修改值和删除键值对")
	fmt.Println("   - 不要在遍历时添加新的键值对")
	fmt.Println("   - 空 map 和 nil map 都可以安全遍历")
	fmt.Println()
}
