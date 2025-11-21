package slice

import "fmt"

// ========== 1.13.2 使用切片 ==========

// demonstrateSliceAccess 演示访问切片
func demonstrateSliceAccess() {
	fmt.Println("=== 1.13.2.1 访问切片 ===")

	s1 := []int{5, 4, 3, 2, 1}
	fmt.Printf("原始切片 s1 = %v\n", s1)

	// 下标访问切片
	e1 := s1[0]
	e2 := s1[1]
	e3 := s1[2]
	fmt.Printf("\n下标访问:\n")
	fmt.Printf("e1 = s1[0] = %d\n", e1)
	fmt.Printf("e2 = s1[1] = %d\n", e2)
	fmt.Printf("e3 = s1[2] = %d\n", e3)

	// 向指定位置赋值
	fmt.Println("\n修改切片元素:")
	s1[0] = 10
	s1[1] = 9
	s1[2] = 8
	fmt.Printf("修改后 s1 = %v\n", s1)

	// range 迭代访问切片
	fmt.Println("\n使用 range 迭代访问切片:")
	for i, v := range s1 {
		fmt.Printf("s1[%d] = %d\n", i, v)
	}
	fmt.Println()
}

// demonstrateSliceLengthAndCapacityUsage 演示切片的长度和容量（使用场景）
func demonstrateSliceLengthAndCapacityUsage() {
	fmt.Println("=== 切片的长度和容量 ===")
	fmt.Println()
	fmt.Println("长度（len）：表示切片可以访问到底层数组的数据范围")
	fmt.Println("容量（cap）：表示切片引用的底层数组的长度")
	fmt.Println("当切片是 nil 时，len() 和 cap() 函数获取的值都是 0")
	fmt.Println("切片的长度小于等于切片的容量")
	fmt.Println()

	// nil 切片
	var nilSlice []int
	fmt.Printf("nilSlice = %v\n", nilSlice)
	fmt.Printf("nilSlice length: %d\n", len(nilSlice))
	fmt.Printf("nilSlice capacity: %d\n", cap(nilSlice))
	fmt.Printf("nilSlice == nil: %v\n", nilSlice == nil)
	fmt.Println()

	// 普通切片
	s2 := []int{9, 8, 7, 6, 5}
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s2 length: %d\n", len(s2))
	fmt.Printf("s2 capacity: %d\n", cap(s2))
	fmt.Println()

	// 从数组创建的切片
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := arr[2:7]
	fmt.Printf("数组 arr = %v\n", arr)
	fmt.Printf("切片 s3 = arr[2:7] = %v\n", s3)
	fmt.Printf("s3 length: %d (索引2到6，共5个元素)\n", len(s3))
	fmt.Printf("s3 capacity: %d (从索引2到数组末尾，共8个元素)\n", cap(s3))
	fmt.Println()
}

// demonstrateSliceAppend 演示切片添加元素
func demonstrateSliceAppend() {
	fmt.Println("=== 1.13.2.2 切片添加元素 ===")
	fmt.Println()
	fmt.Println("切片是变长的，可以向切片追加新的元素")
	fmt.Println("可以使用内置的 append() 向切片追加元素")
	fmt.Println()
	fmt.Println("append() 函数特性：")
	fmt.Println("  - 只有切片类型可以使用")
	fmt.Println("  - 第一个参数必须是切片类型")
	fmt.Println("  - 后面追加的元素参数是变长类型，一次可以追加多个元素")
	fmt.Println("  - 每次 append() 都会返回一个新的切片引用")
	fmt.Println()

	s3 := []int{}
	fmt.Printf("初始 s3 = %v\n", s3)

	// append 函数追加元素
	s3 = append(s3) // 不追加任何元素，无意义的操作
	fmt.Printf("append(s3) 后 s3 = %v\n", s3)

	s3 = append(s3, 1) // 追加一个元素
	fmt.Printf("append(s3, 1) 后 s3 = %v\n", s3)

	s3 = append(s3, 2, 3) // 追加多个元素
	fmt.Printf("append(s3, 2, 3) 后 s3 = %v\n", s3)

	s3 = append(s3, 4, 5, 6) // 追加更多元素
	fmt.Printf("append(s3, 4, 5, 6) 后 s3 = %v\n", s3)
	fmt.Printf("s3 length: %d, capacity: %d\n", len(s3), cap(s3))
	fmt.Println()
}

