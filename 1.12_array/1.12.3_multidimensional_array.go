package array

import "fmt"

// ========== 1.12.3 多维数组 ==========
//
// 多维数组：数组的数组
// Go 中没有限制多维数组的嵌套层数
//
// 声明方式：var <array name> [<length1>][<length2>]... <type>
// 访问方式：与访问普通数组的方式一致，使用多个下标

// demonstrateMultidimensionalArrayDeclaration 演示多维数组的声明
func demonstrateMultidimensionalArrayDeclaration() {
	fmt.Println("=== 1. 多维数组的声明 ===")

	// 二维数组
	fmt.Println("--- 二维数组 ---")
	a := [3][2]int{
		{0, 1},
		{2, 3},
		{4, 5},
	}
	fmt.Printf("a = %v\n", a)
	fmt.Printf("a 的类型: %T\n", a)
	fmt.Println("说明：a 是一个 3x2 的二维数组")

	// 三维数组
	fmt.Println("\n--- 三维数组 ---")
	b := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	fmt.Printf("b = %v\n", b)
	fmt.Printf("b 的类型: %T\n", b)
	fmt.Println("说明：b 是一个 3x2x2 的三维数组")

	// 也可以省略各个位置的初始化，在后续代码中赋值
	fmt.Println("\n--- 省略初始化，后续赋值 ---")
	c := [3][3][3]int{}
	fmt.Printf("初始化后 c = %v\n", c)
	c[2][2][1] = 5
	c[1][2][1] = 4
	fmt.Printf("赋值后 c = %v\n", c)
	fmt.Println("说明：未初始化的元素为类型零值（0）")
	fmt.Println()
}

// demonstrateMultidimensionalArrayAccess 演示多维数组的访问
func demonstrateMultidimensionalArrayAccess() {
	fmt.Println("=== 2. 多维数组的访问 ===")

	// 三维数组
	a := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}
	fmt.Printf("三维数组 a = %v\n", a)

	// 访问第一层（二维数组）
	layer1 := a[0]
	fmt.Printf("\na[0] = %v\n", layer1)
	fmt.Printf("layer1 的类型: %T\n", layer1)

	// 访问第二层（一维数组）
	layer2 := a[0][1]
	fmt.Printf("\na[0][1] = %v\n", layer2)
	fmt.Printf("layer2 的类型: %T\n", layer2)

	// 访问具体元素
	element := a[0][1][1]
	fmt.Printf("\na[0][1][1] = %d\n", element)
	fmt.Printf("element 的类型: %T\n", element)

	// 访问不同位置的元素
	fmt.Println("\n--- 访问不同位置的元素 ---")
	fmt.Printf("a[0][0][0] = %d\n", a[0][0][0])
	fmt.Printf("a[1][0][1] = %d\n", a[1][0][1])
	fmt.Printf("a[2][1][1] = %d\n", a[2][1][1])
	fmt.Println()
}

// demonstrateMultidimensionalArrayTraversal 演示多维数组的遍历
func demonstrateMultidimensionalArrayTraversal() {
	fmt.Println("=== 3. 多维数组的遍历 ===")

	// 三维数组
	a := [3][2][2]int{
		{{0, 1}, {2, 3}},
		{{4, 5}, {6, 7}},
		{{8, 9}, {10, 11}},
	}

	// 方式1：使用 range 嵌套遍历
	fmt.Println("--- 方式1：使用 range 嵌套遍历 ---")
	for i, v := range a {
		fmt.Printf("第一层 index = %d, value = %v\n", i, v)
		for j, inner := range v {
			fmt.Printf("  第二层 index = %d, value = %v\n", j, inner)
			for k, element := range inner {
				fmt.Printf("    第三层 index = %d, value = %d\n", k, element)
			}
		}
	}

	// 方式2：使用传统 for 循环遍历
	fmt.Println("\n--- 方式2：使用传统 for 循环遍历 ---")
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			for k := 0; k < len(a[i][j]); k++ {
				fmt.Printf("a[%d][%d][%d] = %d\n", i, j, k, a[i][j][k])
			}
		}
	}
	fmt.Println()
}

