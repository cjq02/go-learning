package mapcollection

import "fmt"

// ========== 1.14.3 map 作为参数 ==========
//
// map 集合也是引用类型，和切片一样
// 将 map 集合作为参数传给函数或者赋值给另一个变量，
// 它们都指向同一个底层数据结构
// 对 map 集合的修改，都会影响到原始实参

// demonstrateMapAsParameter 演示 map 作为参数传递
func demonstrateMapAsParameter() {
	fmt.Println("=== 1. map 作为参数传递 ===")

	m := make(map[string]int)
	m["a"] = 1
	fmt.Printf("调用函数前 m = %v\n", m)

	receiveMap(m)
	fmt.Printf("调用函数后 m = %v\n", m)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - map 是引用类型，传递的是引用")
	fmt.Println("  - 函数内修改 map 会影响原 map")
	fmt.Println("  - 因为它们指向同一个底层数据结构")
	fmt.Println()
}

// receiveMap 接收 map 作为参数
func receiveMap(param map[string]int) {
	fmt.Printf("  在 receiveMap 函数中，修改前 param[\"a\"] = %d\n", param["a"])
	param["a"] = 2
	param["b"] = 3
	fmt.Printf("  在 receiveMap 函数中，修改后 param = %v\n", param)
}

// demonstrateMapAssignment 演示 map 赋值
func demonstrateMapAssignment() {
	fmt.Println("=== 2. map 赋值 ===")

	m1 := make(map[string]int)
	m1["a"] = 1
	m1["b"] = 2
	fmt.Printf("m1 = %v\n", m1)

	// 赋值给另一个变量
	m2 := m1
	fmt.Printf("m2 := m1 后，m2 = %v\n", m2)

	// 修改 m2
	m2["c"] = 3
	fmt.Printf("修改 m2[\"c\"] = 3 后:\n")
	fmt.Printf("m1 = %v (m1也改变了)\n", m1)
	fmt.Printf("m2 = %v\n", m2)

	// 修改 m1
	m1["d"] = 4
	fmt.Printf("\n修改 m1[\"d\"] = 4 后:\n")
	fmt.Printf("m1 = %v\n", m1)
	fmt.Printf("m2 = %v (m2也改变了)\n", m2)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - m1 和 m2 指向同一个底层数据结构")
	fmt.Println("  - 修改 m1 会影响 m2，修改 m2 会影响 m1")
	fmt.Println("  - 这是引用类型的特性")
	fmt.Println()
}

// demonstrateMapVsArray 对比 map 和数组作为参数的区别
func demonstrateMapVsArray() {
	fmt.Println("=== 3. map vs 数组作为参数对比 ===")

	// 数组：值类型
	arr := [3]int{1, 2, 3}
	fmt.Printf("数组 arr = %v\n", arr)
	modifyArray(arr)
	fmt.Printf("调用函数后 arr = %v (数组未改变，因为是值传递)\n", arr)
	fmt.Println()

	// map：引用类型
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Printf("map m = %v\n", m)
	modifyMap(m)
	fmt.Printf("调用函数后 m = %v (map已改变，因为是引用传递)\n", m)
	fmt.Println()

	fmt.Println("对比总结：")
	fmt.Println("  - 数组：值类型，传递时复制整个数组，函数内修改不影响原数组")
	fmt.Println("  - map：引用类型，传递时复制引用，函数内修改影响原 map")
	fmt.Println()
}

// modifyArray 修改数组（值传递）
func modifyArray(arr [3]int) {
	fmt.Printf("  在 modifyArray 函数中，修改前 arr = %v\n", arr)
	arr[0] = 999
	fmt.Printf("  在 modifyArray 函数中，修改后 arr = %v (这是副本)\n", arr)
}

// modifyMap 修改 map（引用传递）
func modifyMap(m map[string]int) {
	fmt.Printf("  在 modifyMap 函数中，修改前 m = %v\n", m)
	m["a"] = 999
	m["d"] = 4
	fmt.Printf("  在 modifyMap 函数中，修改后 m = %v (影响原map)\n", m)
}

