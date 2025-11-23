package slice

import "fmt"

// ========== 1.13.1 声明与初始化切片 ==========
//
// 切片的声明方式与声明数组的方式非常相似，与数组相比，切片不用声明长度
// 切片是引用类型，底层是对数组的引用
//
// 重要概念：
// 切片(Slice)并不是数组或者数组指针，而是数组的一个引用
// 切片本身是一个标准库中实现的一个特殊的结构体，这个结构体中有三个属性：
//   - array: 指向底层数组的指针
//   - len:   切片的长度（当前元素个数）
//   - cap:   切片的容量（底层数组从切片起始位置到末尾的元素数量）
//
// 具体可以查看 golang 源码仓库中 src/runtime/slice.go 文件：
//   type slice struct {
//       array unsafe.Pointer  // 指向底层数组的指针
//       len   int             // 切片的长度
//       cap   int             // 切片的容量
//   }
//
// 声明方式：var <slice name> []<type>
// 初始化方式：多种方式，包括字面量、make()、从数组创建等

// demonstrateSliceDeclaration 演示切片的声明方式
func demonstrateSliceDeclaration() {
	fmt.Println("=== 1. 切片的声明方式 ===")

	// 方式1：声明并初始化一个空的切片
	var s1 []int = []int{}
	fmt.Printf("方式1: var s1 []int = []int{}\n")
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

	// 方式2：类型推导，并初始化一个空的切片
	var s2 = []int{}
	fmt.Printf("\n方式2: var s2 = []int{}\n")
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	// 方式3：与方式2等价
	s3 := []int{}
	fmt.Printf("\n方式3: s3 := []int{}\n")
	fmt.Printf("s3 = %v, len = %d, cap = %d\n", s3, len(s3), cap(s3))

	// 方式4：与方式1、2、3等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}
	fmt.Printf("\n方式4: s4 := []int{1, 2, 3, 4}\n")
	fmt.Printf("s4 = %v, len = %d, cap = %d\n", s4, len(s4), cap(s4))

	fmt.Println()
}

// demonstrateSliceWithMake 演示使用 make() 函数创建切片
func demonstrateSliceWithMake() {
	fmt.Println("=== 2. 使用 make() 函数创建切片 ===")

	// 方式5：用 make() 函数创建切片，创建 []int 类型的切片，指定切片初始长度为 0
	s5 := make([]int, 0)
	fmt.Printf("方式5: s5 := make([]int, 0)\n")
	fmt.Printf("s5 = %v, len = %d, cap = %d\n", s5, len(s5), cap(s5))

	// 方式6：用 make() 函数创建切片，创建 []int 类型的切片，指定切片初始长度为 2，指定容量参数 4
	s6 := make([]int, 2, 4)
	fmt.Printf("\n方式6: s6 := make([]int, 2, 4)\n")
	fmt.Printf("s6 = %v, len = %d, cap = %d\n", s6, len(s6), cap(s6))
	fmt.Println("说明：len = 2 表示切片长度为 2，cap = 4 表示底层数组容量为 4")

	// make() 的其他用法
	s6a := make([]int, 3) // 只指定长度，容量等于长度
	fmt.Printf("\ns6a := make([]int, 3)\n")
	fmt.Printf("s6a = %v, len = %d, cap = %d\n", s6a, len(s6a), cap(s6a))
	fmt.Println("说明：只指定长度时，容量等于长度")

	s6b := make([]string, 2, 5)
	fmt.Printf("\ns6b := make([]string, 2, 5)\n")
	fmt.Printf("s6b = %v, len = %d, cap = %d\n", s6b, len(s6b), cap(s6b))
	fmt.Println()
}

