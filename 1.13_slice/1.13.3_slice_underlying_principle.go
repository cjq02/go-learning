package slice

import "fmt"

// ========== 1.13.3 切片底层原理 ==========
//
// 切片类型实际上是比较特殊的指针类型，当声明一个切片类型时，就是声明了一个指针
// 这个指针指向的切片结构体，切片结构体中记录的三个属性：数组指针、长度、容量
// 这几个属性在创建一个切片时就定义好，并且在之后都不能再被修改
//
// 重要概念：
// 1. 切片本身是一个特殊的指针
// 2. 不使用 append() 时，多个切片可以共享同一个底层数组
// 3. append() 不触发扩容时，新切片和原切片共享底层数组
// 4. append() 触发扩容时，会创建新的底层数组

// demonstrateSliceWithoutAppend 演示不使用 append() 时的切片行为
func demonstrateSliceWithoutAppend() {
	fmt.Println("=== 1. 不使用 append() 时的切片行为 ===")
	fmt.Println()
	fmt.Println("在不使用 append() 函数的情况下，")
	fmt.Println("在函数内部对切片的修改，都会影响到原始实例。")
	fmt.Println()

	s := make([]int, 3, 6)
	fmt.Printf("s length: %d\n", len(s))
	fmt.Printf("s capacity: %d\n", cap(s))
	fmt.Printf("initial, s = %v\n", s)

	s[1] = 2
	fmt.Printf("set position 1, s = %v\n", s)

	modifySliceInPrinciple(s)
	fmt.Printf("after modifySlice, s = %v\n", s)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - 切片是引用类型，传递的是切片头（指针+长度+容量）")
	fmt.Println("  - 函数内修改切片元素，会影响原切片")
	fmt.Println("  - 因为它们共享同一个底层数组")
	fmt.Println()
}

// modifySliceInPrinciple 修改切片元素（用于底层原理演示）
func modifySliceInPrinciple(param []int) {
	fmt.Printf("  在 modifySlice 函数中，修改前 param = %v\n", param)
	param[0] = 1024
	fmt.Printf("  在 modifySlice 函数中，修改后 param = %v\n", param)
}

// demonstrateAppendWithoutExpansion 演示 append() 不触发扩容时的行为
func demonstrateAppendWithoutExpansion() {
	fmt.Println("=== 2. append() 不触发扩容时的行为 ===")
	fmt.Println()
	fmt.Println("当没有触发切片扩容时：")
	fmt.Println("  - 原来的切片引用，长度和容量不变")
	fmt.Println("  - 新追加的值超过切片可访问范围，访问不到新追加的值")
	fmt.Println("  - 新的切片引用，长度加一，容量不变，可以访问到新追加的值")
	fmt.Println("  - 两个切片共享同一个底层数组")
	fmt.Println()

	s := make([]int, 3, 6)
	fmt.Printf("initial, s = %v\n", s)
	s[1] = 2
	fmt.Printf("after set position 1, s = %v\n", s)

	s2 := append(s, 4)
	fmt.Printf("after append, s2 length: %d\n", len(s2))
	fmt.Printf("after append, s2 capacity: %d\n", cap(s2))
	fmt.Printf("after append, s = %v (长度不变，访问不到新元素)\n", s)
	fmt.Printf("after append, s2 = %v (长度+1，可以访问新元素)\n", s2)

	s[0] = 1024
	fmt.Printf("\nafter set position 0, s = %v\n", s)
	fmt.Printf("after set position 0, s2 = %v (共享底层数组，s2[0]也改变了)\n", s2)

	appendInFunc(s)
	fmt.Printf("after append in func, s = %v\n", s)
	fmt.Printf("after append in func, s2 = %v\n", s2)
	fmt.Println()
	fmt.Println("⚠️ 重要理解：")
	fmt.Println("  - s 和 s2 共享同一个底层数组")
	fmt.Println("  - 修改 s[0] 会影响 s2[0]")
	fmt.Println("  - 但 s 的长度是3，访问不到 s2 追加的元素")
	fmt.Println()
}

