package array

import "fmt"

// ========== 1.12.2 访问数组 ==========
//
// 访问数组的方式：
// 1. 使用下标读取数组中的元素：<value> := <array name>[<position>]
// 2. 使用 range 遍历：for <i>,<v> := range <array name> { ... }
// 3. 获取数组长度：<length variable name> := len(<array name>)

// demonstrateArrayAccessByIndex 演示使用下标访问数组
func demonstrateArrayAccessByIndex() {
	fmt.Println("=== 1. 使用下标访问数组元素 ===")

	a := [5]int{5, 4, 3, 2, 1}
	fmt.Printf("数组 a = %v\n", a)

	// 方式1：使用下标读取数组中的元素
	element := a[2]
	fmt.Printf("a[2] = %d\n", element)

	// 访问不同位置的元素
	fmt.Printf("a[0] = %d\n", a[0])
	fmt.Printf("a[4] = %d\n", a[4])

	// 修改数组元素
	a[1] = 100
	fmt.Printf("修改 a[1] = 100 后，数组 a = %v\n", a)

	// ⚠️ 注意：数组下标从 0 开始，不能越界
	// a[5] = 10  // 编译错误：数组越界
	fmt.Println("说明：数组下标从 0 开始，有效范围是 0 到 len(array)-1")
	fmt.Println()
}

// demonstrateArrayAccessByRange 演示使用 range 遍历数组
func demonstrateArrayAccessByRange() {
	fmt.Println("=== 2. 使用 range 遍历数组 ===")

	a := [5]int{5, 4, 3, 2, 1}
	fmt.Printf("数组 a = %v\n", a)

	// 方式2：使用 range 遍历（同时获取索引和值）
	fmt.Println("\n--- 同时获取索引和值 ---")
	for i, v := range a {
		fmt.Printf("index = %d, value = %d\n", i, v)
	}

	// 只获取索引
	fmt.Println("\n--- 只获取索引 ---")
	for i := range a {
		fmt.Printf("only index, index = %d\n", i)
	}

	// 只获取值（使用空白标识符 _ 忽略索引）
	fmt.Println("\n--- 只获取值（忽略索引）---")
	for _, v := range a {
		fmt.Printf("only value, value = %d\n", v)
	}
	fmt.Println()
}

// demonstrateArrayLength 演示获取数组长度
func demonstrateArrayLength() {
	fmt.Println("=== 3. 获取数组长度 ===")

	a := [5]int{5, 4, 3, 2, 1}
	fmt.Printf("数组 a = %v\n", a)

	// 读取数组长度
	length := len(a)
	fmt.Printf("len(a) = %d\n", length)

	// 使用 len() 和 for 循环遍历数组
	fmt.Println("\n--- 使用 len() 和 for 循环遍历数组 ---")
	for i := 0; i < len(a); i++ {
		fmt.Printf("use len(), index = %d, value = %d\n", i, a[i])
	}
	fmt.Println()
}

// demonstrateArrayAccessPatterns 演示数组访问的各种模式
func demonstrateArrayAccessPatterns() {
	fmt.Println("=== 4. 数组访问的各种模式 ===")

	arr := [5]string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Printf("数组 arr = %v\n", arr)

	// 模式1：顺序访问
	fmt.Println("\n--- 顺序访问 ---")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %s\n", i, arr[i])
	}

	// 模式2：逆序访问
	fmt.Println("\n--- 逆序访问 ---")
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("arr[%d] = %s\n", i, arr[i])
	}

	// 模式3：访问特定范围
	fmt.Println("\n--- 访问特定范围（索引 1 到 3）---")
	for i := 1; i <= 3; i++ {
		fmt.Printf("arr[%d] = %s\n", i, arr[i])
	}

	// 模式4：访问特定元素
	fmt.Println("\n--- 访问特定元素 ---")
	fmt.Printf("第一个元素: arr[0] = %s\n", arr[0])
	fmt.Printf("最后一个元素: arr[%d] = %s\n", len(arr)-1, arr[len(arr)-1])
	fmt.Println()
}

// demonstrateArrayBounds 演示数组边界检查
func demonstrateArrayBounds() {
	fmt.Println("=== 5. 数组边界检查 ===")

	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("数组 arr = %v\n", arr)
	fmt.Printf("数组长度: len(arr) = %d\n", len(arr))

	// 有效索引范围
	fmt.Println("\n有效索引范围：0 到", len(arr)-1)

	// 访问边界元素
	fmt.Printf("第一个元素 arr[0] = %d\n", arr[0])
	fmt.Printf("最后一个元素 arr[%d] = %d\n", len(arr)-1, arr[len(arr)-1])

	// ⚠️ 注意：数组越界会导致运行时 panic
	fmt.Println("\n⚠️ 注意：")
	fmt.Println("  - arr[-1] 会导致编译错误")
	fmt.Println("  - arr[len(arr)] 会导致运行时 panic: index out of range")
	fmt.Println("  - 访问前应检查索引是否在有效范围内")
	fmt.Println()
}

// demonstrateArrayModification 演示数组元素的修改
func demonstrateArrayModification() {
	fmt.Println("=== 6. 修改数组元素 ===")

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("原始数组: arr = %v\n", arr)

	// 修改单个元素
	arr[2] = 100
	fmt.Printf("修改 arr[2] = 100 后: arr = %v\n", arr)

	// 通过循环修改多个元素
	fmt.Println("\n--- 通过循环修改多个元素 ---")
	for i := 0; i < len(arr); i++ {
		arr[i] = (i + 1) * 10
	}
	fmt.Printf("修改后: arr = %v\n", arr)

	// 通过 range 修改（注意：range 返回的是值的副本）
	fmt.Println("\n--- 注意：range 返回的是值的副本 ---")
	arr2 := [3]int{1, 2, 3}
	fmt.Printf("原始数组: arr2 = %v\n", arr2)
	for _, v := range arr2 {
		v = v * 10 // 这不会修改原数组
		fmt.Printf("range 中的 v = %d (这是副本)\n", v)
	}
	fmt.Printf("数组未改变: arr2 = %v\n", arr2)

	// 正确的方式：通过索引修改
	fmt.Println("\n--- 正确方式：通过索引修改 ---")
	for i := range arr2 {
		arr2[i] = arr2[i] * 10
	}
	fmt.Printf("修改后: arr2 = %v\n", arr2)
	fmt.Println()
}

// ArrayAccessDemo 数组访问完整演示
func ArrayAccessDemo() {
	fmt.Println("========== 1.12.2 访问数组 ==========")
	fmt.Println()
	fmt.Println("访问数组的方式：")
	fmt.Println("1. 使用下标读取数组中的元素：<value> := <array name>[<position>]")
	fmt.Println("2. 使用 range 遍历：for <i>,<v> := range <array name> { ... }")
	fmt.Println("3. 获取数组长度：<length variable name> := len(<array name>)")
	fmt.Println()

	demonstrateArrayAccessByIndex()
	demonstrateArrayAccessByRange()
	demonstrateArrayLength()
	demonstrateArrayAccessPatterns()
	demonstrateArrayBounds()
	demonstrateArrayModification()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 使用下标访问：array[index]")
	fmt.Println("✅ 使用 range 遍历：for i, v := range array")
	fmt.Println("✅ 获取长度：len(array)")
	fmt.Println("✅ 数组下标从 0 开始，有效范围是 0 到 len(array)-1")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 数组越界会导致运行时 panic")
	fmt.Println("   - range 返回的值是副本，不能直接修改原数组")
	fmt.Println("   - 修改数组元素需要使用索引：array[i] = value")
	fmt.Println()
}