// demonstrateSliceFromArray 演示从数组创建切片
func demonstrateSliceFromArray() {
	fmt.Println("=== 3. 从数组创建切片 ===")

	// 引用一个数组，初始化切片
	a := [5]int{6, 5, 4, 3, 2}
	fmt.Printf("原始数组 a = %v\n", a)

	// 从数组下标 2 开始，直到数组的最后一个元素
	s7 := a[2:]
	fmt.Printf("\ns7 := a[2:]\n")
	fmt.Printf("s7 = %v, len = %d, cap = %d\n", s7, len(s7), cap(s7))
	fmt.Println("说明：从索引 2 开始到数组末尾")

	// 从数组下标 1 开始，直到数组下标 3 的元素，创建一个新的切片
	s8 := a[1:3]
	fmt.Printf("\ns8 := a[1:3]\n")
	fmt.Printf("s8 = %v, len = %d, cap = %d\n", s8, len(s8), cap(s8))
	fmt.Println("说明：从索引 1 开始到索引 3（不包含 3）")

	// 从 0 到下标 2 的元素，创建一个新的切片
	s9 := a[:2]
	fmt.Printf("\ns9 := a[:2]\n")
	fmt.Printf("s9 = %v, len = %d, cap = %d\n", s9, len(s9), cap(s9))
	fmt.Println("说明：从索引 0 开始到索引 2（不包含 2）")

	// 完整切片：包含所有元素
	s10 := a[:]
	fmt.Printf("\ns10 := a[:]\n")
	fmt.Printf("s10 = %v, len = %d, cap = %d\n", s10, len(s10), cap(s10))
	fmt.Println("说明：包含数组的所有元素")
	fmt.Println()
}

// demonstrateSliceSharingUnderlyingArray 演示切片共享底层数组
func demonstrateSliceSharingUnderlyingArray() {
	fmt.Println("=== 4. 切片共享底层数组 ===")
	fmt.Println("说明：当切片是基于同一个数组创建出来时，修改数组中的值，同样会影响到这些切片")

	a := [5]int{6, 5, 4, 3, 2}
	fmt.Printf("原始数组 a = %v\n", a)

	// 从数组下标 2 开始，直到数组的最后一个元素
	s7 := a[2:]
	// 从数组下标 1 开始，直到数组下标 3 的元素，创建一个新的切片
	s8 := a[1:3]
	// 从 0 到下标 2 的元素，创建一个新的切片
	s9 := a[:2]

	fmt.Printf("\n创建切片后:\n")
	fmt.Printf("s7 = %v\n", s7)
	fmt.Printf("s8 = %v\n", s8)
	fmt.Printf("s9 = %v\n", s9)

	// 修改数组的值
	fmt.Println("\n修改数组 a 的值:")
	a[0] = 9
	a[1] = 8
	a[2] = 7

	fmt.Printf("修改后数组 a = %v\n", a)
	fmt.Printf("\n修改数组后，切片的值也改变了:\n")
	fmt.Printf("s7 = %v (共享 a[2:5])\n", s7)
	fmt.Printf("s8 = %v (共享 a[1:3])\n", s8)
	fmt.Printf("s9 = %v (共享 a[0:2])\n", s9)

	fmt.Println("\n⚠️ 注意：切片是引用类型，多个切片可以共享同一个底层数组")
	fmt.Println("   修改底层数组会影响所有基于该数组创建的切片")
	fmt.Println()
}

// demonstrateSliceLengthAndCapacity 演示切片的长度和容量
func demonstrateSliceLengthAndCapacity() {
	fmt.Println("=== 5. 切片的长度和容量 ===")

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("原始数组 arr = %v\n", arr)

	// 创建不同范围的切片
	s1 := arr[2:5]
	fmt.Printf("\ns1 := arr[2:5]\n")
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))
	fmt.Println("说明：长度 = 3 (5-2)，容量 = 8 (从索引2到数组末尾)")

	s2 := arr[:5]
	fmt.Printf("\ns2 := arr[:5]\n")
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("len(s2) = %d, cap(s2) = %d\n", len(s2), cap(s2))
	fmt.Println("说明：长度 = 5，容量 = 10 (从索引0到数组末尾)")

	s3 := arr[5:]
	fmt.Printf("\ns3 := arr[5:]\n")
	fmt.Printf("s3 = %v\n", s3)
	fmt.Printf("len(s3) = %d, cap(s3) = %d\n", len(s3), cap(s3))
	fmt.Println("说明：长度 = 5，容量 = 5 (从索引5到数组末尾)")

	fmt.Println("\n容量规则：")
	fmt.Println("  - 切片的容量 = 底层数组的长度 - 切片的起始索引")
	fmt.Println("  - 长度 = 切片的结束索引 - 切片的起始索引")
	fmt.Println()
}