// appendInFunc 在函数内使用 append
func appendInFunc(param []int) {
	fmt.Printf("  在 appendInFunc 函数中，修改前 param = %v\n", param)
	param = append(param, 1022)
	fmt.Printf("  in func, param = %v (append后创建了新切片引用)\n", param)
	param[2] = 512
	fmt.Printf("  set position 2 in func, param = %v\n", param)
	fmt.Println("  说明：param 是新的切片引用，但共享底层数组")
}

// demonstrateAppendWithExpansion 演示 append() 触发扩容时的行为
func demonstrateAppendWithExpansion() {
	fmt.Println("=== 3. append() 触发扩容时的行为 ===")
	fmt.Println()
	fmt.Println("当 append() 函数触发扩容后：")
	fmt.Println("  - 实际上是新创建了一个数组实例")
	fmt.Println("  - 把原来的数组中的数据复制到了新数组中")
	fmt.Println("  - 然后创建一个新的切片实例并返回")
	fmt.Println("  - 这时原始切片和新切片指向不同的数组，修改不会相互影响")
	fmt.Println()

	s := make([]int, 2)
	fmt.Printf("initial, s = %v (len=%d, cap=%d)\n", s, len(s), cap(s))

	s2 := append(s, 4)
	fmt.Printf("\nafter append, s length: %d\n", len(s))
	fmt.Printf("after append, s capacity: %d\n", cap(s))
	fmt.Printf("after append, s2 length: %d\n", len(s2))
	fmt.Printf("after append, s2 capacity: %d (容量翻倍，触发了扩容)\n", cap(s2))
	fmt.Printf("after append, s = %v\n", s)
	fmt.Printf("after append, s2 = %v\n", s2)

	s[0] = 1024
	fmt.Printf("\nafter set position 0, s = %v\n", s)
	fmt.Printf("after set position 0, s2 = %v (s2未改变，因为指向不同的数组)\n", s2)

	appendInFuncWithExpansion(s2)
	fmt.Printf("after append in func, s2 = %v (s2未改变)\n", s2)
	fmt.Println()
	fmt.Println("⚠️ 重要理解：")
	fmt.Println("  - 扩容后，s 和 s2 指向不同的底层数组")
	fmt.Println("  - 修改 s[0] 不会影响 s2[0]")
	fmt.Println("  - 这是切片扩容的关键特性")
	fmt.Println()
}

// appendInFuncWithExpansion 在函数内使用 append（会触发扩容）
func appendInFuncWithExpansion(param []int) {
	fmt.Printf("  在 appendInFuncWithExpansion 函数中，修改前 param = %v\n", param)
	param1 := append(param, 511)
	fmt.Printf("  in func, param1 = %v (append后，可能触发扩容)\n", param1)
	param2 := append(param1, 512)
	fmt.Printf("  in func, param2 = %v\n", param2)
	param2[2] = 500
	fmt.Printf("  set position 2 in func, param2 = %v\n", param2)
	fmt.Println("  说明：param1 和 param2 可能指向新的数组（如果触发了扩容）")
}