// demonstrateSliceInsertElement 演示向指定位置添加元素
func demonstrateSliceInsertElement() {
	fmt.Println("=== 向指定位置添加元素 ===")

	s4 := []int{1, 2, 4, 5}
	fmt.Printf("原始切片 s4 = %v\n", s4)
	fmt.Println("目标：在索引2的位置插入元素3")

	// 向指定位置添加元素：s4[:2] + [3] + s4[2:]
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	fmt.Printf("插入后 s4 = %v\n", s4)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - s4[:2] 获取索引0到1的元素 [1, 2]")
	fmt.Println("  - s4[2:] 是切片 [4, 5]")
	fmt.Println("  - s4[2:]... 是展开操作符，将切片 [4, 5] 展开为 4, 5 两个参数")
	fmt.Println("  - append([]int{3}, s4[2:]...) 等价于 append([]int{3}, 4, 5)")
	fmt.Println("  - 结果是 [3, 4, 5]")
	fmt.Println("  - 最后 append(s4[:2], [3, 4, 5]...) 将 [1, 2] 和 [3, 4, 5] 合并")
	fmt.Println()
	fmt.Println("⚠️ 重要：... 是展开操作符（spread operator）")
	fmt.Println("   - 在函数调用时，slice... 将切片展开为多个参数")
	fmt.Println("   - 例如：append(s, []int{1, 2}...) 等价于 append(s, 1, 2)")
	fmt.Println()

	// 另一个示例：在开头插入
	s5 := []int{2, 3, 4}
	fmt.Printf("原始切片 s5 = %v\n", s5)
	s5 = append([]int{1}, s5...) // s5... 将切片展开为 2, 3, 4
	fmt.Printf("在开头插入1后 s5 = %v\n", s5)
	fmt.Println("说明：append([]int{1}, s5...) 等价于 append([]int{1}, 2, 3, 4)")
	fmt.Println()

	// 在末尾插入（等同于 append）
	s6 := []int{1, 2, 3}
	fmt.Printf("原始切片 s6 = %v\n", s6)
	s6 = append(s6, 4)
	fmt.Printf("在末尾插入4后 s6 = %v\n", s6)
	fmt.Println()
}

// demonstrateSliceRemoveElement 演示移除指定位置的元素
func demonstrateSliceRemoveElement() {
	fmt.Println("=== 移除指定位置的元素 ===")

	s5 := []int{1, 2, 3, 5, 4}
	fmt.Printf("原始切片 s5 = %v\n", s5)
	fmt.Println("目标：移除索引3的元素（值为5）")

	// 移除指定位置元素：s5[:3] + s5[4:]
	s5 = append(s5[:3], s5[4:]...)
	fmt.Printf("移除后 s5 = %v\n", s5)
	fmt.Println()
	fmt.Println("说明：")
	fmt.Println("  - s5[:3] 获取索引0到2的元素 [1, 2, 3]")
	fmt.Println("  - s5[4:] 获取索引4及之后的元素 [4]")
	fmt.Println("  - 使用 ... 展开切片，然后合并")
	fmt.Println()

	// 移除第一个元素
	s6 := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始切片 s6 = %v\n", s6)
	s6 = s6[1:]
	fmt.Printf("移除第一个元素后 s6 = %v\n", s6)
	fmt.Println()

	// 移除最后一个元素
	s7 := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始切片 s7 = %v\n", s7)
	s7 = s7[:len(s7)-1]
	fmt.Printf("移除最后一个元素后 s7 = %v\n", s7)
	fmt.Println()
}

