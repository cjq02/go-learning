package rangeiteration

import (
	"fmt"
	"reflect"
)

// ========== 1.15.2 对数组与切片迭代 ==========
//
// 在 Go 中，实际代码执行过程中，使用 range 迭代数组和切片，它们两者的体验是相同的
//
// 遍历方式：
// 1. 只获取索引：for index := range array/slice
// 2. 获取索引和值：for index, value := range array/slice
//
// 切片与数组相比，特殊的地方就在于其长度可变
// 所以在构成二维时，切片中元素的数量可以随意设置，而数组是定长的

// demonstrateOneDimensionalRange 演示遍历一维数组与切片
func demonstrateOneDimensionalRange() {
	fmt.Println("=== 1.15.2.1 遍历一维数组与切片 ===")

	array := [...]int{1, 2, 3}
	slice := []int{4, 5, 6}

	fmt.Printf("数组 array = %v\n", array)
	fmt.Printf("切片 slice = %v\n", slice)
	fmt.Println()

	// 方法1：只拿到数组/切片的下标索引
	fmt.Println("--- 方法1：只获取索引 ---")
	for index := range array {
		fmt.Printf("array -- index=%d value=%d\n", index, array[index])
	}
	for index := range slice {
		fmt.Printf("slice -- index=%d value=%d\n", index, slice[index])
	}
	fmt.Println()

	// 方法2：同时拿到数组/切片的下标索引和对应的值
	fmt.Println("--- 方法2：获取索引和值 ---")
	for index, value := range array {
		fmt.Printf("array -- index=%d index value=%d\n", index, array[index])
		fmt.Printf("array -- index=%d range value=%d\n", index, value)
	}
	fmt.Println()
	for index, value := range slice {
		fmt.Printf("slice -- index=%d index value=%d\n", index, slice[index])
		fmt.Printf("slice -- index=%d range value=%d\n", index, value)
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 数组和切片使用 range 迭代的体验完全相同")
	fmt.Println("  - 方法1：只获取索引，需要通过 array[index] 访问值")
	fmt.Println("  - 方法2：同时获取索引和值，value 是元素的副本")
	fmt.Println()
}

// demonstrateTwoDimensionalRange 演示遍历二维数组与切片
func demonstrateTwoDimensionalRange() {
	fmt.Println("=== 1.15.2.2 遍历二维数组与切片 ===")

	array := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	slice := [][]int{{1, 2}, {3}}

	fmt.Printf("二维数组 array = %v\n", array)
	fmt.Printf("二维切片 slice = %v\n", slice)
	fmt.Println()

	// 只拿到行的索引
	fmt.Println("--- 只获取行索引 ---")
	for index := range array {
		// array[index] 类型是一维数组
		fmt.Printf("array -- index=%d, type=%v, value=%v\n", index, reflect.TypeOf(array[index]), array[index])
	}
	fmt.Println()
	for index := range slice {
		// slice[index] 类型是一维切片
		fmt.Printf("slice -- index=%d, type=%v, value=%v\n", index, reflect.TypeOf(slice[index]), slice[index])
	}
	fmt.Println()

	// 拿到行索引和该行的数据
	fmt.Println("--- 获取行索引和行数据 ---")
	fmt.Println("遍历二维数组：")
	for rowIndex, rowValue := range array {
		fmt.Printf("  rowIndex=%d, type=%v, value=%v\n", rowIndex, reflect.TypeOf(rowValue), rowValue)
	}
	fmt.Println()
	fmt.Println("遍历二维切片：")
	for rowIndex, rowValue := range slice {
		fmt.Printf("  rowIndex=%d, type=%v, value=%v\n", rowIndex, reflect.TypeOf(rowValue), rowValue)
	}
	fmt.Println()

	// 双重遍历，拿到每个元素的值
	fmt.Println("--- 双重遍历，获取每个元素 ---")
	fmt.Println("遍历二维数组：")
	for rowIndex, rowValue := range array {
		for colIndex, colValue := range rowValue {
			fmt.Printf("array[%d][%d]=%d ", rowIndex, colIndex, colValue)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("遍历二维切片：")
	for rowIndex, rowValue := range slice {
		for colIndex, colValue := range rowValue {
			fmt.Printf("slice[%d][%d]=%d ", rowIndex, colIndex, colValue)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 二维数组：每行的长度固定（array[0] 和 array[1] 都是 [3]int）")
	fmt.Println("  - 二维切片：每行的长度可以不同（slice[0] 是 [2]int，slice[1] 是 [1]int）")
	fmt.Println("  - 使用 range 迭代时，两者体验完全一致")
	fmt.Println()
}

// demonstrateSliceVariableLength 演示切片的可变长度特性
func demonstrateSliceVariableLength() {
	fmt.Println("=== 切片的可变长度特性 ===")

	// 二维数组：每行长度固定
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("二维数组 array = %v\n", array)
	fmt.Printf("array[0] 长度 = %d\n", len(array[0]))
	fmt.Printf("array[1] 长度 = %d\n", len(array[1]))
	fmt.Println()

	// 二维切片：每行长度可以不同
	slice := [][]int{
		{1, 2, 3},      // 3个元素
		{4, 5},         // 2个元素
		{6},            // 1个元素
		{7, 8, 9, 10}, // 4个元素
	}
	fmt.Printf("二维切片 slice = %v\n", slice)
	for i, row := range slice {
		fmt.Printf("slice[%d] 长度 = %d, 值 = %v\n", i, len(row), row)
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 数组：每行长度固定，必须在声明时确定")
	fmt.Println("  - 切片：每行长度可变，可以动态设置")
	fmt.Println("  - 这是切片相比数组的优势之一")
	fmt.Println()
}

// demonstrateRangeValueCopy 演示 range 返回的是值副本
func demonstrateRangeValueCopy() {
	fmt.Println("=== range 返回的是值副本 ===")

	slice := []int{1, 2, 3}
	fmt.Printf("原始切片 slice = %v\n", slice)

	fmt.Println("\n遍历时修改 value（不会影响原切片）：")
	for index, value := range slice {
		value = value * 10 // 修改 value 不会影响原切片
		fmt.Printf("  index=%d, value=%d (这是副本)\n", index, value)
	}
	fmt.Printf("遍历后 slice = %v (未改变)\n", slice)
	fmt.Println()

	fmt.Println("遍历时通过索引修改（会影响原切片）：")
	for index := range slice {
		slice[index] = slice[index] * 10 // 通过索引修改会影响原切片
	}
	fmt.Printf("修改后 slice = %v (已改变)\n", slice)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - range 返回的 value 是元素的副本")
	fmt.Println("  - 修改 value 不会影响原数组/切片")
	fmt.Println("  - 需要通过索引 array[index] 或 slice[index] 修改")
	fmt.Println()
}

// demonstrateRangeWithBlankIdentifier 演示使用空白标识符
func demonstrateRangeWithBlankIdentifier() {
	fmt.Println("=== 使用空白标识符忽略值 ===")

	slice := []int{10, 20, 30, 40, 50}
	fmt.Printf("切片 slice = %v\n", slice)
	fmt.Println()

	// 只获取值，忽略索引
	fmt.Println("--- 只获取值，忽略索引 ---")
	sum := 0
	for _, value := range slice {
		sum += value
	}
	fmt.Printf("所有元素的和 = %d\n", sum)
	fmt.Println()

	// 只获取索引，忽略值
	fmt.Println("--- 只获取索引，忽略值 ---")
	for index := range slice {
		fmt.Printf("索引 %d\n", index)
	}
	fmt.Println()

	// 二维数组：只获取行索引
	fmt.Println("--- 二维数组：只获取行索引 ---")
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for rowIndex := range array {
		fmt.Printf("行索引 %d, 行数据 %v\n", rowIndex, array[rowIndex])
	}
	fmt.Println()
}

// RangeArraySliceDemo range 迭代数组与切片完整演示
func RangeArraySliceDemo() {
	fmt.Println("========== 1.15.2 对数组与切片迭代 ==========")
	fmt.Println()
	fmt.Println("在 Go 中，实际代码执行过程中，")
	fmt.Println("使用 range 迭代数组和切片，它们两者的体验是相同的。")
	fmt.Println()
	fmt.Println("切片与数组相比，特殊的地方就在于其长度可变，")
	fmt.Println("所以在构成二维时，切片中元素的数量可以随意设置，")
	fmt.Println("而数组是定长的。")
	fmt.Println()

	demonstrateOneDimensionalRange()
	demonstrateTwoDimensionalRange()
	demonstrateSliceVariableLength()
	demonstrateRangeValueCopy()
	demonstrateRangeWithBlankIdentifier()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 数组和切片使用 range 迭代的体验完全相同")
	fmt.Println("✅ 遍历方式：")
	fmt.Println("   - for index := range array/slice（只获取索引）")
	fmt.Println("   - for index, value := range array/slice（获取索引和值）")
	fmt.Println("   - for _, value := range array/slice（只获取值）")
	fmt.Println()
	fmt.Println("✅ 二维数组/切片：")
	fmt.Println("   - 使用嵌套 range 遍历每个元素")
	fmt.Println("   - 数组每行长度固定，切片每行长度可变")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - range 返回的 value 是元素的副本")
	fmt.Println("   - 修改 value 不会影响原数组/切片")
	fmt.Println("   - 需要通过索引修改原数组/切片")
	fmt.Println()
}

