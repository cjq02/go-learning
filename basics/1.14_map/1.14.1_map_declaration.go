package mapcollection

import "fmt"

// ========== 1.14.1 声明 map 集合 ==========
//
// 在 Go 中，map 集合是无序的键值对集合
// 相比切片和数组，map 集合对索引的自定义程度更高：
//   - 可以使用任意类型作为索引（键）
//   - 也可以存储任意类型的数据（值）
//
// 重要特性：
//   - map 集合中，存储的键值对的顺序是不确定的
//   - 当获取 map 集合中的值时，如果键不存在，则返回类型的零值
//
// 声明方式：
// 1. var <map name> map[<key type>]<value type>
// 2. <map name> := make(map[<key type>]<value type>)
// 3. <map name> := map[<key type>]<value type> { ... }

// demonstrateMapDeclaration1 演示方式1：仅声明 map
func demonstrateMapDeclaration1() {
	fmt.Println("=== 1. 仅声明 map ===")

	var m1 map[string]string
	fmt.Printf("var m1 map[string]string\n")
	fmt.Printf("m1 length: %d\n", len(m1))
	fmt.Printf("m1 = %v\n", m1)
	fmt.Printf("m1 == nil: %v\n", m1 == nil)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - 仅声明的 map 是 nil map")
	fmt.Println("  - nil map 不能直接存储键值对，需要先初始化")
	fmt.Println("  - 可以使用 make() 或字面量初始化")
	fmt.Println()
}

// demonstrateMapDeclaration2 演示方式2：使用 make() 初始化
func demonstrateMapDeclaration2() {
	fmt.Println("=== 2. 使用 make() 初始化 map ===")

	// 方式2.1：不指定容量
	m2 := make(map[string]string)
	fmt.Printf("m2 := make(map[string]string)\n")
	fmt.Printf("m2 length: %d\n", len(m2))
	fmt.Printf("m2 = %v\n", m2)
	fmt.Printf("m2 == nil: %v\n", m2 == nil)
	fmt.Println()

	// 方式2.2：指定容量
	m3 := make(map[string]string, 10)
	fmt.Printf("m3 := make(map[string]string, 10)\n")
	fmt.Printf("m3 length: %d\n", len(m3))
	fmt.Printf("m3 = %v\n", m3)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - make() 创建的 map 不是 nil")
	fmt.Println("  - 可以指定初始容量，减少扩容操作")
	fmt.Println("  - 指定合适的初始容量可以提高性能")
	fmt.Println()
}

// demonstrateMapDeclaration3 演示方式3：字面量初始化
func demonstrateMapDeclaration3() {
	fmt.Println("=== 3. 字面量初始化 map ===")

	// 方式3.1：空 map
	m4 := map[string]string{}
	fmt.Printf("m4 := map[string]string{}\n")
	fmt.Printf("m4 length: %d\n", len(m4))
	fmt.Printf("m4 = %v\n", m4)
	fmt.Printf("m4 == nil: %v\n", m4 == nil)
	fmt.Println()

	// 方式3.2：初始化时插入键值对
	m5 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Printf("m5 := map[string]string{...}\n")
	fmt.Printf("m5 length: %d\n", len(m5))
	fmt.Printf("m5 = %v\n", m5)
	fmt.Println()

	// 方式3.3：多行初始化（推荐格式）
	m6 := map[string]int{
		"apple":  10,
		"banana": 20,
		"cherry": 30,
	}
	fmt.Printf("m6 := map[string]int{...}\n")
	fmt.Printf("m6 = %v\n", m6)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - 字面量初始化可以同时插入键值对")
	fmt.Println("  - 空字面量 {} 创建的 map 不是 nil")
	fmt.Println("  - 键值对的顺序是不确定的（无序）")
	fmt.Println()
}

// demonstrateMapTypes 演示不同类型的 map
func demonstrateMapTypes() {
	fmt.Println("=== 4. 不同类型的 map ===")

	// 字符串到字符串的 map
	m1 := map[string]string{
		"name": "John",
		"city": "Beijing",
	}
	fmt.Printf("map[string]string: %v\n", m1)

	// 字符串到整数的 map
	m2 := map[string]int{
		"age":   25,
		"score": 100,
	}
	fmt.Printf("map[string]int: %v\n", m2)

	// 整数到字符串的 map
	m3 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Printf("map[int]string: %v\n", m3)

	// 整数到整数的 map
	m4 := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	fmt.Printf("map[int]int: %v\n", m4)

	// 布尔值作为键的 map
	m5 := map[bool]string{
		true:  "yes",
		false: "no",
	}
	fmt.Printf("map[bool]string: %v\n", m5)
	fmt.Println()
}