// demonstrateSliceCopy 演示复制切片
func demonstrateSliceCopy() {
	fmt.Println("=== 1.13.2.3 复制切片 ===")
	fmt.Println()
	fmt.Println("可以使用内置函数 copy() 把某个切片中的所有元素复制到另一个切片")
	fmt.Println("复制的长度是它们中最短的切片长度")
	fmt.Println()

	// 示例1：源切片长度小于目标切片
	fmt.Println("--- 示例1：源切片长度小于目标切片 ---")
	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	fmt.Printf("before copy, src1 = %v\n", src1)
	fmt.Printf("before copy, dst1 = %v\n", dst1)

	copied1 := copy(dst1, src1)
	fmt.Printf("copy(dst1, src1) 返回复制的元素数量: %d\n", copied1)

	fmt.Printf("after copy, src1 = %v\n", src1)
	fmt.Printf("after copy, dst1 = %v\n", dst1)
	fmt.Println("说明：只复制了3个元素（src1的长度），dst1的最后一个元素保持零值")
	fmt.Println()

	// 示例2：源切片长度大于目标切片
	fmt.Println("--- 示例2：源切片长度大于目标切片 ---")
	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3)

	fmt.Printf("before copy, src2 = %v\n", src2)
	fmt.Printf("before copy, dst2 = %v\n", dst2)

	copied2 := copy(dst2, src2)
	fmt.Printf("copy(dst2, src2) 返回复制的元素数量: %d\n", copied2)

	fmt.Printf("after copy, src2 = %v\n", src2)
	fmt.Printf("after copy, dst2 = %v\n", dst2)
	fmt.Println("说明：只复制了3个元素（dst2的长度），src2的后2个元素没有被复制")
	fmt.Println()

	// 示例3：相同长度的切片
	fmt.Println("--- 示例3：相同长度的切片 ---")
	src3 := []int{10, 20, 30}
	dst3 := make([]int, 3)

	fmt.Printf("before copy, src3 = %v\n", src3)
	fmt.Printf("before copy, dst3 = %v\n", dst3)

	copied3 := copy(dst3, src3)
	fmt.Printf("copy(dst3, src3) 返回复制的元素数量: %d\n", copied3)

	fmt.Printf("after copy, src3 = %v\n", src3)
	fmt.Printf("after copy, dst3 = %v\n", dst3)
	fmt.Println("说明：所有元素都被复制")
	fmt.Println()

	// 示例4：复制到已有数据的切片
	fmt.Println("--- 示例4：复制到已有数据的切片 ---")
	src4 := []int{100, 200}
	dst4 := []int{1, 2, 3, 4, 5}

	fmt.Printf("before copy, src4 = %v\n", src4)
	fmt.Printf("before copy, dst4 = %v\n", dst4)

	copied4 := copy(dst4, src4)
	fmt.Printf("copy(dst4, src4) 返回复制的元素数量: %d\n", copied4)

	fmt.Printf("after copy, src4 = %v\n", src4)
	fmt.Printf("after copy, dst4 = %v\n", dst4)
	fmt.Println("说明：只覆盖了前2个元素，后面的元素保持不变")
	fmt.Println()

	fmt.Println("⚠️ 注意事项：")
	fmt.Println("  - copy() 返回实际复制的元素数量")
	fmt.Println("  - 复制的长度是源切片和目标切片中较短的长度")
	fmt.Println("  - copy() 不会改变源切片")
	fmt.Println("  - 如果目标切片长度不够，只复制部分元素")
	fmt.Println()
}

// SliceUsageDemo 切片使用完整演示
func SliceUsageDemo() {
	fmt.Println("========== 1.13.2 使用切片 ==========")
	fmt.Println()

	demonstrateSliceAccess()
	demonstrateSliceLengthAndCapacityUsage()
	demonstrateSliceAppend()
	demonstrateSliceInsertElement()
	demonstrateSliceRemoveElement()
	demonstrateSliceCopy()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 访问切片：使用下标 s[i] 或 range 遍历")
	fmt.Println("✅ 长度和容量：len(s) 和 cap(s)")
	fmt.Println("✅ 添加元素：append(s, elements...)")
	fmt.Println("✅ 插入元素：append(s[:i], append([]T{value}, s[i:]...)...)")
	fmt.Println("✅ 移除元素：append(s[:i], s[i+1:]...)")
	fmt.Println("✅ 复制切片：copy(dst, src)")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - append() 返回新的切片引用，需要重新赋值")
	fmt.Println("   - copy() 复制的长度是较短切片的长度")
	fmt.Println("   - nil 切片的 len 和 cap 都是 0")
	fmt.Println("   - 切片的长度小于等于容量")
	fmt.Println()
}