// demonstrateSlicePointerNature 演示切片的指针特性
func demonstrateSlicePointerNature() {
	fmt.Println("=== 4. 切片的指针特性 ===")
	fmt.Println()
	fmt.Println("切片类型实际上是比较特殊的指针类型")
	fmt.Println("当声明一个切片类型时，就是声明了一个指针")
	fmt.Println("这个指针指向的切片结构体包含：")
	fmt.Println("  - array: 指向底层数组的指针")
	fmt.Println("  - len:   切片的长度")
	fmt.Println("  - cap:   切片的容量")
	fmt.Println()

	s1 := make([]int, 3, 6)
	s2 := s1 // 复制切片头（指针+长度+容量）

	fmt.Printf("s1 = %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s1[0] = 100
	fmt.Printf("\n修改 s1[0] = 100 后:\n")
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v (s2也改变了，因为共享底层数组)\n", s2)

	s2 = append(s2, 999)
	fmt.Printf("\nappend(s2, 999) 后:\n")
	fmt.Printf("s1 = %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Println("说明：s2 追加元素后，长度增加，但 s1 的长度不变")
	fmt.Println()
}

// demonstrateSliceExpansionRule 演示切片扩容规则
func demonstrateSliceExpansionRule() {
	fmt.Println("=== 5. 切片扩容规则 ===")
	fmt.Println()

	s := make([]int, 0, 2)
	fmt.Printf("初始: len=%d, cap=%d\n", len(s), cap(s))

	for i := 0; i < 10; i++ {
		oldCap := cap(s)
		s = append(s, i)
		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("append %d: len=%d, cap=%d -> %d (扩容了!)\n", i, len(s), oldCap, newCap)
		} else {
			fmt.Printf("append %d: len=%d, cap=%d (未扩容)\n", i, len(s), cap(s))
		}
	}

	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - Go 切片的扩容策略：当容量不足时，会创建新的底层数组")
	fmt.Println("  - 新容量通常是原容量的 2 倍（当容量 < 1024 时）")
	fmt.Println("  - 扩容后，原切片和新切片指向不同的数组")
	fmt.Println()
}

// demonstrateSliceSharingArray 演示切片共享底层数组
func demonstrateSliceSharingArray() {
	fmt.Println("=== 6. 切片共享底层数组 ===")
	fmt.Println()

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("原始数组 arr = %v\n", arr)

	s1 := arr[2:5] // [2, 3, 4]
	s2 := arr[3:6] // [3, 4, 5]

	fmt.Printf("s1 = arr[2:5] = %v\n", s1)
	fmt.Printf("s2 = arr[3:6] = %v\n", s2)

	// 修改数组
	arr[3] = 100
	fmt.Printf("\n修改 arr[3] = 100 后:\n")
	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("s1 = %v (s1[1]也改变了)\n", s1)
	fmt.Printf("s2 = %v (s2[0]也改变了)\n", s2)

	// 修改切片
	s1[0] = 200
	fmt.Printf("\n修改 s1[0] = 200 后:\n")
	fmt.Printf("arr = %v (arr[2]也改变了)\n", arr)
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v (s2未直接改变，但底层数组变了)\n", s2)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - s1 和 s2 共享同一个底层数组 arr")
	fmt.Println("  - 修改数组会影响所有基于该数组的切片")
	fmt.Println("  - 修改切片也会影响底层数组和其他切片")
	fmt.Println()
}

// SliceUnderlyingPrincipleDemo 切片底层原理完整演示
func SliceUnderlyingPrincipleDemo() {
	fmt.Println("========== 1.13.3 切片底层原理 ==========")
	fmt.Println()
	fmt.Println("切片类型实际上是比较特殊的指针类型，")
	fmt.Println("当声明一个切片类型时，就是声明了一个指针。")
	fmt.Println()
	fmt.Println("这个指针指向的切片结构体，切片结构体中记录的三个属性：")
	fmt.Println("  - array: 指向底层数组的指针")
	fmt.Println("  - len:   切片的长度")
	fmt.Println("  - cap:   切片的容量")
	fmt.Println()
	fmt.Println("这几个属性在创建一个切片时就定义好，并且在之后都不能再被修改。")
	fmt.Println()

	demonstrateSliceWithoutAppend()
	demonstrateAppendWithoutExpansion()
	demonstrateAppendWithExpansion()
	demonstrateSlicePointerNature()
	demonstrateSliceExpansionRule()
	demonstrateSliceSharingArray()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 切片是指针类型，传递的是切片头（指针+长度+容量）")
	fmt.Println("✅ 不使用 append() 时，多个切片共享同一个底层数组")
	fmt.Println("✅ append() 不触发扩容时，新切片和原切片共享底层数组")
	fmt.Println("✅ append() 触发扩容时，会创建新的底层数组，原切片和新切片分离")
	fmt.Println()
	fmt.Println("⚠️ 关键理解：")
	fmt.Println("   - 切片触发扩容前，切片一直共用相同的数组")
	fmt.Println("   - 切片触发扩容后，会创建新的数组，并复制这些数据")
	fmt.Println("   - 切片本身是一个特殊的指针，Go 针对切片类型添加了一些语法糖")
	fmt.Println()
}