// demonstrateSliceTypes 演示不同类型的切片
func demonstrateSliceTypes() {
	fmt.Println("=== 6. 不同类型的切片 ===")

	// 整数切片
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("整数切片: %v\n", intSlice)

	// 字符串切片
	strSlice := []string{"apple", "banana", "cherry"}
	fmt.Printf("字符串切片: %v\n", strSlice)

	// 布尔切片
	boolSlice := []bool{true, false, true}
	fmt.Printf("布尔切片: %v\n", boolSlice)

	// 空切片
	emptySlice := []int{}
	fmt.Printf("空切片: %v, len = %d, cap = %d\n", emptySlice, len(emptySlice), cap(emptySlice))

	// nil 切片
	var nilSlice []int
	fmt.Printf("nil 切片: %v, len = %d, cap = %d\n", nilSlice, len(nilSlice), cap(nilSlice))
	fmt.Printf("nil 切片是否为 nil: %v\n", nilSlice == nil)

	fmt.Println("\n说明：")
	fmt.Println("  - 空切片 []int{} 和 nil 切片 var s []int 的区别")
	fmt.Println("  - 空切片不是 nil，nil 切片是 nil")
	fmt.Println("  - 两者都可以使用，但 nil 切片更节省内存")
	fmt.Println()
}

// demonstrateSliceInternalStructure 演示切片的内部结构
func demonstrateSliceInternalStructure() {
	fmt.Println("=== 7. 切片的内部结构 ===")
	fmt.Println()
	fmt.Println("切片(Slice)并不是数组或者数组指针，而是数组的一个引用。")
	fmt.Println("切片本身是一个标准库中实现的一个特殊的结构体，")
	fmt.Println("这个结构体中有三个属性，分别代表数组指针、长度、容量。")
	fmt.Println()
	fmt.Println("具体可以查看 golang 源码仓库中 src/runtime/slice.go 文件：")
	fmt.Println()
	fmt.Println("  type slice struct {")
	fmt.Println("      array unsafe.Pointer  // 指向底层数组的指针")
	fmt.Println("      len   int             // 切片的长度（当前元素个数）")
	fmt.Println("      cap   int             // 切片的容量（底层数组从切片起始位置到末尾的元素数量）")
	fmt.Println("  }")
	fmt.Println()

	// 演示切片的三个属性
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:7]

	fmt.Println("示例：")
	fmt.Printf("  底层数组 arr = %v\n", arr)
	fmt.Printf("  切片 s := arr[2:7]\n")
	fmt.Printf("  切片 s = %v\n", s)
	fmt.Printf("  len(s) = %d (切片中元素的数量)\n", len(s))
	fmt.Printf("  cap(s) = %d (从索引2到数组末尾的元素数量)\n", cap(s))
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - array: 指向 arr[2] 的指针（切片的起始位置）")
	fmt.Println("  - len:   5 (索引2到6，共5个元素)")
	fmt.Println("  - cap:   8 (从索引2到数组末尾，共8个元素)")
	fmt.Println()
	fmt.Println("⚠️ 重要理解：")
	fmt.Println("  - 切片本身不存储数据，只存储指向底层数组的指针")
	fmt.Println("  - 多个切片可以共享同一个底层数组")
	fmt.Println("  - 修改切片元素实际上修改的是底层数组的元素")
	fmt.Println()
}

