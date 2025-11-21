package array

import "fmt"

// ========== 1.12.1 声明数组 ==========
//
// 数组是具有相同类型的一组已编号且长度固定的数据项序列
// 数组长度是数组类型的一部分，[5]int 和 [10]int 是不同的类型
//
// 声明方式：
// 1. var <array name> [<length>]<type>
// 2. var <array name> = [<length>]<type>{<element1>, <element2>,...}
// 3. var <array name> = [...]<type>{<element1>, <element2>,...}
// 4. var <array name> = [<length>]<type>{<position1>:<element value1>, ...}

// demonstrateArrayDeclaration1 演示数组声明方式1：仅声明
func demonstrateArrayDeclaration1() {
	fmt.Println("=== 1. 仅声明数组（元素为类型零值）===")

	// 仅声明，数组本身已经初始化好了，其中的元素的值为类型的零值
	var a [5]int
	fmt.Println("var a [5]int")
	fmt.Printf("a = %v\n", a)
	fmt.Println("说明：int 类型的零值是 0，所以数组元素都是 0")

	var marr [2]map[string]string
	fmt.Println("\nvar marr [2]map[string]string")
	fmt.Printf("marr = %v\n", marr)
	fmt.Println("说明：map 的零值是 nil，虽然打印出来是 [<nil> <nil>]")
	fmt.Println("⚠️ 注意：不能直接使用 marr[0][\"test\"] = \"1\"，会 panic")
	fmt.Println("      需要先初始化：marr[0] = make(map[string]string)")
	fmt.Println()
}

// demonstrateArrayDeclaration2 演示数组声明方式2：声明并初始化
func demonstrateArrayDeclaration2() {
	fmt.Println("=== 2. 声明并初始化数组 ===")

	// 完整声明方式
	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("var b [5]int = [5]int{1, 2, 3, 4, 5}")
	fmt.Printf("b = %v\n", b)

	// 类型推导声明方式
	var c = [5]string{"c1", "c2", "c3", "c4", "c5"}
	fmt.Println("\nvar c = [5]string{\"c1\", \"c2\", \"c3\", \"c4\", \"c5\"}")
	fmt.Printf("c = %v\n", c)

	// 短变量声明方式
	d := [3]int{3, 2, 1}
	fmt.Println("\nd := [3]int{3, 2, 1}")
	fmt.Printf("d = %v\n", d)
	fmt.Println()
}

// demonstrateArrayDeclaration3 演示数组声明方式3：使用 ... 自动推断长度
func demonstrateArrayDeclaration3() {
	fmt.Println("=== 3. 使用 ... 代替数组长度（编译器自动推断）===")

	// 使用 ... 代替数组长度，编译器会根据初始化时元素个数推断数组长度
	autoLen := [...]string{"auto1", "auto2", "auto3"}
	fmt.Println("autoLen := [...]string{\"auto1\", \"auto2\", \"auto3\"}")
	fmt.Printf("autoLen = %v\n", autoLen)
	fmt.Printf("数组长度 len(autoLen) = %d\n", len(autoLen))
	fmt.Println("说明：编译器根据元素个数推断数组长度为 3")
	fmt.Println()
}

// demonstrateArrayDeclaration4 演示数组声明方式4：指定下标初始化
func demonstrateArrayDeclaration4() {
	fmt.Println("=== 4. 声明时初始化指定下标的元素值 ===")

	// 在已指定数组长度的情况下，对指定下标的元素初始化
	positionInit := [5]string{1: "position1", 3: "position3"}
	fmt.Println("positionInit := [5]string{1: \"position1\", 3: \"position3\"}")
	fmt.Printf("positionInit = %v\n", positionInit)
	fmt.Println("说明：只初始化下标 1 和 3 的元素，其他元素为零值（空字符串）")

	// 可以混合使用位置和索引初始化
	mixedInit := [6]int{1, 2, 4: 100, 200}
	fmt.Println("\nmixedInit := [6]int{1, 2, 4: 100, 200}")
	fmt.Printf("mixedInit = %v\n", mixedInit)
	fmt.Println("说明：前两个元素按顺序初始化，索引4初始化为100，索引5初始化为200")
	fmt.Println()
}