// demonstrateMapVsSlice 对比 map 和切片作为参数的区别
func demonstrateMapVsSlice() {
	fmt.Println("=== 4. map vs 切片作为参数对比 ===")

	// 切片：引用类型
	s := []int{1, 2, 3}
	fmt.Printf("切片 s = %v\n", s)
	modifySlice(s)
	fmt.Printf("调用函数后 s = %v (切片已改变，因为是引用传递)\n", s)
	fmt.Println()

	// map：引用类型
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("map m = %v\n", m)
	modifyMapForComparison(m)
	fmt.Printf("调用函数后 m = %v (map已改变，因为是引用传递)\n", m)
	fmt.Println()

	fmt.Println("对比总结：")
	fmt.Println("  - 切片：引用类型，传递切片头（指针+长度+容量）")
	fmt.Println("  - map：引用类型，传递 map 引用")
	fmt.Println("  - 两者都是引用类型，函数内修改都会影响原值")
	fmt.Println("  - 但切片使用 append() 可能触发扩容，创建新数组")
	fmt.Println()
}

// modifySlice 修改切片
func modifySlice(s []int) {
	fmt.Printf("  在 modifySlice 函数中，修改前 s = %v\n", s)
	s[0] = 999
	fmt.Printf("  在 modifySlice 函数中，修改后 s = %v\n", s)
}

// modifyMapForComparison 修改 map（用于对比）
func modifyMapForComparison(m map[string]int) {
	fmt.Printf("  在 modifyMapForComparison 函数中，修改前 m = %v\n", m)
	m["a"] = 999
	m["c"] = 3
	fmt.Printf("  在 modifyMapForComparison 函数中，修改后 m = %v\n", m)
}

// demonstrateMapNilParameter 演示 nil map 作为参数
func demonstrateMapNilParameter() {
	fmt.Println("=== 5. nil map 作为参数 ===")

	var nilMap map[string]int
	fmt.Printf("nilMap = %v\n", nilMap)
	fmt.Printf("nilMap == nil: %v\n", nilMap == nil)

	// nil map 可以作为参数传递
	handleNilMap(nilMap)
	fmt.Println()

	// 在函数内初始化 nil map
	initializeMapInFunc(nilMap)
	fmt.Printf("函数返回后 nilMap = %v (仍然是 nil)\n", nilMap)
	fmt.Println()

	fmt.Println("⚠️ 注意事项：")
	fmt.Println("  - nil map 可以作为参数传递")
	fmt.Println("  - 但 nil map 不能直接存储键值对")
	fmt.Println("  - 需要在函数内先初始化才能使用")
	fmt.Println()
}

// handleNilMap 处理 nil map
func handleNilMap(m map[string]int) {
	fmt.Printf("  在 handleNilMap 函数中，m == nil: %v\n", m == nil)
	if m == nil {
		fmt.Println("  map 是 nil，需要先初始化")
		m = make(map[string]int)
		m["key"] = 1
		fmt.Printf("  初始化后 m = %v\n", m)
	}
}

// initializeMapInFunc 在函数内初始化 map
func initializeMapInFunc(m map[string]int) {
	fmt.Printf("  在 initializeMapInFunc 函数中，m == nil: %v\n", m == nil)
	m = make(map[string]int)
	m["key"] = 1
	fmt.Printf("  初始化后 m = %v\n", m)
	fmt.Println("  说明：函数内初始化不会影响外部的 nil map")
}

// MapAsParameterDemo map 作为参数完整演示
func MapAsParameterDemo() {
	fmt.Println("========== 1.14.3 map 作为参数 ==========")
	fmt.Println()
	fmt.Println("map 集合也是引用类型，和切片一样")
	fmt.Println("将 map 集合作为参数传给函数或者赋值给另一个变量，")
	fmt.Println("它们都指向同一个底层数据结构")
	fmt.Println("对 map 集合的修改，都会影响到原始实参")
	fmt.Println()

	demonstrateMapAsParameter()
	demonstrateMapAssignment()
	demonstrateMapVsArray()
	demonstrateMapVsSlice()
	demonstrateMapNilParameter()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ map 是引用类型")
	fmt.Println("✅ 作为参数传递时，传递的是引用")
	fmt.Println("✅ 函数内修改 map 会影响原 map")
	fmt.Println("✅ 赋值给另一个变量，两个变量指向同一个 map")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - map 和切片都是引用类型")
	fmt.Println("   - 但 map 没有类似切片 append() 的扩容机制")
	fmt.Println("   - nil map 可以作为参数传递，但需要先初始化才能使用")
	fmt.Println()
}
