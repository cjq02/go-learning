package mapcollection

import "fmt"

// ========== 1.14.2 使用 map 集合 ==========
//
// 获取元素：
//   <value> := <map name>[<key>]
//   <value>,<exist flag> := <map name>[<key>]
//
// 插入或修改键值对：
//   <map name>[<key>] = <value>
//
// 获取长度：
//   length := len(<map name>)
//
// 遍历 map：
//   for <key>, <value> := range <map name> { ... }
//   for <key> := range <map name> { ... }
//
// 删除键值对：
//   delete(<map name>, <key>)

// demonstrateMapGetElement 演示获取 map 元素
func demonstrateMapGetElement() {
	fmt.Println("=== 1. 获取 map 元素 ===")

	m := make(map[string]int, 10)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6

	fmt.Printf("map m = %v\n", m)
	fmt.Println()

	// 方式1：只获取值（如果键不存在，返回零值）
	value1 := m["1"]
	fmt.Printf("方式1: value1 := m[\"1\"]\n")
	fmt.Printf("value1 = %d\n", value1)

	valueUnexist := m["10"]
	fmt.Printf("valueUnexist := m[\"10\"] = %d (键不存在，返回零值)\n", valueUnexist)
	fmt.Println()

	// 方式2：获取值和存在标志（推荐方式）
	value2, exist := m["1"]
	fmt.Printf("方式2: value2, exist := m[\"1\"]\n")
	fmt.Printf("value2 = %d, exist = %v\n", value2, exist)

	valueUnexist2, exist2 := m["10"]
	fmt.Printf("valueUnexist2, exist2 := m[\"10\"]\n")
	fmt.Printf("valueUnexist2 = %d, exist2 = %v (键不存在)\n", valueUnexist2, exist2)
	fmt.Println()

	fmt.Println("⚠️ 重要提示：")
	fmt.Println("  - 方式1：如果键不存在，返回类型的零值（无法区分键是否存在）")
	fmt.Println("  - 方式2：推荐使用，可以明确知道键是否存在")
	fmt.Println()
}

// demonstrateMapInsertAndModify 演示插入和修改键值对
func demonstrateMapInsertAndModify() {
	fmt.Println("=== 2. 插入和修改键值对 ===")

	m := make(map[string]int)
	fmt.Printf("初始 map m = %v\n", m)

	// 插入键值对
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	fmt.Printf("插入后 m = %v\n", m)

	// 修改值
	fmt.Printf("\n修改前 m[\"2\"] = %d\n", m["2"])
	m["2"] = 20
	fmt.Printf("修改后 m[\"2\"] = %d\n", m["2"])
	fmt.Printf("修改后 m = %v\n", m)

	// 插入新键值对
	m["4"] = 4
	fmt.Printf("\n插入新键值对后 m = %v\n", m)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 使用 m[key] = value 可以插入或修改键值对")
	fmt.Println("  - 如果键存在，则修改值")
	fmt.Println("  - 如果键不存在，则插入新的键值对")
	fmt.Println()
}

// demonstrateMapLength 演示获取 map 长度
func demonstrateMapLength() {
	fmt.Println("=== 3. 获取 map 长度 ===")

	m := make(map[string]int)
	fmt.Printf("初始 map m = %v\n", m)
	fmt.Printf("before add, len(m) = %d\n", len(m))

	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	fmt.Printf("插入3个键值对后 m = %v\n", m)
	fmt.Printf("after add, len(m) = %d\n", len(m))

	m["10"] = 10
	fmt.Printf("再插入1个键值对后 m = %v\n", m)
	fmt.Printf("after add more, len(m) = %d\n", len(m))

	delete(m, "10")
	fmt.Printf("删除1个键值对后 m = %v\n", m)
	fmt.Printf("after delete, len(m) = %d\n", len(m))
	fmt.Println()
}

// demonstrateMapIteration 演示遍历 map
func demonstrateMapIteration() {
	fmt.Println("=== 4. 遍历 map 集合 ===")

	m := make(map[string]int, 10)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6

	fmt.Printf("map m = %v\n", m)
	fmt.Println()

	// 方式1：同时获取键和值
	fmt.Println("--- 方式1：同时获取键和值 ---")
	for key, value := range m {
		fmt.Printf("iterate map, m[%s] = %d\n", key, value)
	}
	fmt.Println()

	// 方式2：只获取键
	fmt.Println("--- 方式2：只获取键 ---")
	for key := range m {
		fmt.Printf("iterate map, key = %s\n", key)
	}
	fmt.Println()

	// 方式3：只获取值（使用空白标识符忽略键）
	fmt.Println("--- 方式3：只获取值（忽略键）---")
	for _, value := range m {
		fmt.Printf("iterate map, value = %d\n", value)
	}
	fmt.Println()

	fmt.Println("⚠️ 重要提示：")
	fmt.Println("  - map 是无序的，遍历顺序不确定")
	fmt.Println("  - 每次遍历的顺序可能不同")
	fmt.Println("  - 不要依赖遍历顺序")
	fmt.Println()
}