// demonstrateArrayConstraints 演示数组的限制和注意事项
func demonstrateArrayConstraints() {
	fmt.Println("=== 5. 数组的限制和注意事项 ===")

	// 1. 数组长度是类型的一部分
	var arr1 [3]int
	var arr2 [5]int
	fmt.Printf("arr1 的类型: %T\n", arr1)
	fmt.Printf("arr2 的类型: %T\n", arr2)
	fmt.Println("说明：[3]int 和 [5]int 是不同的类型，不能相互赋值")

	// 2. 初始化时，元素个数不能超过数组声明的长度
	// overLen := [2]int{1, 2, 3}  // 编译错误：array index 2 out of bounds [0:2]
	fmt.Println("\n⚠️ 注意：初始化时，元素个数不能超过数组声明的长度")
	fmt.Println("   例如：[2]int{1, 2, 3} 会导致编译错误")

	// 3. 数组是值类型，赋值会复制整个数组
	arr3 := [3]int{1, 2, 3}
	arr4 := arr3 // 复制整个数组
	arr4[0] = 100
	fmt.Printf("\narr3 = %v\n", arr3)
	fmt.Printf("arr4 = %v\n", arr4)
	fmt.Println("说明：修改 arr4 不会影响 arr3，因为数组是值类型")
	fmt.Println()
}

// demonstrateArrayOperations 演示数组的基本操作
func demonstrateArrayOperations() {
	fmt.Println("=== 6. 数组的基本操作 ===")

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr = %v\n", arr)

	// 访问元素
	fmt.Printf("\n访问元素：arr[0] = %d\n", arr[0])
	fmt.Printf("访问元素：arr[4] = %d\n", arr[4])

	// 修改元素
	arr[2] = 100
	fmt.Printf("\n修改元素：arr[2] = 100\n")
	fmt.Printf("修改后：arr = %v\n", arr)

	// 获取数组长度
	fmt.Printf("\n数组长度：len(arr) = %d\n", len(arr))

	// 遍历数组
	fmt.Println("\n遍历数组（使用索引）：")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("  arr[%d] = %d\n", i, arr[i])
	}

	fmt.Println("\n遍历数组（使用 range）：")
	for index, value := range arr {
		fmt.Printf("  arr[%d] = %d\n", index, value)
	}
	fmt.Println()
}

// ArrayDeclarationDemo 数组声明完整演示
func ArrayDeclarationDemo() {
	fmt.Println("========== 1.12.1 声明数组 ==========")
	fmt.Println()
	fmt.Println("数组是具有相同类型的一组已编号且长度固定的数据项序列。")
	fmt.Println("数组长度是数组类型的一部分，[5]int 和 [10]int 是不同的类型。")
	fmt.Println()
	fmt.Println("四种声明方式：")
	fmt.Println("1. var <array name> [<length>]<type>")
	fmt.Println("2. var <array name> = [<length>]<type>{<element1>, <element2>,...}")
	fmt.Println("3. var <array name> = [...]<type>{<element1>, <element2>,...}")
	fmt.Println("4. var <array name> = [<length>]<type>{<position1>:<element value1>, ...}")
	fmt.Println()

	demonstrateArrayDeclaration1()
	demonstrateArrayDeclaration2()
	demonstrateArrayDeclaration3()
	demonstrateArrayDeclaration4()
	demonstrateArrayConstraints()
	demonstrateArrayOperations()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 数组是值类型，长度固定")
	fmt.Println("✅ 数组长度是类型的一部分")
	fmt.Println("✅ 可以使用 ... 让编译器自动推断长度")
	fmt.Println("✅ 可以指定下标初始化特定元素")
	fmt.Println("✅ 未初始化的元素为类型零值")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 数组赋值会复制整个数组（值类型）")
	fmt.Println("   - 初始化元素个数不能超过数组长度")
	fmt.Println("   - map 类型的零值是 nil，使用前需要初始化")
	fmt.Println()
}
