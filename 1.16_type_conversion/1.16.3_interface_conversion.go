package typeconversion

import "fmt"

// ========== 1.16.3 接口类型转换 ==========
//
// 接口类型只能通过断言将转换为指定类型。
//
// 格式：<variable_name>.(<type_name>)
// variable_name 是变量名称，type_name 是类型名称。
//
// 通过断言方式可以同时得到转换后的值以及转换是否成功的标识。

// InterfaceConversionDemo 演示接口类型转换
func InterfaceConversionDemo() {
	fmt.Println("========== 1.16.3 接口类型转换 ==========")
	fmt.Println()
	fmt.Println("接口类型只能通过断言将转换为指定类型。")
	fmt.Println()
	fmt.Println("格式：<variable_name>.(<type_name>)")
	fmt.Println("  - variable_name 是变量名称")
	fmt.Println("  - type_name 是类型名称")
	fmt.Println()
	fmt.Println("通过断言方式可以同时得到转换后的值以及转换是否成功的标识。")
	fmt.Println()

	demonstrateBasicTypeAssertion()
	demonstrateSwitchTypeAssertion()
	demonstrateInterfaceToStructConversion()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 类型断言格式：variable.(Type)")
	fmt.Println("✅ 返回两个值：转换后的值和转换是否成功的布尔值")
	fmt.Println("✅ 使用 switch v := i.(type) 可以同时进行类型判断和值提取")
	fmt.Println("✅ 可以将接口类型转换为实现该接口的具体结构体类型")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 如果转换失败，第一个返回值是对应类型的零值")
	fmt.Println("   - 使用 switch 的方式可能更常见一些")
	fmt.Println("   - 类型断言失败时，需要检查返回的布尔值")
	fmt.Println()
}

// demonstrateBasicTypeAssertion 演示基本的类型断言
func demonstrateBasicTypeAssertion() {
	fmt.Println("=== 1.16.3.1 基本类型断言 ===")

	var i interface{} = 3
	fmt.Printf("i 的类型: %T, 值: %v\n", i, i)

	// 类型断言：同时获取转换后的值和转换是否成功的标识
	a, ok := i.(int)
	if ok {
		fmt.Printf("类型断言成功: '%d' is an int\n", a)
	} else {
		fmt.Println("类型断言失败: conversion failed")
	}
	fmt.Println()

	// 类型断言失败的情况
	var i2 interface{} = "hello"
	fmt.Printf("i2 的类型: %T, 值: %v\n", i2, i2)

	a2, ok2 := i2.(int)
	if ok2 {
		fmt.Printf("类型断言成功: '%d' is an int\n", a2)
	} else {
		fmt.Println("类型断言失败: i2 不是 int 类型")
		fmt.Printf("  ok2 = %v, a2 = %v (零值)\n", ok2, a2)
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 类型断言格式: variable.(Type)")
	fmt.Println("  - 返回两个值：转换后的值和转换是否成功的布尔值")
	fmt.Println("  - 如果转换失败，第一个返回值是对应类型的零值")
	fmt.Println()
}

// demonstrateSwitchTypeAssertion 演示 switch 类型的类型断言
func demonstrateSwitchTypeAssertion() {
	fmt.Println("=== 1.16.3.2 switch 类型的类型断言 ===")

	// 示例1：int 类型
	var i1 interface{} = 3
	fmt.Printf("i1 = %v\n", i1)
	switch v := i1.(type) {
	case int:
		fmt.Printf("  i1 is an int, value: %d\n", v)
	case string:
		fmt.Printf("  i1 is a string, value: %s\n", v)
	default:
		fmt.Printf("  i1 is unknown type, value: %v\n", v)
	}
	fmt.Println()

	// 示例2：string 类型
	var i2 interface{} = "hello"
	fmt.Printf("i2 = %v\n", i2)
	switch v := i2.(type) {
	case int:
		fmt.Printf("  i2 is an int, value: %d\n", v)
	case string:
		fmt.Printf("  i2 is a string, value: %s\n", v)
	default:
		fmt.Printf("  i2 is unknown type, value: %v\n", v)
	}
	fmt.Println()

	// 示例3：其他类型
	var i3 interface{} = 3.14
	fmt.Printf("i3 = %v\n", i3)
	switch v := i3.(type) {
	case int:
		fmt.Printf("  i3 is an int, value: %d\n", v)
	case string:
		fmt.Printf("  i3 is a string, value: %s\n", v)
	case float64:
		fmt.Printf("  i3 is a float64, value: %f\n", v)
	default:
		fmt.Printf("  i3 is unknown type, value: %v\n", v)
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 使用 switch 的方式可能更常见一些")
	fmt.Println("  - switch v := i.(type) 可以同时进行类型判断和值提取")
	fmt.Println()
}

// Supplier 接口定义
type Supplier interface {
	Get() string
}

// DigitSupplier 实现 Supplier 接口
type DigitSupplier struct {
	value int
}

// Get 实现 Supplier 接口的方法
func (i *DigitSupplier) Get() string {
	return fmt.Sprintf("%d", i.value)
}

// demonstrateInterfaceToStructConversion 演示接口类型转换为结构体接口类型
func demonstrateInterfaceToStructConversion() {
	fmt.Println("=== 1.16.3.3 接口类型转换为结构体接口类型 ===")

	var a Supplier = &DigitSupplier{value: 1}
	fmt.Printf("a 的类型: %T\n", a)
	fmt.Printf("a.Get() = %s\n", a.Get())
	fmt.Println()

	// 将接口类型转换为具体的结构体指针类型
	b, ok := a.(*DigitSupplier)
	if ok {
		fmt.Printf("类型断言成功: b = %+v, ok = %v\n", b, ok)
		fmt.Printf("b.value = %d\n", b.value)
	} else {
		fmt.Printf("类型断言失败: ok = %v\n", ok)
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 可以将接口类型转换为实现该接口的具体结构体类型")
	fmt.Println("  - 使用类型断言可以安全地获取具体的结构体实例")
	fmt.Println()
}

