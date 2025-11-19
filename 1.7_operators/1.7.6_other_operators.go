package operators

import "fmt"

// ========== 1.7.6 其他运算符 ==========

// demonstrateAddressOperators 地址运算符演示
func demonstrateAddressOperators() {
	fmt.Println("=== 地址运算符 ===")

	a := 4
	fmt.Printf("变量 a 的值: %d\n", a)

	// & 取地址运算符
	var ptr *int
	ptr = &a
	fmt.Printf("指针 ptr 指向的值: %d\n", *ptr)

	// 修改指针指向的值
	*ptr = 10
	fmt.Printf("通过指针修改后，变量 a 的值: %d\n", a)
	fmt.Printf("指针 ptr 指向的值: %d\n", *ptr)

	// 显示地址
	fmt.Printf("变量 a 的地址: %p\n", &a)
	fmt.Printf("指针 ptr 的值（地址）: %p\n", ptr)
}

// demonstrateChannelOperators 通道运算符演示
func demonstrateChannelOperators() {
	fmt.Println("\n=== 通道运算符 ===")

	// 创建通道
	ch := make(chan int, 3)

	// <- 发送运算符
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Printf("向通道发送了 3 个值\n")

	// <- 接收运算符
	value1 := <-ch
	value2 := <-ch
	value3 := <-ch

	fmt.Printf("从通道接收的值: %d, %d, %d\n", value1, value2, value3)

	// 关闭通道
	close(ch)
	fmt.Println("通道已关闭")
}

// demonstratePointerArithmetic 指针运算演示（Go中不支持）
func demonstratePointerArithmetic() {
	fmt.Println("\n=== 指针运算（Go不支持）===")

	fmt.Println("Go 语言不支持指针运算，不像 C/C++ 那样可以进行指针的加减运算。")
	fmt.Println("这是为了安全性和简洁性考虑。")

	a := 10
	ptr := &a

	fmt.Printf("指针地址: %p\n", ptr)
	fmt.Printf("指针指向的值: %d\n", *ptr)

	// 以下代码会编译错误：
	// ptr++    // 错误：不支持指针运算
	// ptr + 1  // 错误：不支持指针运算

	fmt.Println("Go 通过 slice 和数组来提供类似的功能，同时保证内存安全。")
}

// demonstrateSliceOperators 切片运算符演示
func demonstrateSliceOperators() {
	fmt.Println("\n=== 切片运算符 ===")

	// 创建切片
	slice := []int{10, 20, 30, 40, 50}
	fmt.Printf("原始切片: %v\n", slice)

	// 切片操作 [start:end]
	subSlice := slice[1:4] // 索引 1 到 3（不包括 4）
	fmt.Printf("slice[1:4]: %v\n", subSlice)

	// 从开始到指定位置
	fromStart := slice[:3] // 等价于 slice[0:3]
	fmt.Printf("slice[:3]: %v\n", fromStart)

	// 从指定位置到结束
	toEnd := slice[2:] // 等价于 slice[2:len(slice)]
	fmt.Printf("slice[2:]: %v\n", toEnd)

	// 完整的切片
	fullSlice := slice[:] // 等价于 slice[0:len(slice)]
	fmt.Printf("slice[:]: %v\n", fullSlice)
}

// demonstrateMapOperators 映射运算符演示
func demonstrateMapOperators() {
	fmt.Println("\n=== 映射运算符 ===")

	// 创建映射
	m := make(map[string]int)
	m["apple"] = 10
	m["banana"] = 20
	m["orange"] = 30

	fmt.Printf("映射内容: %v\n", m)

	// 访问运算符 []
	value := m["apple"]
	fmt.Printf("m[\"apple\"] = %d\n", value)

	// 检查键是否存在
	if val, exists := m["grape"]; exists {
		fmt.Printf("葡萄的价格: %d\n", val)
	} else {
		fmt.Println("葡萄不存在于映射中")
	}

	// 删除运算符 delete()
	delete(m, "banana")
	fmt.Printf("删除 banana 后: %v\n", m)

	// 长度运算符 len()
	fmt.Printf("映射长度: %d\n", len(m))
}

// demonstrateTypeAssertion 类型断言运算符演示
func demonstrateTypeAssertion() {
	fmt.Println("\n=== 类型断言运算符 ===")

	var i interface{} = "hello"

	// 类型断言 .(type)
	if str, ok := i.(string); ok {
		fmt.Printf("类型断言成功: %q\n", str)
	} else {
		fmt.Println("类型断言失败")
	}

	// 断言为其他类型
	if num, ok := i.(int); ok {
		fmt.Printf("整数值: %d\n", num)
	} else {
		fmt.Println("不是整数类型")
	}
}

// OtherOperatorsDemo 其他运算符演示主函数
func OtherOperatorsDemo() {
	fmt.Println("========== 1.7.6 其他运算符 ==========")

	demonstrateAddressOperators()
	demonstrateChannelOperators()
	demonstratePointerArithmetic()
	demonstrateSliceOperators()
	demonstrateMapOperators()
	demonstrateTypeAssertion()

	fmt.Println("\n=== 其他运算符总结 ===")
	fmt.Println("✅ & * 地址和解引用运算符")
	fmt.Println("✅ <- 通道发送和接收运算符")
	fmt.Println("✅ [] 切片和映射访问运算符")
	fmt.Println("✅ .() 类型断言运算符")
	fmt.Println("✅ Go 不支持指针运算（安全性考虑）")
	fmt.Println("✅ 通过高级类型提供类似功能")
}
