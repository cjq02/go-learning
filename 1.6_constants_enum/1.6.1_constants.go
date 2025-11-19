package constants

import "fmt"

// 常量（Constants）示例
//
// 常量特性：
// 1. 常量的值是在编译期确定的，所以常量定义时必须赋值
// 2. 不能使用方法的返回值为常量赋值
// 3. 常量被定义后，其值不能再被修改
// 4. 常量（包括全局常量和局部常量）被定义后可以不使用
// 5. 常量只能使用基本数据类型：数字、字符串和布尔类型

// Demo 演示常量的各种定义方式和特性
func ConstantsDemo() {
	fmt.Println("=== 常量定义方式示例 ===")

	// 方式 1: const <name> <type> = <value>
	const a int = 1
	fmt.Printf("方式1 - a (int): %d\n", a)

	// 方式 2: const <name> = <value> (类型推导)
	const b = "test"
	fmt.Printf("方式2 - b (string): %s\n", b)

	// 方式 3: const <name3>, <name4>, ... = <value3>, <value4>, ...
	const c, d = 2, "hello"
	fmt.Printf("方式3 - c (int): %d, d (string): %s\n", c, d)

	// 方式 4: const <name5>, <name6>, ... <type> = <value5>, <value6>, ...
	const e, f bool = true, false
	fmt.Printf("方式4 - e (bool): %v, f (bool): %v\n", e, f)

	// 方式 5: 使用小括号包裹多个常量声明
	const (
		h    byte = 3
		i         = "value"
		j, k      = "v", 4
		l, m      = 5, false
	)
	fmt.Printf("方式5 - h (byte): %d, i (string): %s\n", h, i)
	fmt.Printf("方式5 - j (string): %s, k (int): %d\n", j, k)
	fmt.Printf("方式5 - l (int): %d, m (bool): %v\n", l, m)

	const (
		n = 6
	)
	fmt.Printf("方式5 - n (int): %d\n", n)

	// 演示常量不能修改（如果尝试修改会编译错误）
	// a = 2  // 编译错误: cannot assign to a

	// 演示常量可以不使用
	const unusedConst = 999
	// unusedConst 没有被使用，但不会报错

	// 演示常量只能使用基本数据类型
	const (
		numInt    = 42
		numFloat  = 3.14
		str       = "Hello, World!"
		boolTrue  = true
		boolFalse = false
	)

	fmt.Printf("\n基本数据类型常量:\n")
	fmt.Printf("整数: %d\n", numInt)
	fmt.Printf("浮点数: %f\n", numFloat)
	fmt.Printf("字符串: %s\n", str)
	fmt.Printf("布尔值: %v, %v\n", boolTrue, boolFalse)

	// 以下代码会编译错误，因为常量不能使用复杂数据类型：
	// const slice = []int{1, 2, 3}        // 编译错误
	// const arr = [3]int{1, 2, 3}         // 编译错误
	// const m = map[string]int{"a": 1}    // 编译错误
	// const ptr = &numInt                  // 编译错误
	// const st = struct{ x int }{x: 1}     // 编译错误

	// 演示常量不能使用函数返回值
	// const result = getValue()  // 编译错误: const initializer getValue() is not a constant

	// 局部常量示例
	const localConst = "局部常量"
	fmt.Printf("\n局部常量: %s\n", localConst)
}

// 这个函数用于演示常量不能使用函数返回值
// func getValue() int {
//     return 10
// }