// demonstrateMapNilVsEmpty 演示 nil map 和空 map 的区别
func demonstrateMapNilVsEmpty() {
	fmt.Println("=== 5. nil map 和空 map 的区别 ===")

	// nil map
	var nilMap map[string]int
	fmt.Printf("var nilMap map[string]int\n")
	fmt.Printf("nilMap = %v\n", nilMap)
	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)
	fmt.Printf("len(nilMap) = %d\n", len(nilMap))

	// 空 map
	emptyMap := map[string]int{}
	fmt.Printf("\nemptyMap := map[string]int{}\n")
	fmt.Printf("emptyMap = %v\n", emptyMap)
	fmt.Printf("emptyMap == nil: %v\n", emptyMap == nil)
	fmt.Printf("len(emptyMap) = %d\n", len(emptyMap))

	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - nil map：var m map[string]int，值为 nil")
	fmt.Println("  - 空 map：m := map[string]int{}，值不为 nil")
	fmt.Println("  - 两者都可以使用，但 nil map 不能直接存储键值对")
	fmt.Println("  - 推荐使用 make() 或字面量初始化，而不是 nil map")
	fmt.Println()
}

// demonstrateMapCapacity 演示 map 的容量
func demonstrateMapCapacity() {
	fmt.Println("=== 6. map 的容量 ===")

	// 不指定容量
	m1 := make(map[string]int)
	fmt.Printf("m1 := make(map[string]int)\n")
	fmt.Printf("len(m1) = %d\n", len(m1))
	fmt.Println()

	// 指定容量
	m2 := make(map[string]int, 10)
	fmt.Printf("m2 := make(map[string]int, 10)\n")
	fmt.Printf("len(m2) = %d (初始长度为0)\n", len(m2))
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - make(map[K]V, capacity) 中的 capacity 是提示性的初始容量")
	fmt.Println("  - 实际容量可能会根据实现调整")
	fmt.Println("  - 指定合适的容量可以减少扩容操作，提高性能")
	fmt.Println("  - 但不需要精确指定，Go 会自动管理")
	fmt.Println()
}

// MapDeclarationDemo map 集合声明完整演示
func MapDeclarationDemo() {
	fmt.Println("========== 1.14.1 声明 map 集合 ==========")
	fmt.Println()
	fmt.Println("在 Go 中，map 集合是无序的键值对集合。")
	fmt.Println("相比切片和数组，map 集合对索引的自定义程度更高：")
	fmt.Println("  - 可以使用任意类型作为索引（键）")
	fmt.Println("  - 也可以存储任意类型的数据（值）")
	fmt.Println()
	fmt.Println("重要特性：")
	fmt.Println("  - map 集合中，存储的键值对的顺序是不确定的")
	fmt.Println("  - 当获取 map 集合中的值时，如果键不存在，则返回类型的零值")
	fmt.Println()

	demonstrateMapDeclaration1()
	demonstrateMapDeclaration2()
	demonstrateMapDeclaration3()
	demonstrateMapTypes()
	demonstrateMapNilVsEmpty()
	demonstrateMapCapacity()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 声明方式：")
	fmt.Println("   1. var m map[K]V (nil map)")
	fmt.Println("   2. m := make(map[K]V) 或 make(map[K]V, capacity)")
	fmt.Println("   3. m := map[K]V{} 或 map[K]V{key: value, ...}")
	fmt.Println()
	fmt.Println("✅ map 特性：")
	fmt.Println("   - 无序的键值对集合")
	fmt.Println("   - 键可以是任意可比较类型")
	fmt.Println("   - 值可以是任意类型")
	fmt.Println("   - 键不存在时返回零值")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - nil map 不能直接存储键值对，需要先初始化")
	fmt.Println("   - map 是无序的，遍历顺序不确定")
	fmt.Println("   - 指定合适的初始容量可以提高性能")
	fmt.Println()
}
