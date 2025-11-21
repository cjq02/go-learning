package array

import "fmt"

// ========== 1.12.4 数组作为参数 ==========
//
// 数组的部分特性类似基础数据类型，当数组作为参数传递时，在函数中并不能改变外部实参的值
// 如果想要修改外部实参的值，需要把数组的指针作为参数传递给函数
//
// 原因：数组是值类型，传递时会复制整个数组

// Custom 自定义结构体类型
type Custom struct {
	i int
}

// 全局数组：指针数组
var carr [5]*Custom = [5]*Custom{
	{6},
	{7},
	{8},
	{9},
	{10},
}

// demonstrateArrayAsValueParameter 演示数组作为值参数传递
func demonstrateArrayAsValueParameter() {
	fmt.Println("=== 1. 数组作为值参数传递 ===")
	fmt.Println("说明：数组是值类型，传递时会复制整个数组")

	a := [5]int{5, 4, 3, 2, 1}
	fmt.Printf("调用函数前，a = %v\n", a)

	receiveArray(a)
	fmt.Printf("调用函数后，a = %v\n", a)
	fmt.Println("⚠️ 注意：函数内部的修改不会影响外部的数组")
	fmt.Println()
}

// receiveArray 接收数组作为值参数
func receiveArray(param [5]int) {
	fmt.Printf("在 receiveArray 函数中，修改前 param = %v\n", param)
	param[1] = -5
	fmt.Printf("在 receiveArray 函数中，修改后 param = %v\n", param)
	fmt.Println("说明：param 是 a 的副本，修改 param 不会影响 a")
}

// demonstrateArrayAsPointerParameter 演示数组作为指针参数传递
func demonstrateArrayAsPointerParameter() {
	fmt.Println("=== 2. 数组作为指针参数传递 ===")
	fmt.Println("说明：传递数组指针可以修改外部数组")

	a := [5]int{5, 4, 3, 2, 1}
	fmt.Printf("调用函数前，a = %v\n", a)

	receiveArrayPointer(&a)
	fmt.Printf("调用函数后，a = %v\n", a)
	fmt.Println("✅ 注意：函数内部的修改会影响外部的数组")
	fmt.Println()
}

// receiveArrayPointer 接收数组指针作为参数
func receiveArrayPointer(param *[5]int) {
	fmt.Printf("在 receiveArrayPointer 函数中，修改前 param = %v\n", param)
	param[1] = -5
	fmt.Printf("在 receiveArrayPointer 函数中，修改后 param = %v\n", param)
	fmt.Println("说明：param 是指向 a 的指针，修改 param 会影响 a")
}

// demonstrateArrayOfPointers 演示指针数组作为参数传递
func demonstrateArrayOfPointers() {
	fmt.Println("=== 3. 指针数组作为参数传递 ===")
	fmt.Println("说明：数组的元素是指针类型时，传递的是指针数组的副本，但指针指向的对象是共享的")

	fmt.Println("--- 在 main 函数中 ---")
	for i := range carr {
		fmt.Printf("carr[%d] = %p, value = %v\n", i, carr[i], (*carr[i]).i)
	}

	printFuncParamPointer(carr)

	fmt.Println("--- 在 main 函数中（调用后）---")
	for i := range carr {
		fmt.Printf("carr[%d] = %p, value = %v\n", i, carr[i], (*carr[i]).i)
	}
	fmt.Println()
}

// printFuncParamPointer 接收指针数组作为参数
func printFuncParamPointer(param [5]*Custom) {
	fmt.Println("--- 在 printFuncParamPointer 函数中 ---")
	for i := range param {
		(*param[i]).i = (*param[i]).i + 1
		fmt.Printf("param[%d] = %p, value = %v\n", i, param[i], (*param[i]).i)
	}
	fmt.Println("说明：param 是 carr 的副本，但 param[i] 和 carr[i] 指向同一个对象")
}

// demonstrateArrayParameterComparison 对比数组值传递和指针传递
func demonstrateArrayParameterComparison() {
	fmt.Println("=== 4. 数组值传递 vs 指针传递对比 ===")

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}

	fmt.Printf("原始数组 arr1 = %v\n", arr1)
	fmt.Printf("原始数组 arr2 = %v\n", arr2)

	// 值传递
	modifyByValue(arr1)
	fmt.Printf("\n值传递后 arr1 = %v\n", arr1)

	// 指针传递
	modifyByPointer(&arr2)
	fmt.Printf("指针传递后 arr2 = %v\n", arr2)

	fmt.Println("\n总结：")
	fmt.Println("  - 值传递：函数内部修改不影响外部（arr1 未改变）")
	fmt.Println("  - 指针传递：函数内部修改影响外部（arr2 已改变）")
	fmt.Println()
}