// demonstrateTwoDimensionalArray 演示二维数组的详细示例
func demonstrateTwoDimensionalArray() {
	fmt.Println("=== 4. 二维数组详细示例 ===")

	// 创建一个 3x4 的二维数组（矩阵）
	matrix := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println("矩阵 matrix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}

	// 访问矩阵的特定元素
	fmt.Printf("\nmatrix[1][2] = %d\n", matrix[1][2])

	// 访问整行
	fmt.Printf("matrix[0] (第一行) = %v\n", matrix[0])

	// 修改元素
	matrix[1][2] = 100
	fmt.Printf("\n修改 matrix[1][2] = 100 后:\n")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// demonstrateMultidimensionalArrayInitialization 演示多维数组的不同初始化方式
func demonstrateMultidimensionalArrayInitialization() {
	fmt.Println("=== 5. 多维数组的不同初始化方式 ===")

	// 方式1：完整初始化
	fmt.Println("--- 方式1：完整初始化 ---")
	arr1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("arr1 = %v\n", arr1)

	// 方式2：部分初始化（未初始化的元素为零值）
	fmt.Println("\n--- 方式2：部分初始化 ---")
	arr2 := [2][3]int{
		{1, 2}, // 第三个元素为 0
		{4},    // 第二、三个元素为 0
	}
	fmt.Printf("arr2 = %v\n", arr2)

	// 方式3：使用索引初始化特定元素
	fmt.Println("\n--- 方式3：使用索引初始化特定元素 ---")
	arr3 := [3][3]int{
		0: {0: 1, 2: 3}, // 第一行：索引0和2有值
		2: {1: 5},       // 第三行：索引1有值
	}
	fmt.Printf("arr3 = %v\n", arr3)

	// 方式4：先声明后赋值
	fmt.Println("\n--- 方式4：先声明后赋值 ---")
	arr4 := [2][2]int{}
	arr4[0][0] = 10
	arr4[0][1] = 20
	arr4[1][0] = 30
	arr4[1][1] = 40
	fmt.Printf("arr4 = %v\n", arr4)
	fmt.Println()
}

// demonstrateMultidimensionalArrayLimitations 演示多维数组的限制和注意事项
func demonstrateMultidimensionalArrayLimitations() {
	fmt.Println("=== 6. 多维数组的限制和注意事项 ===")

	// 1. 数组长度是类型的一部分
	var arr1 [2][3]int
	var arr2 [3][2]int
	fmt.Printf("arr1 的类型: %T\n", arr1)
	fmt.Printf("arr2 的类型: %T\n", arr2)
	fmt.Println("说明：[2][3]int 和 [3][2]int 是不同的类型")

	// 2. 多维数组是值类型
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	arr4 := arr3 // 复制整个多维数组
	arr4[0][0] = 100
	fmt.Printf("\narr3 = %v\n", arr3)
	fmt.Printf("arr4 = %v\n", arr4)
	fmt.Println("说明：修改 arr4 不会影响 arr3，因为数组是值类型")

	// 3. 内层数组长度必须一致
	// var arr5 [2][3]int = [2][3]int{{1, 2}, {3, 4, 5}}  // 编译错误
	fmt.Println("\n⚠️ 注意：内层数组长度必须一致")
	fmt.Println("   例如：[2][3]int 要求每个内层数组都是长度为 3 的数组")
	fmt.Println()
}

// MultidimensionalArrayDemo 多维数组完整演示
func MultidimensionalArrayDemo() {
	fmt.Println("========== 1.12.3 多维数组 ==========")
	fmt.Println()
	fmt.Println("多维数组：数组的数组")
	fmt.Println("Go 中没有限制多维数组的嵌套层数")
	fmt.Println()
	fmt.Println("声明方式：var <array name> [<length1>][<length2>]... <type>")
	fmt.Println("访问方式：与访问普通数组的方式一致，使用多个下标")
	fmt.Println()

	demonstrateMultidimensionalArrayDeclaration()
	demonstrateMultidimensionalArrayAccess()
	demonstrateMultidimensionalArrayTraversal()
	demonstrateTwoDimensionalArray()
	demonstrateMultidimensionalArrayInitialization()
	demonstrateMultidimensionalArrayLimitations()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 多维数组是数组的数组")
	fmt.Println("✅ 可以使用多个下标访问不同层级的元素")
	fmt.Println("✅ 遍历多维数组需要使用嵌套循环")
	fmt.Println("✅ 多维数组是值类型，赋值会复制整个数组")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 数组长度是类型的一部分")
	fmt.Println("   - 内层数组长度必须一致")
	fmt.Println("   - 未初始化的元素为类型零值")
	fmt.Println()
}