// demonstrateSliceVsArray 演示切片和数组的区别
func demonstrateSliceVsArray() {
	fmt.Println("=== 8. 切片和数组的区别 ===")
	fmt.Println()

	// 1. 声明语法
	fmt.Println("--- 1. 声明语法 ---")
	fmt.Println("数组：必须指定长度")
	fmt.Println("  var arr [5]int        // ✅ 正确：指定长度为 5")
	fmt.Println("  var arr []int         // ❌ 错误：数组必须指定长度")
	fmt.Println()
	fmt.Println("切片：不需要指定长度")
	fmt.Println("  var s []int           // ✅ 正确：切片不需要指定长度")
	fmt.Println("  var s [5]int          // ❌ 错误：这是数组，不是切片")
	fmt.Println()

	// 2. 类型特性
	fmt.Println("--- 2. 类型特性 ---")
	var arr1 [5]int
	var arr2 [10]int
	var s1 []int
	var s2 []int

	fmt.Printf("arr1 的类型: %T\n", arr1)
	fmt.Printf("arr2 的类型: %T\n", arr2)
	fmt.Printf("s1 的类型: %T\n", s1)
	fmt.Printf("s2 的类型: %T\n", s2)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - 数组：[5]int 和 [10]int 是不同的类型（长度是类型的一部分）")
	fmt.Println("  - 切片：[]int 就是 []int（长度不是类型的一部分）")
	fmt.Println()

	// 3. 内存存储
	fmt.Println("--- 3. 内存存储 ---")
	arr := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("数组 arr = %v\n", arr)
	fmt.Printf("切片 s = %v\n", s)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - 数组：直接存储数据，数组变量就是数据本身")
	fmt.Println("  - 切片：存储切片头（指针+长度+容量），数据在底层数组中")
	fmt.Println()

	// 4. 参数传递
	fmt.Println("--- 4. 参数传递 ---")
	arr3 := [3]int{1, 2, 3}
	s3 := []int{1, 2, 3}

	fmt.Printf("传递前 arr3 = %v\n", arr3)
	fmt.Printf("传递前 s3 = %v\n", s3)

	modifyArray(arr3)
	modifySlice(s3)

	fmt.Printf("传递后 arr3 = %v (数组未改变，因为是值传递)\n", arr3)
	fmt.Printf("传递后 s3 = %v (切片已改变，因为是引用传递)\n", s3)
	fmt.Println()

	// 5. 长度和容量
	fmt.Println("--- 5. 长度和容量 ---")
	arr4 := [5]int{1, 2, 3, 4, 5}
	s4 := []int{1, 2, 3, 4, 5}

	fmt.Printf("数组 arr4 = %v: len = %d, cap = %d (数组只有长度)\n", arr4, len(arr4), cap(arr4))
	fmt.Printf("切片 s4 = %v: len = %d, cap = %d (切片有长度和容量)\n", s4, len(s4), cap(s4))
	fmt.Println()

	// 6. 长度可变性
	fmt.Println("--- 6. 长度可变性 ---")
	fmt.Println("数组：长度固定，不能改变")
	arr5 := [3]int{1, 2, 3}
	fmt.Printf("arr5 = %v, len = %d\n", arr5, len(arr5))
	fmt.Println("  // arr5 = append(arr5, 4)  // ❌ 错误：数组不支持 append")

	fmt.Println("\n切片：长度可变，可以动态增长")
	s5 := []int{1, 2, 3}
	fmt.Printf("s5 = %v, len = %d, cap = %d\n", s5, len(s5), cap(s5))
	s5 = append(s5, 4, 5)
	fmt.Printf("append 后 s5 = %v, len = %d, cap = %d\n", s5, len(s5), cap(s5))
	fmt.Println()

	// 7. 对比总结
	fmt.Println("--- 7. 快速对比表 ---")
	fmt.Println("┌─────────────┬──────────┬──────────┐")
	fmt.Println("│   特性      │   数组   │   切片   │")
	fmt.Println("├─────────────┼──────────┼──────────┤")
	fmt.Println("│ 声明        │ [5]int   │ []int    │")
	fmt.Println("│ 类型        │ 值类型   │ 引用类型 │")
	fmt.Println("│ 长度        │ 固定     │ 可变     │")
	fmt.Println("│ 长度是类型  │ ✅ 是    │ ❌ 否    │")
	fmt.Println("│ 内存        │ 直接存储 │ 切片头   │")
	fmt.Println("│ 传递        │ 值传递   │ 引用传递 │")
	fmt.Println("│ 容量        │ = 长度   │ ≥ 长度   │")
	fmt.Println("└─────────────┴──────────┴──────────┘")
	fmt.Println()

	fmt.Println("⚠️ 重要提示：")
	fmt.Println("  - 数组：固定大小、值语义，适合固定大小的数据集合")
	fmt.Println("  - 切片：动态大小、引用语义，适合大多数场景（推荐）")
	fmt.Println()
}