// modifyByValue 通过值传递修改数组
func modifyByValue(arr [3]int) {
	fmt.Println("在 modifyByValue 中修改数组")
	for i := range arr {
		arr[i] = arr[i] * 10
	}
	fmt.Printf("修改后 arr = %v（这是副本）\n", arr)
}

// modifyByPointer 通过指针传递修改数组
func modifyByPointer(arr *[3]int) {
	fmt.Println("在 modifyByPointer 中修改数组")
	for i := range arr {
		arr[i] = arr[i] * 10
	}
	fmt.Printf("修改后 arr = %v（这是原数组）\n", arr)
}

// demonstrateArrayParameterPerformance 演示数组参数传递的性能考虑
func demonstrateArrayParameterPerformance() {
	fmt.Println("=== 5. 数组参数传递的性能考虑 ===")

	fmt.Println("小数组（3个元素）：")
	fmt.Printf("  值传递：复制 3 个 int（24 字节）\n")
	fmt.Printf("  指针传递：复制 1 个指针（8 字节）\n")

	fmt.Println("\n大数组（1000个元素）：")
	fmt.Printf("  值传递：复制 1000 个 int（8000 字节）\n")
	fmt.Printf("  指针传递：复制 1 个指针（8 字节）\n")

	fmt.Println("\n性能建议：")
	fmt.Println("  - 小数组：值传递和指针传递性能差异不大")
	fmt.Println("  - 大数组：建议使用指针传递，避免大量数据复制")
	fmt.Println("  - 或者：使用切片（slice），切片是引用类型")
	fmt.Println()
}

// demonstrateArrayParameterBestPractices 演示数组参数的最佳实践
func demonstrateArrayParameterBestPractices() {
	fmt.Println("=== 6. 数组参数的最佳实践 ===")

	arr := [5]int{1, 2, 3, 4, 5}

	// 实践1：只读操作，使用值传递
	fmt.Println("--- 实践1：只读操作，使用值传递 ---")
	sum := calculateSum(arr)
	fmt.Printf("数组 %v 的和 = %d\n", arr, sum)
	fmt.Println("说明：只读操作不需要修改原数组，值传递即可")

	// 实践2：需要修改，使用指针传递
	fmt.Println("\n--- 实践2：需要修改，使用指针传递 ---")
	fmt.Printf("修改前 arr = %v\n", arr)
	resetArray(&arr)
	fmt.Printf("修改后 arr = %v\n", arr)

	// 实践3：使用切片（推荐）
	fmt.Println("\n--- 实践3：使用切片（推荐）---")
	fmt.Println("说明：在实际开发中，更推荐使用切片（slice）")
	fmt.Println("     切片是引用类型，传递时不会复制数据")
	fmt.Println("     例如：func processSlice(s []int) { ... }")
	fmt.Println()
}

// calculateSum 计算数组元素的和（只读操作）
func calculateSum(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// resetArray 重置数组（需要修改）
func resetArray(arr *[5]int) {
	for i := range arr {
		arr[i] = 0
	}
}

// ArrayAsParameterDemo 数组作为参数完整演示
func ArrayAsParameterDemo() {
	fmt.Println("========== 1.12.4 数组作为参数 ==========")
	fmt.Println()
	fmt.Println("数组的部分特性类似基础数据类型，当数组作为参数传递时，")
	fmt.Println("在函数中并不能改变外部实参的值。")
	fmt.Println()
	fmt.Println("如果想要修改外部实参的值，需要把数组的指针作为参数传递给函数。")
	fmt.Println()
	fmt.Println("原因：数组是值类型，传递时会复制整个数组")
	fmt.Println()

	demonstrateArrayAsValueParameter()
	demonstrateArrayAsPointerParameter()
	demonstrateArrayOfPointers()
	demonstrateArrayParameterComparison()
	demonstrateArrayParameterPerformance()
	demonstrateArrayParameterBestPractices()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 数组是值类型，传递时会复制整个数组")
	fmt.Println("✅ 值传递：函数内部修改不影响外部数组")
	fmt.Println("✅ 指针传递：函数内部修改会影响外部数组")
	fmt.Println("✅ 指针数组：传递的是指针数组的副本，但指针指向的对象是共享的")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 大数组传递时，建议使用指针避免大量数据复制")
	fmt.Println("   - 实际开发中，更推荐使用切片（slice）")
	fmt.Println("   - 切片是引用类型，传递时不会复制数据")
	fmt.Println()
}