// demonstrateMapDelete 演示删除 map 中的键值对
func demonstrateMapDelete() {
	fmt.Println("=== 5. 删除 map 中的键值对 ===")

	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["10"] = 10

	fmt.Printf("初始 map m = %v\n", m)

	// 检查键是否存在
	_, exist10 := m["10"]
	fmt.Printf("before delete, exist 10: %v\n", exist10)

	// 删除键值对
	delete(m, "10")
	fmt.Printf("delete(m, \"10\") 后 m = %v\n", m)

	_, exist10 = m["10"]
	fmt.Printf("after delete, exist 10: %v\n", exist10)
	fmt.Println()

	// 删除不存在的键（不会报错）
	fmt.Println("--- 删除不存在的键 ---")
	fmt.Printf("删除前 m = %v\n", m)
	delete(m, "999") // 删除不存在的键
	fmt.Printf("delete(m, \"999\") 后 m = %v (不会报错)\n", m)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - delete(map, key) 删除指定的键值对")
	fmt.Println("  - 如果键不存在，不会报错，什么也不做")
	fmt.Println("  - 删除后，键值对从 map 中移除")
	fmt.Println()
}

// demonstrateMapDeleteDuringIteration 演示在遍历时删除 map 中的键
func demonstrateMapDeleteDuringIteration() {
	fmt.Println("=== 6. 在遍历时删除 map 中的键 ===")

	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5

	fmt.Printf("初始 map m = %v\n", m)
	fmt.Println()

	// 在遍历时删除所有键
	fmt.Println("--- 在遍历时删除所有键 ---")
	for key := range m {
		fmt.Printf("iterate map, will delete key: %s\n", key)
		delete(m, key)
	}
	fmt.Printf("遍历删除后 m = %v\n", m)
	fmt.Println()

	fmt.Println("⚠️ 注意事项：")
	fmt.Println("  - 可以在遍历时删除键值对")
	fmt.Println("  - 这是安全的操作，不会导致问题")
	fmt.Println("  - 但删除后，当前遍历不会立即反映变化")
	fmt.Println()
}

// demonstrateMapZeroValue 演示 map 的零值特性
func demonstrateMapZeroValue() {
	fmt.Println("=== 7. map 的零值特性 ===")

	m := map[string]int{
		"apple":  10,
		"banana": 20,
	}

	// 获取存在的键
	value1, exist1 := m["apple"]
	fmt.Printf("m[\"apple\"] = %d, exist = %v\n", value1, exist1)

	// 获取不存在的键
	value2, exist2 := m["cherry"]
	fmt.Printf("m[\"cherry\"] = %d, exist = %v (键不存在，返回零值)\n", value2, exist2)

	// 不同类型的零值
	m2 := map[string]string{
		"name": "John",
	}
	value3, exist3 := m2["age"]
	fmt.Printf("m2[\"age\"] = \"%s\", exist = %v (字符串零值是空字符串)\n", value3, exist3)

	m3 := map[string]bool{
		"active": true,
	}
	value4, exist4 := m3["deleted"]
	fmt.Printf("m3[\"deleted\"] = %v, exist = %v (布尔零值是 false)\n", value4, exist4)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 当键不存在时，返回类型的零值")
	fmt.Println("  - int 的零值是 0")
	fmt.Println("  - string 的零值是空字符串 \"\"")
	fmt.Println("  - bool 的零值是 false")
	fmt.Println("  - 使用 value, ok := m[key] 可以区分键是否存在")
	fmt.Println()
}

// MapUsageDemo map 集合使用完整演示
func MapUsageDemo() {
	fmt.Println("========== 1.14.2 使用 map 集合 ==========")
	fmt.Println()

	demonstrateMapGetElement()
	demonstrateMapInsertAndModify()
	demonstrateMapLength()
	demonstrateMapIteration()
	demonstrateMapDelete()
	demonstrateMapDeleteDuringIteration()
	demonstrateMapZeroValue()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 获取元素：")
	fmt.Println("   - value := m[key] (不推荐，无法区分键是否存在)")
	fmt.Println("   - value, ok := m[key] (推荐，可以判断键是否存在)")
	fmt.Println()
	fmt.Println("✅ 插入/修改：m[key] = value")
	fmt.Println("✅ 获取长度：len(m)")
	fmt.Println("✅ 遍历 map：for key, value := range m")
	fmt.Println("✅ 删除键值对：delete(m, key)")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - map 是无序的，遍历顺序不确定")
	fmt.Println("   - 键不存在时返回类型的零值")
	fmt.Println("   - 删除不存在的键不会报错")
	fmt.Println("   - 可以在遍历时安全地删除键值对")
	fmt.Println()
}