// modifyArray 修改数组（值传递，不影响原数组）
func modifyArray(arr [3]int) {
	arr[0] = 999
	fmt.Printf("  函数内修改数组: arr = %v\n", arr)
}

// modifySlice 修改切片（引用传递，影响原切片）
func modifySlice(s []int) {
	s[0] = 999
	fmt.Printf("  函数内修改切片: s = %v\n", s)
}

// SliceDeclarationDemo 切片声明与初始化完整演示
func SliceDeclarationDemo() {
	fmt.Println("========== 1.13.1 声明与初始化切片 ==========")
	fmt.Println()
	fmt.Println("切片的声明方式与声明数组的方式非常相似，")
	fmt.Println("与数组相比，切片不用声明长度。")
	fmt.Println()
	fmt.Println("切片(Slice)并不是数组或者数组指针，而是数组的一个引用。")
	fmt.Println("切片本身是一个标准库中实现的一个特殊的结构体，")
	fmt.Println("这个结构体中有三个属性：")
	fmt.Println("  - array: 指向底层数组的指针")
	fmt.Println("  - len:   切片的长度（当前元素个数）")
	fmt.Println("  - cap:   切片的容量（底层数组从切片起始位置到末尾的元素数量）")
	fmt.Println()
	fmt.Println("切片是引用类型，底层是对数组的引用。")
	fmt.Println("当切片是基于同一个数组创建出来时，修改数组中的值，")
	fmt.Println("同样会影响到这些切片。")
	fmt.Println()

	demonstrateSliceDeclaration()
	demonstrateSliceWithMake()
	demonstrateSliceFromArray()
	demonstrateSliceSharingUnderlyingArray()
	demonstrateSliceLengthAndCapacity()
	demonstrateSliceTypes()
	demonstrateSliceInternalStructure()
	demonstrateSliceVsArray()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 切片声明：var <slice name> []<type>")
	fmt.Println("✅ 初始化方式：")
	fmt.Println("   - 字面量：[]int{1, 2, 3}")
	fmt.Println("   - make()：make([]int, len, cap)")
	fmt.Println("   - 从数组创建：arr[start:end]")
	fmt.Println()
	fmt.Println("✅ 切片的内部结构（src/runtime/slice.go）：")
	fmt.Println("   type slice struct {")
	fmt.Println("       array unsafe.Pointer  // 指向底层数组的指针")
	fmt.Println("       len   int             // 切片的长度")
	fmt.Println("       cap   int             // 切片的容量")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("✅ 切片是引用类型，共享底层数组")
	fmt.Println("✅ 长度（len）：切片中元素的数量")
	fmt.Println("✅ 容量（cap）：底层数组从切片起始位置到末尾的元素数量")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 切片本身不存储数据，只存储指向底层数组的指针")
	fmt.Println("   - 修改底层数组会影响所有基于该数组创建的切片")
	fmt.Println("   - 多个切片可以共享同一个底层数组")
	fmt.Println("   - 空切片 []int{} 不是 nil")
	fmt.Println("   - nil 切片 var s []int 是 nil")
	fmt.Println()
}
